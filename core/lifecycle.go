package core

type Lifecycle string

const (
	LifecycleJobInit     = "job_init"
	LifecycleStepInit    = "step_init"
	LifecycleStepCleanup = "step_cleanup"
	LifecycleJobCleanup  = "job_cleanup"
)
