package dao

import (
	"os"
	"path/filepath"
	"strings"

	. "github.com/davlord/boob/core/model"
)

const (
	FILE_MODE_WRITE = os.O_APPEND | os.O_CREATE | os.O_WRONLY
	FILE_PERMISSION = 0644
	DIR_PERMISSION  = 0755
)

func CreateBookmark(bookmark *Bookmark) error {

	f, err := openDatabaseFile(FILE_MODE_WRITE)
	if err != nil {
		return err
	}
	defer f.Close()

	serializedBookmark := serializeBookmark(bookmark)
	if _, err = f.WriteString(serializedBookmark + "\n"); err != nil {
		return err
	}

	return nil
}

func openDatabaseFile(modeFlags int) (*os.File, error) {
	filePath := getDatabaseFile()
	createDirectoryIfNeeded(filePath)

	return os.OpenFile(filePath, modeFlags, FILE_PERMISSION)
}

func getDatabaseFile() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "boob", "boobs")
}

func createDirectoryIfNeeded(file string) {
	dir := filepath.Dir(file)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, DIR_PERMISSION)
	}
}

func serializeBookmark(bookmark *Bookmark) string {
	serializedTags := strings.Join(bookmark.Tags, ",")
	return bookmark.Url + " [" + serializedTags + "]"
}
