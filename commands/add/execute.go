package add

import (
	"fmt"

	"github.com/davlord/boob/core/dao"
	. "github.com/davlord/boob/core/model"
)

func Execute(args []string) error {
	var bookmark Bookmark = parseArguments(args)
	err := dao.CreateBookmark(&bookmark)
	if err != nil {
		return err
	}
	fmt.Println("bookmark added :" + bookmark.String())
	return nil
}
