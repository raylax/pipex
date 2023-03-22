package core

import (
	"bytes"
)

type Logger interface {
	Log(args ...any)
	Logf(format string, args ...any)
	Warn(args ...any)
	Warnf(format string, args ...any)
	Error(args ...any)
	Errorf(format string, args ...any)
}

type LogWriter struct {
	temp   []byte
	logger Logger
	write  func(logger Logger, message string)
}

func NewLogWriter(logger Logger, write func(logger Logger, message string)) *LogWriter {
	return &LogWriter{logger: logger, write: write}
}

func (l *LogWriter) Write(p []byte) (n int, err error) {
	l.temp = append(l.temp, p...)

	lines := bytes.SplitAfter(l.temp, []byte("\n"))

	for _, line := range lines[:len(lines)-1] {
		l.write(l.logger, string(line[:len(line)-1]))
	}

	lastLine := lines[len(lines)-1]

	if len(lastLine) == 0 {
		l.temp = nil
		return len(p), nil
	} else {
		l.temp = lines[len(lines)-1]
	}

	return len(p), nil
}

func (l *LogWriter) Close() {
	if l.temp == nil {
		return
	}
	l.write(l.logger, string(l.temp))
}
