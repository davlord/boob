package command

import "github.com/davlord/boob/core/dao"

type Command func(*dao.BookmarkDao, []string) error
