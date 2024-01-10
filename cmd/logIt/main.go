package main

import (
	"github.com/urfave/cli/v2"
	"github.com/wlsyne/logIt/constants"
	"github.com/wlsyne/logIt/utils/config"
	"github.com/wlsyne/logIt/utils/git"
	"github.com/wlsyne/logIt/write"
	"log"
	"os"
)

func handleWrite(cliCTX *cli.Context, configInfo config.Config, changeLogPath string) error {
	err := config.Validator(configInfo, constants.WriteMode)
	if err != nil {
		log.Fatal(err)
		return err
	}

	userName, err := git.GetGitUserName()

	if err != nil {
		log.Fatal(err)
		return err
	}

	commitList, err := git.GetCommitList()

	if err != nil {
		log.Fatal(err)
		return err
	}

	writeItems, err := write.WritePrompt(commitList)

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = write.WriteResult(write.WriteResultParams{
		Config:     configInfo,
		Author:     userName,
		WriteItems: writeItems,
		FilePath:   changeLogPath,
	})

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func main() {
	//get config and validate
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	configFilePath := wd + "/" + constants.ConfigFileName
	changeLogPath := wd + "/" + constants.ChangelogFileName

	configInfo, err := config.GetConfig(configFilePath)

	if err != nil {
		log.Fatal(err)
	}

	app := &cli.App{
		Name:  "logIt",
		Usage: "logIt is a tool for writing logs",
		Commands: []*cli.Command{
			{
				Name:    "write",
				Aliases: []string{"w"},
				Usage:   "writes a changeLog",
				Action: func(context *cli.Context) error {
					return handleWrite(context, configInfo, changeLogPath)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
