package worker

import (
	"bytes"
	"context"
	"github.com/docker/docker/client"
	"testing"
	"time"
)

func TestShellContainer_Exec(t *testing.T) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		t.Error(err)
	}
	shell := &ShellContainer{
		client:      cli,
		containerId: "e773a0876ea2",
	}
	runShellExecCases(t, shell)
}

func TestShellHost_Exec(t *testing.T) {
	runShellExecCases(t, &ShellHost{})
}

func runShellExecCases(t *testing.T, shell Shell) {
	defer shell.Destroy()

	t.Run("echo", func(t *testing.T) {
		var stdout bytes.Buffer
		code, err := shell.Exec(context.Background(), []string{"bash", "-c", "echo -n hi pipex"}, &stdout, nil)
		if stdout.String() != "hi pipex" {
			t.Logf("code:%v err:%v", code, err)
			t.Error(stdout.String())
		}
	})

	t.Run("sleep echo", func(t *testing.T) {
		var stdout bytes.Buffer
		code, err := shell.Exec(context.Background(), []string{"bash", "-c", "sleep 1 && echo -n hi pipex"}, &stdout, nil)
		if stdout.String() != "hi pipex" {
			t.Logf("code:%v err:%v", code, err)
			t.Error(stdout.String())
		}
	})

	t.Run("echo stderr", func(t *testing.T) {
		var stderr bytes.Buffer
		code, err := shell.Exec(context.Background(), []string{"bash", "-c", "echo -n hi pipex >&2"}, nil, &stderr)
		if stderr.String() != "hi pipex" {
			t.Logf("code:%v err:%v", code, err)
			t.Error(stderr.String())
		}
	})

	t.Run("error exit", func(t *testing.T) {
		code, err := shell.Exec(context.Background(), []string{"sh", "-c", "exit 123"}, nil, nil)
		if code != 123 {
			t.Logf("code:%v err:%v", code, err)
			t.Error(code)
		}
	})

	t.Run("error command", func(t *testing.T) {
		code, err := shell.Exec(context.Background(), []string{"sh", "-c", "not-a-valid-command-pipex"}, nil, nil)
		if code != 127 {
			t.Logf("code:%v err:%v", code, err)
			t.Error(123)
		}
	})

	t.Run("timeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		code, err := shell.Exec(ctx, []string{"sh", "-c", "sleep 2 && exit 123"}, nil, nil)
		if code != -1 || err == nil || err.Error() != "context deadline exceeded" {
			t.Errorf("code:%v err:%v", code, err)
		}
	})

}
