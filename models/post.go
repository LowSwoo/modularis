package models

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

type Post struct {
	Id      string
	Title   string
	Author  string
	Content any
	err     string
}

func NewPost(title, author string, content any) *Post {
	hash := md5.Sum([]byte(time.Now().String() + title))
	return &Post{
		Title:   title,
		Author:  author,
		Content: content,
		Id:      hex.EncodeToString(hash[:]),
	}
}

func (p Post) Error() string {
	return p.err
}

var PostDoesntExist error = &Post{err: "post doesn't exist"}
var PostAlreadyExist error = &Post{err: "post already exist"}
