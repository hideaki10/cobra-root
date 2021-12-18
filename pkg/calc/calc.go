package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "calc",
	Short: "Calculate arithmetic expressiondadasd",
	Long:  "Calculate arithmetic expressionlong",

	//Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
