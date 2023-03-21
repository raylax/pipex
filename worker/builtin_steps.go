package worker

type StepRun struct {
}

func (s StepRun) Id() string {
	return "run"
}

func (s StepRun) Run(jc *JobContext) error {
	return nil
}

func (s StepRun) Cleanup(jc *JobContext) StepResult {
	return StepResultSkip
}
