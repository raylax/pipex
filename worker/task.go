package worker

type Task interface {
	Id() string
	Execute(ctx *TaskContext) error
	Cleanup(ctx *TaskContext) error
}
