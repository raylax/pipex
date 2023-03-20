package worker

type Mode = string

const (
	ModeHost      Mode = "host"
	ModeContainer Mode = "container"
)
