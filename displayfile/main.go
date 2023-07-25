package main

import (
	"fmt"
	"os"
)

func main() {
	var args []string = os.Args

	if len(args) == 1 {
		fmt.Println("File name missing")
	} else if len(args) > 2 {
		fmt.Println("Too many arguments")
	} else if len(args) == 2 {
		file, err := os.Open(args[1])
		if err != nil {
			fmt.Println(err.Error())
		}

		var text []byte = make([]byte, 34)

		file.Read(text)

		fmt.Printf("%s", text)

		file.Close()
	}
}
