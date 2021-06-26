package openapi

import "testing"

type mockTaskWorker struct {
	point  int32
	status Status
}
func (m mockTaskWorker) ExecutingPoint() int32 {
	return m.point
}
func (m mockTaskWorker) Work() Status {
	return m.status
}

func TestWorkerManagerRun(t *testing.T) {
	capacity := 10
	tests := []struct {
		name string
		readyWorker []TaskWorker
		runningWorker []TaskWorker
		tics int
		expectedPoint int
	}{
		{
			name: "normal",
			readyWorker: []TaskWorker{&mockTaskWorker{point: 3}, &mockTaskWorker{point: 2}},
			tics: 1,
			expectedPoint: 5,
		},
		{
			name: "run ready until capacity",
			readyWorker: []TaskWorker{&mockTaskWorker{point: 6}, &mockTaskWorker{point: 5}, &mockTaskWorker{point: 1}},
			tics: 1,
			expectedPoint: 7,
		},
		{
			name: "run ready until capacity - running",
			readyWorker: []TaskWorker{&mockTaskWorker{point: 8}, &mockTaskWorker{point: 5}},
			runningWorker: []TaskWorker{&mockTaskWorker{point: 3}},
			tics: 1,
			expectedPoint: 8,
		},
		{
			name: "switch to ready worker",
			readyWorker: []TaskWorker{&mockTaskWorker{point: 8}},
			runningWorker: []TaskWorker{&mockTaskWorker{point: 3, status: Finished}},
			tics: 2,
			expectedPoint: 8,
		},
		{
			name: "running forever",
			readyWorker: []TaskWorker{&mockTaskWorker{point: 8}},
			runningWorker: []TaskWorker{&mockTaskWorker{point: 3, status: Running}},
			tics: 2,
			expectedPoint: 3,
		},
		{
			name: "remove completed worker",
			readyWorker: []TaskWorker{&mockTaskWorker{point: 1, status: Completed}},
			runningWorker: []TaskWorker{&mockTaskWorker{point: 3, status: Completed}},
			tics: 2,
			expectedPoint: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manager := NewWorkerManager(capacity)
			manager.readyWorkers = tt.readyWorker
			manager.runningWorkers = tt.runningWorker
			var got int
			for i := 0; i < tt.tics; i++ {
				got = manager.Run()
			}
			if got != tt.expectedPoint {
				t.Errorf("Run = %v, want %v", got, tt.expectedPoint)
			}
		})
	}
}
