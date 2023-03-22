package worker

import (
	gocontext "context"
	"errors"
	"fmt"
	"github.com/raylax/pipex/core"
)

type StepRun struct {
}

func (s StepRun) Id() string {
	return "run"
}

func (s StepRun) Execute(ctx *StepContext) error {
	var shell []string
	var commands []string
	define := ctx.Define
	switch {
	case define.Sh != nil:
		shell = []string{"sh", "-c"}
		commands = define.Sh
	case define.Bash != nil:
		shell = []string{"bash", "-c"}
		commands = define.Bash
	case define.Cmd != nil:
		shell = []string{"cmd", "/c"}
		commands = define.Cmd
	case define.Powershell != nil:
		shell = []string{"powershell", "-c"}
		commands = define.Powershell
	default:
		return errorStepSkip
	}
	context, cancel := ctx.NewContextWithTimeout(define.Timeout)
	defer cancel()
	for _, command := range commands {
		err := s.executeCommand(ctx, context, shell, command)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s StepRun) executeCommand(ctx *StepContext, context gocontext.Context, shell []string, command string) error {
	stdout := core.NewLogWriter(ctx.Logger, func(logger core.Logger, message string) {
		logger.Log(message)
	})
	defer stdout.Close()
	stderr := core.NewLogWriter(ctx.Logger, func(logger core.Logger, message string) {
		logger.Error(message)
	})
	defer stdout.Close()
	code, err := ctx.Shell.Exec(context, append(shell, command), stdout, stderr)
	if err != nil {
		return err
	}
	if code != 0 {
		return errors.New(fmt.Sprintf("command failed with exit code %d", code))
	}
	return nil
}

func (s StepRun) Cleanup(ctx *StepContext) error {
	return errorStepSkip
}
