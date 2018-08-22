package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "aufait",
		Short: "aufait is a CLI for setting up repositories for management under the dependency management tool Au Fait",
		Long:  `Fully automated dependency management tool, set it up and forget it.`,
	}
	project string
	repos   []string
	paths   []string
)

func init() {
	rootCmd.AddCommand(addCmd)
}

// Execute entry point for aufait cli
func Execute() {
	rootCmd.Execute()
}
