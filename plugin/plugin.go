package plugin

import "errors"

type Type int

const (
	TypePipeline Type = 1 << 0
	TypeTask     Type = 1 << 1
	TypeStep     Type = 1 << 2
)

var errorSkip = errors.New("skip")

type Plugin interface {
	ID() string
}

type PipelinePlugin interface {
	OnExecute(ctx *PipelineContext) error
	OnCleanup(ctx *PipelineContext) error
}

type TaskPlugin interface {
	OnExecute(ctx *TaskContext) error
	OnCleanup(ctx *TaskContext) error
}

type StepPlugin interface {
	OnExecute(ctx *StepContext) error
	OnCleanup(ctx *StepContext) error
}
