package cmd

import (
	"DearShiro_GO/core/scanner"
	"github.com/spf13/cobra"
)

var gadgetFuzzCmd = &cobra.Command{
	Use:   "gadgetfuzz [url]",
	Short: "To fuzz the available gadget",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		urlRaw := args[0]
		scanner := &scanner.GadgetScanner{TargetUrl: urlRaw, TargetShiroKey: key}
		scanner.Scan()
	},
}

func init() {
	rootCmd.AddCommand(gadgetFuzzCmd)
	gadgetFuzzCmd.Flags().StringVarP(&key, "key", "k", "kPH+bIxk5D2deZiIxcaaaA==", "The shiro key")
}
