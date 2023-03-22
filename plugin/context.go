package plugin

import "github.com/raylax/pipex/core"

type context struct {
	Logger core.Logger
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
