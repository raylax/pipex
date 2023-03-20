package worker

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"io"
	"os/exec"
)

type Shell interface {
	Exec(ctx context.Context, command string, stdout io.Writer, stderr io.Writer) (code int, err error)
	Destroy()
}

type ShellHost struct {
}

func (s *ShellHost) Exec(ctx context.Context, command string, stdout io.Writer, stderr io.Writer) (code int, err error) {
	cmd := exec.CommandContext(ctx, "bash", "-c", command)

	if stdout != nil {
		cmd.Stdout = stdout
	}
	if stderr != nil {
		cmd.Stderr = stderr
	}

	err = cmd.Run()

	if ctx.Err() != nil {
		return -1, ctx.Err()
	}

	if exiterr, ok := err.(*exec.ExitError); ok {
		return exiterr.ExitCode(), nil
	}

	return
}

func (s *ShellHost) Destroy() {
	// NOP
}

type ShellContainer struct {
	client      *client.Client
	containerId string
}

func (s *ShellContainer) Exec(ctx context.Context, command string, stdout io.Writer, stderr io.Writer) (code int, err error) {

	createResponse, err := s.client.ContainerExecCreate(ctx, s.containerId, types.ExecConfig{
		AttachStderr: true,
		AttachStdout: true,
		Cmd:          []string{"sh", "-c", command},
	})
	if err != nil {
		return
	}
	execID := createResponse.ID

	attachResponse, err := s.client.ContainerExecAttach(ctx, execID, types.ExecStartCheck{})
	if err != nil {
		return
	}
	outputDone := make(chan error)

	if stdout == nil {
		stdout = io.Discard
	}
	if stderr == nil {
		stderr = io.Discard
	}

	go func() {
		_, err = stdcopy.StdCopy(stdout, stderr, attachResponse.Reader)
		outputDone <- err
	}()

	select {
	case err = <-outputDone:
		if err != nil {
			return -1, err
		}
		break
	case <-ctx.Done():
		return -1, ctx.Err()
	}

	inspectResponse, err := s.client.ContainerExecInspect(ctx, execID)
	if err != nil {
		return -1, err
	}

	return inspectResponse.ExitCode, nil
}

func (s *ShellContainer) Destroy() {
	_ = s.client.Close()
}
