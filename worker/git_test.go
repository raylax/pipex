package worker

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestGit_CloneHttp(t *testing.T) {
	withTempDir(func(path string) {
		git := NewGitWithHTTP(path, "https://gitee.com/raylax/test.git", "dev", "", "")
		runCases(t, git)
	})
}

func TestGit_CloneSSH_RSA(t *testing.T) {
	var keysData = `-----BEGIN OPENSSH PRIVATE KEY-----
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
-----END OPENSSH PRIVATE KEY-----
`
	withTempDir(func(path string) {
		git, _ := NewGitWithSSH(path, "git@gitee.com:raylax/test.git", "dev", keysData, "")
		runCases(t, git)
	})
}

func TestGit_CloneSSH_ECDSA(t *testing.T) {
	var keysData = `-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAaAAAABNlY2RzYS
1zaGEyLW5pc3RwMjU2AAAACG5pc3RwMjU2AAAAQQQTwk/2ef4nCxI3fHOOlYfCzKz9O/RZ
3qCLqPI/4ybbIlB26hsE0kr8y0oavl7vy2C652Rf90mmaAwHJAaegki3AAAAqHUjH/V1Ix
/1AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBBPCT/Z5/icLEjd8
c46Vh8LMrP079FneoIuo8j/jJtsiUHbqGwTSSvzLShq+Xu/LYLrnZF/3SaZoDAckBp6CSL
cAAAAgKE3qS3cy0prjVkRnoQ/qz+3nQ+Vumm2/iPZY6wg7frkAAAANZGV2QGRldi5sb2Nh
bAECAw==
-----END OPENSSH PRIVATE KEY-----
`
	withTempDir(func(path string) {
		git, _ := NewGitWithSSH(path, "git@gitee.com:raylax/test.git", "dev", keysData, "")
		runCases(t, git)
	})
}

func runCases(t *testing.T, git *Git) {

	withTempDir(func(path string) {
		repo, err := git.Clone(context.Background(), os.Stdout)
		if err != nil {
			t.Error(err)
		}

		_, err = repo.Branch("dev")

		if err != nil {
			t.Error(err)
		}

		head, _ := repo.Head()
		println(head.String())

	})
}

func withTempDir(fn func(path string)) {
	path, err := ioutil.TempDir("", "git-test")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(path)
	fn(path)
}
