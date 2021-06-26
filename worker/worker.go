package openapi

type Status int
const (
	Running Status = iota + 1
	Finished
	Completed
)

type TaskWorker interface {
	ExecutingPoint() int32
	Work() Status
}

type taskWorker struct {
	tasks []int32
}
func (t taskWorker) completed() bool {
	return len(t.tasks) == 0
}
func (t taskWorker) ExecutingPoint() int32 {
	if t.completed() {
		return 0
	}
	return t.tasks[0]
}
func (t *taskWorker) Work() Status {
	if t.completed() {
		return Completed
	}
	t.tasks[0]--
	if t.tasks[0] == 0 {
		t.tasks = t.tasks[1:]
		return Finished
	}
	return Running
}

func NewTaskWorker(tasks []int32) TaskWorker {
	return &taskWorker{tasks: tasks}
}
