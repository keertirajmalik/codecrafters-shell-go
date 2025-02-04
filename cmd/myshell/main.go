package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		command = strings.TrimSpace(command)
		if command == "exit 0" {
			os.Exit(0)
		}

		commandArgs := strings.Fields(command)
		switch commandArgs[0] {
		case "echo":
			fmt.Println(strings.Join(commandArgs[1:], " "))

		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
		}
	}
}
