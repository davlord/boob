package add

import (
	"fmt"

	"github.com/davlord/boob/core/dao"
	. "github.com/davlord/boob/core/model"
)

func Execute(args []string) {
	var bookmark Bookmark = parseArguments(args)
	dao.CreateBookmark(&bookmark)
	fmt.Println(bookmark)
}
