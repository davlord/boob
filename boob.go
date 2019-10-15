package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/davlord/boob/command"
	"github.com/davlord/boob/core/dao"
)

type commandContext struct {
	database    string
	executeFunc command.Command
	executeArgs []string
}

func (cmd *commandContext) execute() error {
	// build dao
	dao := dao.BookmarkDao{
		Database: cmd.database,
	}

	// execute command
	return cmd.executeFunc(&dao, cmd.executeArgs)
}

func main() {
	commandContext, err := parseArgs()
	if err != nil {
		showErrorAndExit(err)
	}

	if err := commandContext.execute(); err != nil {
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

func parseArgs() (*commandContext, error) {
	argsLen := len(os.Args)
	if argsLen < 2 {
		return nil, errors.New("missing command")
	}

	commandContext := commandContext{}

	// select command
	var err error
	if argsLen > 2 {
		commandContext.executeFunc, err = getCommand(os.Args[2])
	}
	if commandContext.executeFunc != nil {
		commandContext.database = os.Args[1]
	} else {
		commandContext.executeFunc, err = getCommand(os.Args[1])
	}
	if err != nil {
		return nil, err
	}

	// build command arguments
	if commandContext.database != "" && argsLen > 3 {
		commandContext.executeArgs = os.Args[3:]
	} else if commandContext.database == "" && argsLen > 2 {
		commandContext.executeArgs = os.Args[2:]
	}

	return &commandContext, nil
}
