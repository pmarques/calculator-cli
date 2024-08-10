package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// Used for (Cobra) flags.
var Debug bool

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "calculator",
	Short: "Calculator is a simple CLI calculator as a cobra example",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&Debug, "debug", false, "debug output")
}
