package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	arg := os.Args[1:]

	if len(arg) < 1 {
		fmt.Println("File name missing")
		return
	}
	if len(arg) > 1 {
		fmt.Println("Too many arguments")
		return
	}
	var str string
	for _, ch := range arg {
		str += string(ch)
	}
	file, er := os.Open(str)
	if er != nil {
		return
	}
	defer file.Close()

	io.Copy(os.Stdout, file)
}
