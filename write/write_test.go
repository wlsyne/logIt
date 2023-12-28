package write

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/wlsyne/logIt/utils/config"
	"github.com/wlsyne/logIt/utils/git"
	"os"
	"testing"
)

func readlineBufUtil(params []string) string {
	bufSize := 4096
	var output string

	for _, param := range params {
		len := bufSize - 1 - len(param)%bufSize
		output += fmt.Sprintf("%s\n%s", param, string(make([]byte, len, len)))
	}

	return output
}

func TestWriteResult(t *testing.T) {
	// Mock for config
	mockConfig := config.Config{
		Title:      "test",
		GitBaseUrl: "https://www.example.com",
		ChatIds:    []int{123, 456},
		BotWebhook: "https://www.example.com/webhook",
	}
	//Create temp file
	file, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatalf("Error creating temp file: %v", err)
	}

	err = WriteResult(
		WriteResultParams{
			Config: mockConfig,
			Author: "synwu",
			WriteItems: []WriteItem{
				{
					Commit:  "commit1",
					Type:    "✨ Feat",
					Content: "test",
				},
			},
			FilePath: file.Name(),
		})
	if err != nil {
		t.Fatalf("Error writing file: %v", err)
	}
	//	check File content
	info, err := os.ReadFile(file.Name())
	if err != nil {
		t.Fatalf("Error reading file: %v", err)
	}

	fmt.Println(string(info))
	assert.Equal(t, "# test\n- ✨ Feat: test  [#commit1](https://www.example.com/commit1)\n> Published by <@synwu>\n", string(info))
}

func TestWritePrompt(t *testing.T) {
	//Create pipe to simulate user input
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Error creating pipe: %v", err)
	}
	oldStdin := os.Stdin
	defer func() {
		r.Close()
		w.Close()
		os.Stdin = oldStdin
	}()
	os.Stdin = r
	// Mock for commitList
	commitList := []git.CommitItem{
		{
			Hash:    "commit1",
			Content: "test1",
		},
		{
			Hash:    "commit2",
			Content: "test2",
		},
		{
			Hash:    "commit3",
			Content: "test3",
		},
	}

	// Test for selecting common options and finish
	input := readlineBufUtil([]string{"", "Test", "", "\u001B[B\u001B[B\u001B[B\u001B[B\u001B[B\u001B[B\u001B[B\u001B[B"})
	if _, err := w.WriteString(input); err != nil {
		t.Fatalf("Error writing to pipe: %v", err)
	}
	result, err := WritePrompt(commitList)
	assert.NoError(t, err)
	assert.IsType(t, []WriteItem{}, result)

	// Test for selecting cancel
	input = readlineBufUtil([]string{"", "Test", "", "\u001B[B\u001B[B\u001B[B\u001B[B\u001B[B\u001B[B\u001B[B\u001B[B\u001B[B"})
	if _, err := w.WriteString(input); err != nil {
		t.Fatalf("Error writing to pipe: %v", err)
	}
	result, err = WritePrompt(commitList)

	assert.Error(t, err)
	assert.EqualError(t, err, "user canceled")
}
