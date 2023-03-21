package worker

import "context"

type JobContext struct {
	ctx       context.Context
	Workspace string
	Shell     Shell

	Steps []*Step
	Step  *Step
}

func (c *JobContext) NewCtxWithTimeout(timeout Timeout) (context.Context, context.CancelFunc) {
	return context.WithTimeout(c.ctx, timeout.ToDuration())
}
