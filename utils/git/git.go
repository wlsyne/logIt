package git

import (
	"errors"
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

// TODO: migrate it to an new project
func mapSliceFunc[Input, Output any](originalSlice []Input, f func(value Input, index int) Output) []Output {
	result := make([]Output, len(originalSlice))
	for index, value := range originalSlice {
		result[index] = f(value, index)
	}
	return result
}

func GetCommitList() ([]CommitItem, error) {
	//	get commitList by go-git
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
	commitItems := mapSliceFunc(commitList, func(value string, index int) CommitItem {
		item := strings.Split(value, ",")
		return CommitItem{
			Hash:    item[0],
			Content: item[1],
		}
	})
	return []CommitItem{{
		Hash:    "hash",
		Content: "content",
	}, {
		Hash:    "hash",
		Content: "1",
	}}, nil

}
