package data

import (
	"errors"
	gososerial "github.com/EmYiQing/Gososerial"
)

const SimplePrincipalCollectionHex = "ACED0005737200326F72672E6170616368652E736869726F2E7375626A6563742E53696D706C655072696E636970616C436F6C6C656374696F6EA87F5825C6A3084A0300014C000F7265616C6D5072696E636970616C7374000F4C6A6176612F7574696C2F4D61703B78707077010078"

type payloadMap struct {
	Command   string
	NamedFunc map[string]func(cmd string) []byte
}

var instance *payloadMap

//func (this *payloadMap) AddAllPayload() {
//	this.NamedFunc = make(map[string]func(cmd string) []byte)
//	// TODO: The SerialUiD is different from shiro
//	this.NamedFunc["NOCC"] = gososerial.GetCB1
//	// CCK1 is OK
//	this.NamedFunc["CCK1"] = gososerial.GetCCK1
//}

func NewPayloadMap() *payloadMap {
	if instance == nil {
		instance = new(payloadMap)
		instance.NamedFunc = make(map[string]func(cmd string) []byte)
		// TODO: The SerialUiD is different from shiro
		instance.NamedFunc["NOCC"] = gososerial.GetCB1
		// CCK1 is OK
		instance.NamedFunc["CCK1"] = gososerial.GetCCK1
		return instance
	}
	return instance
}

func (this *payloadMap) GetPayload(gadgetName, command string) ([]byte, error) {
	serialFunc, ok := this.NamedFunc[gadgetName]
	if !ok {
		return make([]byte, 0), errors.New("payload not found")
	}
	return serialFunc(command), nil
}

func (this *payloadMap) AddPayload(payloadName string, serialFunc func(cmd string) []byte) {
	// TODO
}
