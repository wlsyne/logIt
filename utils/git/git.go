package git

import (
	"errors"
	"os"
	"os/exec"
)

func GetGitUserName() (string, error) {
	//	get userName by go-git
	wd, err := os.Getwd()
	if err != nil {
		return "", errors.New("Error getting current working directory")
	}
	cmd := exec.Command("git", "config", "user.name")
	cmd.Dir = wd
	return "", nil
}
