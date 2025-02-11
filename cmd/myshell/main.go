package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	Exit = "exit"
	Echo = "echo"
	Type = "type"
)

func main() {
	path := os.Getenv("PATH")

	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("error while reading command: ", err)
			return
		}

		commandArgs := strings.Fields(command)
		switch commandArgs[0] {
		case Exit:
			os.Exit(0)

		case Echo:
			fmt.Println(strings.Join(commandArgs[1:], " "))

		case Type:
			fmt.Println(checkBuiltInCommand(commandArgs[1], path))

		default:
			fmt.Printf("%s: command not found\n", strings.TrimSpace(command))
		}
	}
}
