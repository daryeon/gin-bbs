package models

import (
	"github.com/jinzhu/gorm"
)

/*
	帖子

///
*/
type Post struct {
	gorm.Model
	Title    string
	Content  string
	Author   string
	CreateAt string `json:"create_at" binding:"required"`
}

var Db *gorm.DB

func InitDb() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "forum.db?_loc=Asia/Shanghai")
	if err == nil {
		Db = db
		Db.AutoMigrate(&Post{})
		return db, err
	}
	return nil, err
}

func GetPosts() ([]Post, error) {
	var posts []Post
	err := Db.Find(&posts).Error
	return posts, err
}

func AddPost(post Post) (bool, error) {
	if post.Title == "" {
		return false, nil
	}
	err := Db.Create(&post).Error
	return true, err
}
