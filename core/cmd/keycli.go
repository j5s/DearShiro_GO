package cmd

import (
	"DearShiro_GO/core/scanner"
	"github.com/spf13/cobra"
)

var keyFuzzCmd = &cobra.Command{
	Use:   "keyfuzz [url]",
	Short: "To fuzz shiro key",
	// The target url
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rawUrl := args[0]
		scanner := &scanner.KeyScanner{TargetUrl: rawUrl}
		scanner.Scan()
	},
	Example: "keyfuzz http://127.0.0.1:8000",
}

func init() {
	rootCmd.AddCommand(keyFuzzCmd)
}
