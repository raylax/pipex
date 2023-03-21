package worker

type logger interface {
	StartJob(name string)
	StartStep(name string)
	EndStep()
	EndJob()
}

type stdLogger struct {
}

func (s stdLogger) StartJob(name string) {
	//TODO implement me
	panic("implement me")
}

func (s stdLogger) StartStep(name string) {
	//TODO implement me
	panic("implement me")
}

func (s stdLogger) EndStep() {
	//TODO implement me
	panic("implement me")
}

func (s stdLogger) EndJob() {
	//TODO implement me
	panic("implement me")
}
