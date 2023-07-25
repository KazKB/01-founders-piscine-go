package main

import (
	"os"
	"strconv"
)

func printResult(result int) {
	var temp int = result
	var output, output2 []byte
	var negative bool

	if result > -9223372036854775808 && result < 9223372036854775807 {
		if result > 0 || result < 0 {
			if result < 0 {
				negative = true
				result *= -1
			}

			for result > 0 {
				temp = result % 10
				output = append(output, byte(temp)+48)
				result /= 10
			}

			for i := len(output) - 1; i >= 0; i-- {
				output2 = append(output2, output[i])
			}

			var negativeSign string = "-"

			if negative {
				os.Stdout.Write([]byte(negativeSign))
			}

			os.Stdout.Write(output2)

		} else if result == 0 {
			var zero string = "0"
			os.Stdout.Write([]byte(zero))
		}

		var newline string = "\n"
		os.Stdout.Write([]byte(newline))
	}
}

func main() {
	var args []string = os.Args

	if len(args) == 4 {
		var num1, num2, result int
		var operation rune
		var error string
		var number1, number2 bool

		for i := range args[1] {
			if (args[1][i] > 47 && args[1][i] < 58) || args[1][0] == 45 {
				number1 = true
			} else {
				number1 = false
				break
			}
		}

		for i := range args[3] {
			if (args[3][i] > 47 && args[3][i] < 58) || args[3][0] == 45 {
				number2 = true
			} else {
				number2 = false
				break
			}
		}

		num1, _ = strconv.Atoi(args[1])
		num2, _ = strconv.Atoi(args[3])
		operation = rune(args[2][0])

		if (num1 > -9223372036854775808 && num1 < 9223372036854775807) && (num2 > -9223372036854775808 && num2 < 9223372036854775807) && (operation == 37 || operation == 42 || operation == 43 || operation == 45 || operation == 47) && number1 && number2 {
			if num2 == 0 && operation == 37 {
				error := "No modulo by 0\n"

				os.Stdout.Write([]byte(error))
			} else if num2 == 0 && operation == 47 {
				error = "No division by 0\n"

				os.Stdout.Write([]byte(error))
			} else if operation == 37 {
				result = num1 % num2
				printResult(result)
			} else if operation == 42 {
				result = num1 * num2
				printResult(result)
			} else if operation == 43 {
				result = num1 + num2
				printResult(result)
			} else if operation == 45 {
				result = num1 - num2
				printResult(result)
			} else if operation == 47 {
				result = num1 / num2
				printResult(result)
			}
		}
	}
}
