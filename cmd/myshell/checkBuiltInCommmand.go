package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func checkBuiltInCommand(command string) string {
	var supportedCommands = map[string]bool{
		Exit: true,
		Echo: true,
		Type: true,
	}

	if supportedCommands[command] {
		return fmt.Sprintf("%s is a shell builtin", command)
	}

	ok, path := checkExecutableExist(command)
	if ok {
		return fmt.Sprintf("%s is %s", command, path)
	}
	return fmt.Sprintf("%s: not found", command)
}

func checkExecutableExist(command string) (bool, string) {
	for _, envPath := range strings.Split(os.Getenv("PATH"), ":") {
		executablePath := envPath + "/" + command
		_, err := os.Stat(executablePath)
		if !errors.Is(err, os.ErrNotExist) {
			return true, executablePath
		}
	}
	return false, ""
}
