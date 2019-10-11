package command

import (
	"strconv"

	"github.com/davlord/boob/core/dao"
)

func Remove(args []string) error {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	if err := dao.DeleteBookmarkByIndex(id - 1); err != nil {
		return err
	}

	return nil
}
