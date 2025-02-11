package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func checkBuiltInCommand(command, path string) string {
	var supportedCommands = map[string]bool{
		Exit: true,
		Echo: true,
		Type: true,
	}

	if supportedCommands[command] {
		return fmt.Sprintf("%s is a shell builtin", command)
	}

	for _, envPath := range strings.Split(path, ":") {
		executablePath := envPath + "/" + command
		if checkExecutableExist(executablePath) {
			return fmt.Sprintf("%s is %s", command, executablePath)
		}
	}
	return fmt.Sprintf("%s: not found", command)
}

func checkExecutableExist(path string) bool {
	_, err := os.Stat(path)

	return !errors.Is(err, os.ErrNotExist)
}
