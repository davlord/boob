package main

import (
	"fmt"
	"os"

	"github.com/davlord/boob/commands/add"
	"github.com/davlord/boob/commands/browse"
	"github.com/davlord/boob/commands/print"
)

func main() {
	if len(os.Args) < 2 {
		invalidCommandExit()
	}

	var err error = nil
	var executeArgs []string = nil
	if len(os.Args) > 2 {
		executeArgs = os.Args[2:]
	}
	switch os.Args[1] {

	case "add":
		err = add.Execute(executeArgs)
	case "print":
		err = print.Execute(executeArgs)
	case "browse":
		err = browse.Execute(executeArgs)
	default:
		invalidCommandExit()
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
	}

}

func invalidCommandExit() {
	fmt.Println("invalid command")
	os.Exit(1)
}
