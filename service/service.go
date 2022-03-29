package service

import (
	"bytes"
	"encoding/base64"
	"fileshare/dao"
	"fileshare/models"
	"fileshare/utils"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"time"
)

type fileService struct {
	FileRepository dao.FileRepository
}

type FileService interface {
	UploadFile(fileGroupID string, file multipart.File, header *multipart.FileHeader) models.File
	ViewFileGroup(id string) models.FileGroup
	DownloadFile(id string) models.File
}

func NewFileService(FileRepository dao.FileRepository) FileService {
	return fileService{
		FileRepository: FileRepository,
	}
}

func (f fileService) UploadFile(fileGroupID string, file multipart.File, header *multipart.FileHeader) models.File{
	id := utils.GenerateID(6)
	var Buf bytes.Buffer
	io.Copy(&Buf, file)
	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		if err := os.Mkdir("uploads", os.ModePerm); err != nil {
			msg := "Could not create uploads directory"
			log.Fatal(msg,err)
		}
	}
	err := ioutil.WriteFile("uploads/"+id, Buf.Bytes(), os.ModePerm)
	if err != nil{
		msg := "Unable to save file"
		log.Fatal(msg,err)
	}
	Buf.Reset()
	fileData := models.File{
		ID:         id,
		GroupId:      fileGroupID,
		Name:       header.Filename,
		Size:       header.Size,
		MediaType:  header.Header.Get("Content-Type"),
		UploadDate: time.Now(),
		URL:        utils.CreateRoute(id, "download"),     
	}
	log.Println("Uploaded file:", fileData)
	return f.FileRepository.SaveFile(fileData)
}

func (f fileService) ViewFileGroup(id string) models.FileGroup {
	fileGroup := f.FileRepository.GetFileGroup(id)
	fileGroup.QR = base64.StdEncoding.EncodeToString(utils.CreateQR(utils.CreateRoute(fileGroup.ID, "view")))
	return fileGroup
}

func (f fileService) DownloadFile(id string) models.File {
	file := f.FileRepository.GetFile(id)
	return file
}