package command

import (
	"strconv"

	"github.com/davlord/boob/core/dao"
)

type Remove struct{}

func (remove Remove) Execute(args []string) error {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	if err := dao.DeleteBookmarkByIndex(id - 1); err != nil {
		return err
	}

	return nil
}
