package golang

import (
	"codegen/internal/flagkeys"
	"codegen/internal/generate"
	"codegen/internal/generate/plugins/golang"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Cmd for `generate“ command, handling code generation for flag accessors
var Cmd = &cobra.Command{
	Use:   "go",
	Short: "Generate Golang flag accessors for OpenFeature.",
	Long:  `Generate Golang flag accessors for OpenFeature.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		params := golang.Params{
			// Probably some conversion applied here, toLower and remove special characters.
			GoPackage: viper.GetString(flagkeys.GoPackageName),
		}
		gen := golang.NewGenerator(params)
		err := generate.CreateFlagAccessors(gen)
		return err
	},
}

func init() {
	Cmd.Flags().String(flagkeys.GoPackageName, "", "Name of the Go package to be generated.")
	Cmd.MarkFlagRequired(flagkeys.GoPackageName)
	viper.BindPFlag(flagkeys.GoPackageName, Cmd.Flags().Lookup(flagkeys.GoPackageName))

}
