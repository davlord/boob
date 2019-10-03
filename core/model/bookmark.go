package model

import (
	"fmt"
	"strings"
)

type Bookmark struct {
	Url  string
	Tags []string
}

func (b Bookmark) String() string {
	tagsString := strings.Join(b.Tags, ",")
	return fmt.Sprintf("%s %s", b.Url, tagsString)
}
