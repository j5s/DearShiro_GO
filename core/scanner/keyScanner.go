package scanner

import (
	"DearShiro_GO/core/conn"
	"bufio"
	"fmt"
	"io"
	"os"
)

type KeyScanner struct {
	TargetUrl string
}

func (this KeyScanner) Scan() {
	// TODO: some bug occur in file path
	file, err := os.Open("/Users/jin/Desktop/WorkSpace/GolandProjects/DearShiro_GO/resources/key")
	//file, err := os.Open("resources/key")
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
		connection := conn.NewConnection(this.TargetUrl)
		isFalseKey = connection.CheckFalseKey(key)
	}
	if !isFalseKey {
		println("############################################")
		println("[*]Found key: " + string(key))
	}
}
