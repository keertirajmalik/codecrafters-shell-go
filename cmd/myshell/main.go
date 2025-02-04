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
		case "exit":
			os.Exit(0)

		case "echo":
			fmt.Println(strings.Join(commandArgs[1:], " "))

		case "type":
			fmt.Println(checkBuiltInCommand(commandArgs[1]))

		default:
			fmt.Printf("%s: command not found\n", strings.TrimSpace(command))
		}
	}
}

func checkBuiltInCommand(command string) string {
	var supportedCommands = map[string]bool{
		Exit: true,
		Echo: true,
		Type: true,
	}

	if !supportedCommands[command] {
		return fmt.Sprintf("%s: not found", command)
	}
	return fmt.Sprintf("%s is a shell builtin", command)
}
