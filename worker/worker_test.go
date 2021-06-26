package openapi

import (
	"testing"
)

func TestTaskWorker(t *testing.T) {
	tests := []struct {
		name        string
		tasks       []int32
		tics        int
		expectPoint int32
	}{
		{
			name:        "task finished",
			tasks:       []int32{5, 3},
			tics:        8,
			expectPoint: 0,
		},
		{
			name:        "task already finished",
			tasks:       []int32{5, 3},
			tics:        9,
			expectPoint: 0,
		},
		{
			name:        "task running",
			tasks:       []int32{5, 3},
			tics:        2,
			expectPoint: 3,
		},
		{
			name:        "task switched",
			tasks:       []int32{5, 3},
			tics:        5,
			expectPoint: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			worker := NewTaskWorker("Low", tt.tasks)
			for i := 0; i < tt.tics; i++ {
				worker.Work()
			}
			if got := worker.ExecutingPoint(); got != tt.expectPoint {
				t.Errorf("ExecutingPoint = %v, want %v", got, tt.expectPoint)
			}
		})
	}
}
