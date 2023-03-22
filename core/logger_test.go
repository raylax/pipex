package core

import (
	"fmt"
	"testing"
)

func TestLoggerWriter_Write(t *testing.T) {

	tests := []struct {
		name   string
		input  []byte
		output string
	}{
		{
			name:   "single line",
			input:  []byte("hello world\n"),
			output: "hello world\n",
		},
		{
			name:   "single line\n",
			input:  []byte("hello world\n"),
			output: "hello world\n",
		},
		{
			name:   "multiple lines",
			input:  []byte("foo\nbar\nbaz"),
			output: "foo\nbar\nbaz\n",
		},
		{
			name:   "multiple lines",
			input:  []byte("foo\nbar\nbaz\n"),
			output: "foo\nbar\nbaz\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var logger logger
			writer := &LogWriter{
				logger: &logger,
				write: func(logger Logger, message string) {
					logger.Log(message + "\n")
				},
			}

			_, _ = writer.Write(tt.input)
			writer.Close()

			if got := logger.String(); got != tt.output {
				t.Errorf("unexpected output: got %q, want %q", got, tt.output)
			}
		})
	}
}

func TestLoggerWriter_Write2(t *testing.T) {
	var logger logger
	writer := &LogWriter{
		logger: &logger,
		write: func(logger Logger, message string) {
			logger.Log(message + "\n")
		},
	}

	_, _ = writer.Write([]byte("a"))
	_, _ = writer.Write([]byte("b\n"))
	_, _ = writer.Write([]byte("cdef"))
	writer.Close()

	output := "ab\ncdef\n"
	if got := logger.String(); got != output {
		t.Errorf("unexpected output: got %q, want %q", got, output)
	}
}

type logger struct {
	data []byte
}

func (l *logger) writef(format string, args ...any) {
	l.data = append(l.data, []byte(fmt.Sprintf(format, args...))...)
}

func (l *logger) write(args ...any) {
	l.data = append(l.data, []byte(fmt.Sprint(args...))...)
}

func (l *logger) Log(args ...any) {
	l.write(args...)
}

func (l *logger) Logf(format string, args ...any) {
	l.writef(format, args...)
}

func (l *logger) Warn(args ...any) {
	l.write(args...)
}

func (l *logger) Warnf(format string, args ...any) {
	l.writef(format, args...)
}

func (l *logger) Error(args ...any) {
	l.write(args...)
}

func (l *logger) Errorf(format string, args ...any) {
	l.writef(format, args...)
}

func (l *logger) String() string {
	return string(l.data)
}
