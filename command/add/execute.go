package add

import (
	"errors"
	"fmt"

	"github.com/davlord/boob/core/dao"
	. "github.com/davlord/boob/core/model"
)

func Execute(args []string) error {
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
