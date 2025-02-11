package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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
		case Exit:
			os.Exit(0)

		case Echo:
			fmt.Println(strings.Join(commandArgs[1:], " "))

		case Type:
			fmt.Println(checkBuiltInCommand(commandArgs[1]))

		default:
			ok, _ := checkExecutableExist(commandArgs[0])
			if !ok {
				fmt.Printf("%s: command not found\n", strings.TrimSpace(command))
				continue
			}

			cmd := exec.Command(commandArgs[0], commandArgs[1:]...)
			stdout, err := cmd.Output()
			if err != nil {
				fmt.Printf("error while running command:%s\n", err)
				continue
			}
			fmt.Print(string(stdout))
		}
	}
}
