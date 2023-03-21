package builtins

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	history       = make([]string, 0) // store commands. will only store last 20 commands
	errInvalidArg = errors.New("invalid argument")
)

// add command to history
func AddCmd(command string) {
	// if history is not full then add command, else remove first item and add new command
	if len(history) < 20 {
		history = append(history, command)
	} else {
		history = history[1:]
		history = append(history, command)
	}
}

func History(args ...string) error {
	switch len(args) {
	case 0: // print history
		for i, command := range history {
			fmt.Printf("%v %v \n", i+1, command)
		}
		return nil
	case 1:
		if args[0] == "-c" { // clear history
			history = nil
		} else if args[0] == "-h" { // print history without event numbers
			for _, command := range history {
				fmt.Printf("%v \n", command)
			}
		} else if args[0] == "-r" { // print history in reverse order, newest commands first
			for i := len(history) - 1; i >= 0; i-- {
				fmt.Printf("%v %v \n", i+1, history[i])
			}
		} else { // print only last n history commands
			n, err := strconv.Atoi(args[0])               // check if argument is an int
			if err == nil && n > 0 && n <= len(history) { // check if int is within range of size of history
				for i := len(history) - n; i < len(history); i++ {
					fmt.Printf("%v %v \n", i+1, history[i])
				}
			} else if err == nil { // n is an int but not within range of size of history
				return fmt.Errorf("%w: expected integer less than or equal to history list length", errInvalidArg)
			} else if err != nil { // n is not an int and not a valid history option
				return fmt.Errorf("%w: expected valid history option", errInvalidArg)
			}
		}
		return nil
	default:
		return fmt.Errorf("%w: expected zero or one arguments", errInvalidArg)
	}
}
