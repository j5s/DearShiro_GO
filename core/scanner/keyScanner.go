package scanner

import (
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

	for true {
		line, _, end := reader.ReadLine()
		if end == io.EOF {
			break
		}
		fmt.Println(string(line))
	}
}
