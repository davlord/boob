package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/davlord/boob/command"
	"github.com/davlord/boob/core/dao"
)

func main() {
	if len(os.Args) < 2 {
		showErrorAndExit(errors.New("missing command"))
	}

	// select command
	var database string
	var cmd command.Command
	var err error
	if len(os.Args) > 2 {
		cmd, err = getCommand(os.Args[2])
	}
	if cmd != nil {
		database = os.Args[1]
	} else {
		cmd, err = getCommand(os.Args[1])
	}
	if err != nil {
		showErrorAndExit(err)
	}

	// build command arguments
	var executeArgs []string = nil
	if database != "" && len(os.Args) > 3 {
		executeArgs = os.Args[3:]
	} else if database == "" && len(os.Args) > 2 {
		executeArgs = os.Args[2:]
	}

	// build dao
	dao := dao.BookmarkDao{
		Database: database,
	}

	// execute command
	if err := cmd(&dao, executeArgs); err != nil {
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
	case "edit":
		return command.Edit, nil
	default:
		return nil, errors.New("unknown command")
	}
}

func showErrorAndExit(err error) {
	fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
	os.Exit(1)
}
