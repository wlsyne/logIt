package write

import (
	"errors"
	"github.com/manifoldco/promptui"
	"github.com/wlsyne/go-func/sliceFunc"
	"github.com/wlsyne/logIt/utils/git"
	"os"
)

type ChangelogType string
type WriteItem struct {
	Commit  string
	Type    ChangelogType
	Content string
}

const (
	Feat        ChangelogType = "✨ Feat"
	Doc         ChangelogType = "📝 Doc"
	Fix         ChangelogType = "🐛 Fix"
	Style       ChangelogType = "🎨 Style"
	SpeedUp     ChangelogType = "⚡️ SpeedUp"
	Config      ChangelogType = "🔧 Config"
	Test        ChangelogType = "✅ Test"
	BreakChange ChangelogType = "💥 BreakChange"
	Finish      ChangelogType = "Finish"
	Cancel      ChangelogType = "Cancel"
)

var ChangelogTypeList = []ChangelogType{
	Feat,
	Doc,
	Fix,
	Style,
	SpeedUp,
	Config,
	Test,
	BreakChange,
	Finish,
	Cancel,
}

func convertCommitList(commitList []git.CommitItem) []string {
	results := sliceFunc.MapSlice(commitList, func(value git.CommitItem, index int) string {
		return value.Content
	})
	return results
}

func WritePrompt(commitList []git.CommitItem) ([]WriteItem, error) {
	var results []WriteItem
	for {
		//	Select a type
		selectPrompt := promptui.Select{
			Label: "Select Type",
			Items: ChangelogTypeList,
			Stdin: os.Stdin,
		}

		_, selectedType, err := selectPrompt.Run()

		if err != nil {
			return nil, err
		}

		if selectedType == string(Finish) {
			return results, nil
		}

		if selectedType == string(Cancel) {
			return nil, errors.New("user canceled")
		}

		//	Write content
		inputPrompt := promptui.Prompt{
			Label: "type content",
			Stdin: os.Stdin,
		}

		content, err := inputPrompt.Run()

		if err != nil {
			return nil, err
		}

		//	Select a commit
		convertedItems := convertCommitList(commitList)
		commitPrompt := promptui.Select{
			Label: "Select Commit",
			Items: convertedItems,
			Stdin: os.Stdin,
		}

		_, selectedCommit, err := commitPrompt.Run()

		if err != nil {
			return nil, err
		}

		results = append(results, WriteItem{
			Commit:  selectedCommit,
			Type:    ChangelogType(selectedType),
			Content: content,
		})
	}
}
