package main

import (
	ShiroScanner "DearShiro_GO/core/scanner"
	"flag"
)

var module string
var target = new(ShiroScanner.ShiroTarget)

func init() {
	flag.StringVar(&module, "m", "", "scan module")
	flag.StringVar(&target.Base, "b", "", "shiro url")
	flag.StringVar(&target.Key, "k", "", "shiro key")
	flag.StringVar(&target.Gadget, "g", "", "gadget want to use")
	flag.StringVar(&target.Command, "c", "", "command want to execute")
}

func main() {
	//module := "key"
	//base := "http://127.0.0.1:8000/login.jsp"
	//key := "123"
	//gadget := "NoCC"
	//command := "whoami"
	factory := ShiroScanner.InitFactory(module)
	scanner := factory(target)
	scanner.Scan()
}
