package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Bookmark struct {
	gorm.Model
	Name string `json:"name"`
	Url  string `json:"url"`
}

func InitDatabase() error {
	db, err := gorm.Open(sqlite.Open("bookmark.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&Bookmark{})

	return nil
}
