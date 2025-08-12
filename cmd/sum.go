package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func sum(numbers []int, debug bool) int {
	result := numbers[0]

	for _, num := range numbers[1:] {
		if debug {
			fmt.Printf("Add %d to %d\n", num, result)
		}

		result += num
	}

	return result
}

func newSumCmd(rootCmd *cobra.Command) {
	sumCmd := &cobra.Command{
		//nolint:dupword
		Use:   "sum number number [...number]",
		Short: "Sum all numbers",
		//nolint:mnd
		Args: cobra.MinimumNArgs(2),
		Run: func(_ *cobra.Command, args []string) {
			debug, _ := rootCmd.Flags().GetBool("debug")

			numbers := StringListToIntList(args)

			result := sum(numbers, debug)

			fmt.Println(result)
		},
	}

	rootCmd.AddCommand(sumCmd)
}
