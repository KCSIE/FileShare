package dao

import "fileshare/models"

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB() (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("fileshare.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.File{})
	return db
}