package print

import (
	"fmt"
	"strings"

	"github.com/davlord/boob/core/dao"
)

func Execute(args []string) error {
	if args == nil {
		return printBookmarks()
	}

	if args[0] == "tags" {
		return printTags()
	}

	return nil
}

func printBookmarks() error {
	bookmarks, err := dao.GetAllBookmarks()
	if err != nil {
		return err
	}

	for i, bm := range bookmarks {
		fmt.Printf("%d\t%s\t%s\n", i+1, bm.Url, strings.Join(bm.Tags, ","))
	}
	return nil
}

func printTags() error {
	tags, err := dao.GetAllTags()
	if err != nil {
		return err
	}
	for _, tag := range tags {
		fmt.Printf("%s\n", tag)
	}
	return nil
}
