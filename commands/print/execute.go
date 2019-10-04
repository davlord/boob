package print

import (
	"fmt"
	"strings"

	"github.com/davlord/boob/core/dao"
)

func Execute() error {
	bookmarks, err := dao.GetAllBookmarks()
	if err != nil {
		return err
	}

	for i, bm := range bookmarks {
		fmt.Printf("%d\t%s\t%s\n", i+1, bm.Url, strings.Join(bm.Tags, ","))
	}

	return nil
}
