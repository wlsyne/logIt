package git

import (
	"errors"
	"fmt"
	"github.com/wlsyne/go-func/sliceFunc"
	"os"
	"os/exec"
	"strings"
)

type CommitItem struct {
	Hash    string
	Content string
}

func GetGitUserName() (string, error) {
	//	get userName by go-git
	wd, err := os.Getwd()
	if err != nil {
		return "", errors.New("Error getting current working directory")
	}
	cmd := exec.Command("git", "config", "user.name")
	cmd.Dir = wd
	out, err := cmd.Output()
	if err != nil {
		return "", errors.New("Error getting git user name, please check the current directory is a git repository")
	}
	return string(out), nil
}

func GetCommitList() ([]CommitItem, error) {
	//	get commitList
	wd, err := os.Getwd()
	if err != nil {
		return nil, errors.New("Error getting current working directory")
	}
	cmd := exec.Command("git", "log", "--pretty=format:%H,%s")
	cmd.Dir = wd
	out, err := cmd.Output()
	if err != nil {
		return nil, errors.New("Error getting git commit list, please check the current directory is a git repository")
	}

	commitList := strings.Split(string(out), "\n")
	commitItems := sliceFunc.MapSlice(commitList, func(value string, index int) CommitItem {
		item := strings.Split(value, ",")
		return CommitItem{
			Hash:    item[0],
			Content: item[1],
		}
	})

	fmt.Println(commitItems)
	return commitItems, nil

}
