package controller

import (
	"fileshare/models"
	"fileshare/service"
	"fileshare/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

type FileController struct {
	fileService service.FileService
}

func NewFileController(fileService service.FileService) FileController {
	return FileController{fileService: fileService}
}

func (h *FileController)HandleUpload(c *gin.Context){
	fileGroup := models.FileGroup{
		ID: utils.GenerateID(6),
	}
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 64<<20) // limit for upload
	err := c.Request.ParseMultipartForm(64<<20) // limit for memory
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	form, _ := c.MultipartForm()
	fileHeaders := form.File["files"]
	for _, header := range fileHeaders {
		file,err := header.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		defer file.Close()
		fileData := h.fileService.UploadFile(fileGroup.ID, file, header) // filegroup id, file itself, file info, return file metadata
		fileGroup.Files = append(fileGroup.Files, fileData)
	}
	c.Redirect(http.StatusFound,utils.CreateRoute(fileGroup.ID,"view"))
}

func (h *FileController)HandleView(c *gin.Context){
	fileGroup := h.fileService.ViewFileGroup(c.Param("id"))
	c.HTML(http.StatusOK, "download/index.tmpl", fileGroup)
}

func (h *FileController)HandleDownload(c *gin.Context){
	file := h.fileService.DownloadFile(c.Param("id"))
	c.Header("Content-Disposition", "attachment; filename="+file.Name)
	c.Header("Content-Type", file.MediaType+"; charset=utf-8")
	c.File("uploads/"+file.ID)
}