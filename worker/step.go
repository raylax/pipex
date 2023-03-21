package worker

type StepType int

type StepResult int

const (
	StepResultSkip StepResult = iota
	StepResultHandle
)

const (
	StepTypeBuiltin StepType = iota
	StepTypePlugin
)

type BuiltinStep interface {
	Id() string
	Run(ctx *JobContext) error
	Cleanup(ctx *JobContext) StepResult
}

type Step struct {
	Type   StepType
	Name   string
	Define StepDefine
}
