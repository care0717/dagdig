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
	Priority() int
}

type taskWorker struct {
	priority int
	tasks    []int32
}

func (t taskWorker) Priority() int {
	return t.priority
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
		if t.completed() {
			return Completed
		}
		return Finished
	}
	return Running
}

func NewTaskWorker(priority string, tasks []int32) TaskWorker {
	if priority == "Low" {
		return &taskWorker{priority: 0, tasks: tasks}
	} else {
		return &taskWorker{priority: 1, tasks: tasks}
	}
}
