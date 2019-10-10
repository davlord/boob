package command

import (
	"errors"
	"os/exec"
	"strconv"

	"github.com/davlord/boob/core/dao"
)

type Browse struct{}

func (browse Browse) Execute(args []string) error {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	// check if a bookmark with the same URL already exists
	existingBookmark, err := dao.FindBookmarkByIndex(id - 1)
	if err != nil {
		return err
	}
	if existingBookmark == nil {
		return errors.New("invalid id, bookmark not found")
	}

	cmd := exec.Command("xdg-open", existingBookmark.Url)
	return cmd.Run()
}
