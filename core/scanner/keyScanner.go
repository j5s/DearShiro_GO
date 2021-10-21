package scanner

import (
	"DearShiro_GO/core/conn"
	"bufio"
	"fmt"
	"io"
	"os"
)

type KeyScanner struct {
	Target *ShiroTarget
}

func (this KeyScanner) Scan() {
	file, err := os.Open("resources/key")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	var isFalseKey = true
	var key []byte
	var end error

	for isFalseKey {
		key, _, end = reader.ReadLine()
		if end == io.EOF {
			break
		}
		connection := conn.NewConnection(this.Target.Base)
		isFalseKey = connection.CheckFalseKey(key)
	}
	if !isFalseKey {
		println("[*]Found key: " + string(key))
	}
}
