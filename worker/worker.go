package openapi

type TaskWorker struct {
	tasks []int32
}
func (t TaskWorker) Complete() bool {
	return len(t.tasks) == 0
}
func (t TaskWorker) ExecutingPoint() int32 {
	if t.Complete() {
		return 0
	}
	return t.tasks[0]
}
func (t *TaskWorker) Work() {
	if t.Complete() {
		return
	}
	t.tasks[0]--
	if t.tasks[0] == 0 {
		t.tasks = t.tasks[1:]
	}
}

func NewTaskWorker(tasks []int32) TaskWorker {
	return TaskWorker{tasks: tasks}
}
