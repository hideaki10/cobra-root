package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add two integers",
	Long:  `Add two intergers a and b; result = a +b `,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var a, b int
		var err error
		a, err = strconv.Atoi(args[0])
		if err != nil {
			panic("Argument to `add` must be an integer")
		}
		b, err = strconv.Atoi(args[1])
		if err != nil {
			panic("Argument to `add` must be an integer")
		}

		result := calc.Add(a, b, check)
		fmt.Print(result)
	},
}

func init() {
	addCmd.Flags().BoolVar(
		&check,
		"check",
		false,
		"check controls if overflow/underflow check is performed")
	rootCmd.AddCommand(addCmd)
}

func Add(a, b int, check bool) int {
	if check {
		checkAdd(a, b)
	}
	return a + b
}
