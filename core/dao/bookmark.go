package dao

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	. "github.com/davlord/boob/core/model"
)

const (
	FILE_MODE_WRITE = os.O_APPEND | os.O_CREATE | os.O_WRONLY
	FILE_MODE_READ  = os.O_RDONLY
	FILE_PERMISSION = 0644
	DIR_PERMISSION  = 0755
)

func CreateBookmark(bookmark *Bookmark) error {

	f, err := openDatabaseFile(FILE_MODE_WRITE)
	if err != nil {
		return err
	}
	defer f.Close()

	existingBookmark, err := findBookmarkByUrl(bookmark.Url)
	if err != nil {
		return err
	}
	if existingBookmark != nil {
		return errors.New("a bookmark with the same URL already exists")
	}

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

func unserializeBookmark(serializedBookmark string) *Bookmark {
	var lineRegexp = regexp.MustCompile(`([^\ ]+)\ \[([^\]]*)\]`)
	var parts []string = lineRegexp.FindStringSubmatch(serializedBookmark)
	return &Bookmark{
		Url:  parts[1],
		Tags: strings.Split(parts[2], ","),
	}
}

func findBookmarkByUrl(url string) (*Bookmark, error) {
	f, err := openDatabaseFile(FILE_MODE_READ)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var foundBookmark *Bookmark = nil

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		serializedBookmark := scanner.Text()
		if unserializedBookmark := unserializeBookmark(serializedBookmark); unserializedBookmark.Url == url {
			foundBookmark = unserializedBookmark
			break
		}
	}

	return foundBookmark, nil
}
