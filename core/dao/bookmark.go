package dao

import (
	"bufio"
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

	serializedBookmark := serializeBookmark(bookmark)
	if _, err = f.WriteString(serializedBookmark + "\n"); err != nil {
		return err
	}

	return nil
}

func FindBookmarkByUrl(url string) (*Bookmark, error) {
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

func FindBookmarkByIndex(index int) (*Bookmark, error) {
	f, err := openDatabaseFile(FILE_MODE_READ)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var i int = 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		serializedBookmark := scanner.Text()
		if i == index {
			bookmark := unserializeBookmark(serializedBookmark)
			return bookmark, nil
		}
		i++
	}

	return nil, nil
}

func GetAllBookmarks() ([]*Bookmark, error) {
	f, err := openDatabaseFile(FILE_MODE_READ)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var bookmarks []*Bookmark = make([]*Bookmark, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		serializedBookmark := scanner.Text()
		var bookmark *Bookmark = unserializeBookmark(serializedBookmark)
		bookmarks = append(bookmarks, bookmark)
	}

	return bookmarks, nil
}

func GetAllTags() ([]string, error) {
	// get all bookmarks
	bookmarks, err := GetAllBookmarks()
	if err != nil {
		return nil, err
	}

	// extract unique tags from all bookmarks
	tagSet := make(map[string]struct{})
	for _, bm := range bookmarks {
		for _, tag := range bm.Tags {
			tagSet[tag] = struct{}{}
		}
	}

	// build tags slice as a result
	tags := make([]string, 0, len(tagSet))
	for tag, _ := range tagSet {
		tags = append(tags, tag)
	}
	return tags, nil
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
