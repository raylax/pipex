package worker

type StepType int

const (
	StepTypeBuiltin StepType = iota
	StepTypePlugin
)

type BuiltinStep interface {
	RequiredKeys() []string
}

type Step struct {
	Type StepType
	Name string
}
