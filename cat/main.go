package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var args []string = os.Args

	if len(args) == 1 {
		input, _ := io.ReadAll(os.Stdin)

		os.Stdout.Write(input)
	} else if len(args) > 1 {
		for i := 1; i < len(args); i++ {
			file, err := os.Open(args[i])

			if err != nil {
				z01.PrintRune(69)
				z01.PrintRune(82)
				z01.PrintRune(82)
				z01.PrintRune(79)
				z01.PrintRune(82)
				z01.PrintRune(58)
				z01.PrintRune(32)

				for i := range err.Error() {
					z01.PrintRune(rune(err.Error()[i]))
				}
				z01.PrintRune('\n')
				os.Exit(1)
			} else {
				fileInfo, _ := file.Stat()

				var text []byte = make([]byte, fileInfo.Size())

				file.Read(text)

				for i := range text {
					z01.PrintRune(rune(text[i]))
				}

				file.Close()
			}
		}
	}
}
