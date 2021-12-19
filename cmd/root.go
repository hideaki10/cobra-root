package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var check bool

var rootCmd = &cobra.Command{
	Use:   "calc",
	Short: "Calculate arithmetic expression",
	Long:  "Calculate arithmetic expression",

	//Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
