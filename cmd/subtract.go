package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var subtractCmd = &cobra.Command{
	Use:   "subtract",
	Short: "Subtract two integers",
	Long:  `Subtract two intergers a and b; result = a -b `,
	Args:  cobra.ExactArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		var a, b int
		var err error
		a, err = strconv.Atoi(args[0])
		if err != nil {
			panic("Arguments to `substract` must be integers")
		}
		b, err = strconv.Atoi(args[1])
		if err != nil {
			panic("Arguments to `substract` must be integers")
		}
		result := calc.Subtract(a, b, check)
		fmt.Print(result)

	},
}

func init() {
	subtractCmd.Flags().BoolVar(
		&check,
		"check",
		false,
		"check controls if overflow/underflow check is performed")
	rootCmd.AddCommand(subtractCmd)
}
