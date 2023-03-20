package plugin

import "github.com/raylax/pipex/worker"

type Type string

const (
	TypeLifecycle Type = "lifecycle"
	TypeStep      Type = "step"
)

type Result struct {
}

type Plugin interface {
	ID() string
	Init(ctx worker.JobContext) error
	Destroy()
}

type LifecyclePlugin interface {
	Plugin
	OnJobInit(ctx *worker.JobContext) Result
	OnStepInit(ctx *worker.JobContext) Result
	OnStepCleanup(ctx *worker.JobContext) Result
	OnJobCleanup(ctx *worker.JobContext) Result
}

type StepPlugin interface {
	Plugin
	OnInit(ctx *worker.JobContext) Result
	OnRun(ctx *worker.JobContext) Result
	OnCleanup(ctx *worker.JobContext) Result
}
