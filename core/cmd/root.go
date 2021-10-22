package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	key     string
	gadget  string
	command string
)

var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
