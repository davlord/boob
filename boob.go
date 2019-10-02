package main

import (
	"fmt"
	"os"

	"github.com/davlord/boob/commands/add"
)

func main() {
	if len(os.Args) < 2 {
		invalidCommandExit()
	}

	switch os.Args[1] {

	case "add":
		add.Execute(os.Args[2:])
	default:
		invalidCommandExit()
	}

}

func invalidCommandExit() {
	fmt.Println("invalid command")
	os.Exit(1)
}
