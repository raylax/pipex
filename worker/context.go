package worker

type JobContext struct {
	Workspace string
	Shell     *Shell
}
