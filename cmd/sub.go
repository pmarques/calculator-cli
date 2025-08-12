package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func sub(numbers []int, debug bool) int {
	result := numbers[0]

	for _, num := range numbers[1:] {
		if debug {
			fmt.Printf("Sub %d from %d\n", num, result)
		}

		result -= num
	}

	return result
}

func newSubCmd() *cobra.Command {
	subCmd := &cobra.Command{
		//nolint:dupword
		Use:   "sub number number [...number]",
		Short: "Subtract all numbers",
		//nolint:mnd
		Args: cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			debug, _ := cmd.Root().Flags().GetBool("debug")

			numbers := StringListToIntList(args)

			result := sub(numbers, debug)

			fmt.Println(result)
		},
	}

	return subCmd
}
