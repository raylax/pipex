package worker

import (
	"context"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"io"
)

type Git struct {
	url    string
	branch string
	path   string
	auth   transport.AuthMethod
}

func NewGitWithHTTP(path, url, branch, username, password string) *Git {
	var auth transport.AuthMethod
	if username != "" || password != "" {
		auth = &http.BasicAuth{
			Username: username,
			Password: password,
		}
	}
	return newGit(path, url, branch, auth)
}

func NewGitWithSSH(path, url, branch, sshKeys, sshKeysPassword string) (*Git, error) {
	auth, err := ssh.NewPublicKeys("git", []byte(sshKeys), sshKeysPassword)
	if err != nil {
		return nil, err
	}
	return newGit(path, url, branch, auth), nil
}

func newGit(path, url, branch string, auth transport.AuthMethod) *Git {
	return &Git{
		url:    url,
		branch: branch,
		path:   path,
		auth:   auth,
	}
}

func (g *Git) Clone(ctx context.Context, stdout io.Writer) (*git.Repository, error) {
	return git.PlainCloneContext(ctx, g.path, false, &git.CloneOptions{
		URL:               g.url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Progress:          stdout,
		Auth:              g.auth,
		ReferenceName:     plumbing.NewBranchReferenceName(g.branch),
	})
}
