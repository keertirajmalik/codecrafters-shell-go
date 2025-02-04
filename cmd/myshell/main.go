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
		}

		if strings.TrimSpace(command) == "exit 0" {
			os.Exit(0)
		}
		fmt.Fprintf(os.Stdout, "%s: command not found\n", strings.TrimSpace(command))
	}
}
