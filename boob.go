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
	switch os.Args[1] {

	case "add":
		err = add.Execute(os.Args[2:])
	case "print":
		err = print.Execute()
	case "browse":
		err = browse.Execute(os.Args[2])
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
