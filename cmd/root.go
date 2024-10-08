package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Used for (Cobra) flags.
	var Debug bool

	// rootCmd represents the base command when called without any subcommands.
	rootCmd := &cobra.Command{
		Use:   "calculator",
		Short: "Calculator is a simple CLI calculator as a cobra example",
	}

	rootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "debug output")

	// Testing 2 different initialisatons
	rootCmd.AddCommand(newSubCmd())
	newSumCmd(rootCmd)

	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("Error running cobra command %v\n", err)
	}
}
