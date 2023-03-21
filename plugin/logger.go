package plugin

type Logger interface {
	Log(args ...any)
	Logf(format string, args ...any)
	Warn(args ...any)
	Warnf(format string, args ...any)
	Error(args ...any)
	Errorf(format string, args ...any)
}
