package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var sumCmd = &cobra.Command{
	Use:   "sum",
	Short: "Sum all numbers",
	Run:   sum,
}

func init() {
	rootCmd.AddCommand(sumCmd)
}

func sum(_ *cobra.Command, args []string) {
	sub := StringToInt(args[0])

	for _, arg := range args[1:] {
		num := StringToInt(arg)

		if Debug {
			fmt.Printf("Add %d to %d\n", num, sub)
		}

		sub += num
	}

	fmt.Println(sub)
}
