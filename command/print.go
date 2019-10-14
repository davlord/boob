package command

import (
	"fmt"
	"sort"
	"strings"

	"github.com/davlord/boob/core/dao"
)

func Print(dao *dao.BookmarkDao, args []string) error {
	if args == nil {
		return printBookmarks(dao)
	}

	if args[0] == "tags" {
		return printTags(dao)
	}

	return nil
}

func printBookmarks(dao *dao.BookmarkDao) error {
	bookmarks, err := dao.GetAllBookmarks()
	if err != nil {
		return err
	}

	for i, bm := range bookmarks {
		fmt.Printf("%d\t%s\t%s\n", i+1, bm.Url, strings.Join(bm.Tags, ","))
	}
	return nil
}

func printTags(dao *dao.BookmarkDao) error {
	tags, err := dao.GetAllTags()
	if err != nil {
		return err
	}
	sort.Strings(tags)
	for _, tag := range tags {
		fmt.Printf("%s\n", tag)
	}
	return nil
}
