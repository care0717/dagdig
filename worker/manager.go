package openapi

type WorkerManager struct {
	capacity       int
	readyWorkers   []TaskWorker
	runningWorkers []TaskWorker
}

func NewWorkerManager(capacity int) WorkerManager {
	return WorkerManager{capacity: capacity}
}
func (m *WorkerManager) Add(worker TaskWorker) {
	m.readyWorkers = append(m.readyWorkers, worker)
}
func (m *WorkerManager) Run() int {
	var runningPoint int
	currentRunningWorkers := make([]TaskWorker, len(m.runningWorkers))
	for i, w := range m.runningWorkers {
		runningPoint += int(w.ExecutingPoint())
		currentRunningWorkers[i] = w
	}
	var nextReadyWorkers []TaskWorker
	for _, w := range m.readyWorkers {
		if m.capacity >= runningPoint+int(w.ExecutingPoint()) {
			runningPoint += int(w.ExecutingPoint())
			currentRunningWorkers = append(currentRunningWorkers, w)
		} else {
			nextReadyWorkers = append(nextReadyWorkers, w)
		}
	}
	var nextRunningWorkers []TaskWorker
	for _, w := range currentRunningWorkers {
		status := w.Work()
		switch status {
		case Running:
			nextRunningWorkers = append(nextRunningWorkers, w)
		case Finished:
			nextReadyWorkers = append(nextReadyWorkers, w)
		}
	}
	m.readyWorkers = nextReadyWorkers
	m.runningWorkers = nextRunningWorkers
	return runningPoint
}
