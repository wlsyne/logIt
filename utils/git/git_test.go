package git

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetGitUserName(t *testing.T) {
	userName, _ := GetGitUserName()
	assert.IsType(t, "string", userName)
}

//func TestGetCommitList(t *testing.T) {
//	commitList, _ := GetCommitList()
//	assert.IsType(t, []CommitItem{}, commitList)
//}
