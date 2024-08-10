package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "Subtract all numbers",
	Run:   sub,
}

func init() {
	rootCmd.AddCommand(subCmd)
}

func sub(_ *cobra.Command, args []string) {
	sub := StringToInt(args[0])

	for _, arg := range args[1:] {
		num := StringToInt(arg)
		sub -= num
	}

	fmt.Println(sub)
}
