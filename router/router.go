package router

import (
	"fileshare/controller"
	"github.com/gin-gonic/gin"
)

type Router struct {
	fileHandler controller.FileController
}

func (r *Router) With(engine *gin.Engine) {
	v1 := engine.Group("v1")
	{
		v1.POST("/upload", r.fileHandler.HandleUpload)
		v1.GET("/view/:id", r.fileHandler.HandleView)
		v1.GET("/download/:id", r.fileHandler.HandleDownload)
	}
}

func NewRouter(fileHandler controller.FileController) *Router {
	return &Router{
		fileHandler: fileHandler,
	}
}
