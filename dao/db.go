package dao

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"fileshare/models"
)

func NewDB() (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("fileshare.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.File{})
	return db
}

func CleanDB() (db *gorm.DB){
	db, err := gorm.Open(sqlite.Open("fileshare.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Unscoped().Where("state = ?", "1").Delete(&models.File{})
	return
}
