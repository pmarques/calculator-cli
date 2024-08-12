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
		Use:   "sum",
		Short: "Sum all numbers",
		Run: func(_ *cobra.Command, args []string) {
			debug, _ := rootCmd.Root().Flags().GetBool("debug")

			numbers := StringListToIntList(args)

			result := sum(numbers, debug)

			fmt.Println(result)
		},
	}

	rootCmd.AddCommand(sumCmd)
}
