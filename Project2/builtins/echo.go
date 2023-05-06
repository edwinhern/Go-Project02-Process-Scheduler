package builtins

import (
	"fmt"
)

func Echo(args ...string) error {
	switch len(args) {
	case 0:
		// No argument given
		fmt.Printf("%s\n", "")
	default:
		for i := 0; i < len(args); i++ {
			str := args[i]
			// Print arg if " not found at beginning of argument
			if str[0] != '"' {
				fmt.Printf("%s\n", str)
			} else {
				var start = i
				var count = i
				var endChar = str[len(str)-1]

				// Loop until finding arg that ends in "
				for endChar != '"' {
					i++
					count++
					str = args[i]
					endChar = str[len(str)-1]
				}

				// Print of the arguments in "", The arguments with a " have them deleted before output to terminal
				for j := start; j <= count; j++ {
					if j == start && j == count {
						catStr := args[j]
						catStr = catStr[1 : len(catStr)-1]
						fmt.Printf("%s ", catStr)
					} else if j == start && j != count {
						catStr := args[j]
						catStr = catStr[1:]
						fmt.Printf("%s ", catStr)
					} else if j != start && j == count {
						catStr := args[j]
						catStr = catStr[:len(catStr)-1]
						fmt.Printf("%s ", catStr)
					} else {
						catStr := args[j]
						fmt.Printf("%s ", catStr)
					}
				}
				fmt.Printf("%s\n", "")
			}
		}
	}
	return nil
}
