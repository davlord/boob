package add

import (
	"fmt"
	"os"
	"strings"

	. "github.com/davlord/boob/core/model"
)

func parseArguments(args []string) Bookmark {
	if len(args) != 2 {
		invalidCommandExit()
	}

	url := args[0]
	tags := parseTagsArgument(args[1])

	return Bookmark{
		Url:  url,
		Tags: tags,
	}
}

func invalidCommandExit() {
	fmt.Println("add {url} {tag1,tag2,...}")
	os.Exit(1)
}

func parseTagsArgument(tagsList string) []string {
	return strings.Split(tagsList, ",")
}
