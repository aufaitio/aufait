package cmd

import (
	"fmt"
	"github.com/quantumew/aufait/app"
	"github.com/spf13/cobra"
	"os"
)

var (
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "Configure the aufait tool",
		Run:   runAdd,
	}
	path        string
	remote      string
	branch      string
	listenerURL string
	remoteName  string
)

func init() {
	addCmd.PersistentFlags().StringVarP(&path, "path", "p", ".", "Configure a local repositories")
	addCmd.PersistentFlags().StringVarP(&remote, "remote", "r", "", "Override the default remote")
	addCmd.PersistentFlags().StringVarP(&remoteName, "remote-name", "n", "origin", "Override the default origin remote name")
	addCmd.PersistentFlags().StringVarP(&branch, "branch", "b", "master", "Override the output of git branch")
	addCmd.PersistentFlags().StringVarP(&listenerURL, "url", "u", "https://quantumew.github.io/aufait/", "Override the output of git branch")
}

func runAdd(cmd *cobra.Command, args []string) {
	cli, err := app.NewCLI(listenerURL)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = cli.ConfigureLocalRepository(path, remote, branch, remoteName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
