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

func init() {
	// ?
	rootCmd.SetHelpTemplate("Usage:\n   [flags]\n   [command]\nAvailable Commands:\n  commandexec  use the gadget to exec command\n  gadgetfuzz  To fuzz the available gadget\n  keyfuzz     To fuzz shiro key\n\nFlags:\n  -h, --help   help for this command\n\nUse \" [command] --help\" for more information about a command.")
}
