package indexcontroller

import (
	"forum/config"
	errorcontroller "forum/controllers/error"
	middlewarecontroller "forum/controllers/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type indexController struct {
	config               *config.Config
	engine               *gin.Engine
	database             *gorm.DB
	middlewareController *middlewarecontroller.MiddlewareController
	errorController      *errorcontroller.ErrorController
}

func NewIndexController(config *config.Config, engine *gin.Engine, database *gorm.DB, middlewareController *middlewarecontroller.MiddlewareController, errorController *errorcontroller.ErrorController) *indexController {
	c := &indexController{
		engine: engine,
	}

	engine.GET("/", c.get)

	return c
}
