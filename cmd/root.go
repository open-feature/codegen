package cmd

import (
	"fmt"
	"os"

	"codegen/cmd/generate"

	"github.com/spf13/cobra"
)

var (
	Version string
	Commit  string
	Date    string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "openfeature",
	Short: "CLI for OpenFeature.",
	Long:  `CLI for OpenFeature related functionalities.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string, commit string, date string) {
	Version = version
	Commit = commit
	Date = date
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(generate.Root)
	rootCmd.AddCommand(versionCmd)
}
