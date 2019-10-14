package command

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/davlord/boob/core/dao"
	. "github.com/davlord/boob/core/model"
)

func Edit(args []string) error {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	bookmarkToEditIndex := id - 1

	var bookmark Bookmark = parseArguments(args[1:])

	// check if a bookmark with the same URL already exists
	existingBookmarkIndex, err := dao.FindBookmarkIndexByUrl(bookmark.Url)
	if err != nil {
		return err
	}
	if existingBookmarkIndex >= 0 && existingBookmarkIndex != bookmarkToEditIndex {
		return errors.New("a bookmark with the same URL already exists")
	}

	// update bookmark
	err = dao.UpdateBookmarkByIndex(bookmarkToEditIndex, &bookmark)
	if err != nil {
		return err
	}
	fmt.Printf("bookmark %d edited :%s", id, bookmark.String())
	return nil
}
