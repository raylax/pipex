package worker

import (
	"context"
	"github.com/raylax/pipex/core"
)

type PipelineContext struct {
	ctx    context.Context
	Logger core.Logger
}

func (c *PipelineContext) NewContextWithTimeout(timeout Timeout) (context.Context, context.CancelFunc) {
	return context.WithTimeout(c.ctx, timeout.ToDuration())
}

type TaskContext struct {
	ctx       context.Context
	Pipeline  *PipelineContext
	Workspace string
	Shell     Shell
	Logger    core.Logger
}

func (c *TaskContext) NewContextWithTimeout(timeout Timeout) (context.Context, context.CancelFunc) {
	return context.WithTimeout(c.ctx, timeout.ToDuration())
}

type StepContext struct {
	ctx       context.Context
	Task      *TaskContext
	Workspace string
	Shell     Shell
	Define    StepDefine
	Logger    core.Logger
}

func (c *StepContext) NewContextWithTimeout(timeout Timeout) (context.Context, context.CancelFunc) {
	return context.WithTimeout(c.ctx, timeout.ToDuration())
}
