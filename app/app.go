package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"fileshare/controller"
	"fileshare/dao"
	"fileshare/router"
	"fileshare/service"
)

type Server struct {
	engine    *gin.Engine
	apiRouter *router.Router
}

func (s *Server) Start() {
	s.apiRouter.With(s.engine)
	err := s.engine.Run(fmt.Sprintf(":%d", 8080))
	if err != nil {
		panic(err)
	}
}

func NewServer(engine *gin.Engine, apiRouter *router.Router) *Server {
	return &Server{
		engine:    engine,
		apiRouter: apiRouter,
	}
}

func NewGinEngine() *gin.Engine {
	router := gin.Default()
	router.Static("/pubilc/static", "./public/static")
	router.LoadHTMLGlob("public/templates/**/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload/index.tmpl", nil)
	})
	return router
}

// dependency injection
func InitServer() *Server {
	engine := NewGinEngine()
	db := dao.NewDB()
	fileRepository := dao.NewFileRepository(db)
	fileService := service.NewFileService(fileRepository)
	fileController := controller.NewFileController(fileService)
	routerRouter := router.NewRouter(fileController)
	server := NewServer(engine, routerRouter)
	return server
}

