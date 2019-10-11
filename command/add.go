package command

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/davlord/boob/core/dao"
	. "github.com/davlord/boob/core/model"
)

func Add(args []string) error {
	var bookmark Bookmark = parseArguments(args)

	// check if a bookmark with the same URL already exists
	existingBookmark, err := dao.FindBookmarkByUrl(bookmark.Url)
	if err != nil {
		return err
	}
	if existingBookmark != nil {
		return errors.New("a bookmark with the same URL already exists")
	}

	// create bookmark
	err = dao.CreateBookmark(&bookmark)
	if err != nil {
		return err
	}
	fmt.Println("bookmark added :" + bookmark.String())
	return nil
}

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
