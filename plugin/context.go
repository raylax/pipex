package plugin

type context struct {
	Logger Logger
	Env    Env
}

type PipelineContext struct {
	context
}

type TaskContext struct {
	PipelineContext
	Task Task
}

type StepContext struct {
	TaskContext
	Step Step
}
