package worker

import "errors"

var errorStepSkip = errors.New("STEP SKIP")

type Step interface {
	Id() string
	Execute(ctx *StepContext) error
	Cleanup(ctx *StepContext) error
}
