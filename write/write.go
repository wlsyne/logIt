package write

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/wlsyne/go-func/sliceFunc"
	"github.com/wlsyne/logIt/utils/config"
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

type WriteResultParams struct {
	Config     config.Config
	Author     string
	WriteItems []WriteItem
	FilePath   string
}

func WriteResult(params WriteResultParams) error {
	filePath, author, writeItems, packageConfig := params.FilePath, params.Author, params.WriteItems, params.Config

	//Handle changelog title
	content := "# " + packageConfig.Title + "\n"

	//handle changelog content
	for _, writeItem := range writeItems {
		content += "- " + string(writeItem.Type) + ": " + writeItem.Content + "  " + "[#" + writeItem.Commit + "](" + packageConfig.GitBaseUrl + "/" + writeItem.Commit + ")\n"
	}

	//handle changelog footer
	content += "> Published by <@" + author + ">" + "\n"

	//Write to file
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	if err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}
	fmt.Println(content)
	_, err = file.WriteString(content)
	fmt.Println(err)

	if err != nil {
		return err
	}

	return nil
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
