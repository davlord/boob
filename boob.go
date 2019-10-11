package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/davlord/boob/command"
)

func main() {
	if len(os.Args) < 2 {
		showErrorAndExit(errors.New("missing command"))
	}

	// select command
	cmd, err := getCommand(os.Args[1])
	if err != nil {
		showErrorAndExit(err)
	}

	// build command arguments
	var executeArgs []string = nil
	if len(os.Args) > 2 {
		executeArgs = os.Args[2:]
	}

	// execute command
	if err := cmd(executeArgs); err != nil {
		showErrorAndExit(err)
	}
}

func getCommand(commandName string) (command.Command, error) {
	switch commandName {

	case "add":
		return command.Add, nil
	case "print":
		return command.Print, nil
	case "browse":
		return command.Browse, nil
	case "remove":
		return command.Remove, nil
	default:
		return nil, errors.New("unknown command")
	}
}

func showErrorAndExit(err error) {
	fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
	os.Exit(1)
}
