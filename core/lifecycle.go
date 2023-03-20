package core

type Lifecycle string

const (
	LifecycleJobInit     Lifecycle = "job_init"
	LifecycleStepInit    Lifecycle = "step_init"
	LifecycleStepCleanup Lifecycle = "step_cleanup"
	LifecycleJobCleanup  Lifecycle = "job_cleanup"
)
