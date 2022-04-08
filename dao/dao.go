package dao

import (
	"fileshare/models"

	"gorm.io/gorm"
)

type FileRepository interface {
	SaveFile(file models.File) models.File   
	GetFileGroup(id string) models.FileGroup 
	GetFile(id string) models.File   
}

type fileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) FileRepository {
	return &fileRepository{db: db}
}

func (f *fileRepository) SaveFile(file models.File) models.File {
	f.db.Create(&file)
	return file
}

func (f *fileRepository) GetFileGroup(id string) models.FileGroup {
	var files []models.File
	f.db.Table("files").Where("group_id = ?", id).Scan(&files)
	fileGroup := models.FileGroup{ID: id}
	fileGroup.Files = append(fileGroup.Files, files...)
	return fileGroup
}

func (f *fileRepository) GetFile(id string) models.File {
	var file models.File
	f.db.Table("files").Where("id = ?", id).Scan(&file)
	return file
}