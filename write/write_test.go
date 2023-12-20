package write

import (
	"github.com/manifoldco/promptui"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetWriteResult(t *testing.T) {

	// When users select cancel,

	//	When users select finish,

	//	When users select other common options like: feat, fix etc.,

}

func TestSelectPrompt(t *testing.T) {
	// Define the prompt options
	options := []string{"Option 1", "Option 2", "Option 3"}

	// Define the prompt
	prompt := promptui.Select{
		Label: "Select an option",
		Items: options,
	}

	// Simulate user input
	prompt.Templates = &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "\U000027A4 {{ . | cyan }}",
		Inactive: "  {{ . | white }}",
		Selected: "\U00002705 {{ . | green }}",
	}

	// Run the prompt
	_, result, err := prompt.Run()

	// Verify the result
	assert.NoError(t, err)
	assert.Equal(t, "Option 1", result)
}

// Do you know a library called promptui in go, if I use it to write a select prompt, how can I test it? I mean to give some mock options and simulate the user's actions like: user press down arrow, select the second option, press enter, etc.,
