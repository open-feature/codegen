package react

import (
	"codegen/internal/generate"
	"codegen/internal/generate/plugins/react"

	"github.com/spf13/cobra"
)

// Cmd for "generate" command, handling code generation for flag accessors
var Cmd = &cobra.Command{
	Use:   "react",
	Short: "Generate typesafe React Hooks.",
	Long:  `Generate typesafe React Hooks compatible with the OpenFeature React SDK.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		params := react.Params{}
		gen := react.NewGenerator(params)
		err := generate.CreateFlagAccessors(gen)
		return err
	},
}

func init() {
}
