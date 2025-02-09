package topiccontroller

import (
	"forum/config"
	errorcontroller "forum/controllers/error"
	middlewarecontroller "forum/controllers/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type topicController struct {
	config               *config.Config
	engine               *gin.Engine
	database             *gorm.DB
	middlewareController *middlewarecontroller.MiddlewareController
	errorController      *errorcontroller.ErrorController
}

func NewTopicController(config *config.Config, engine *gin.Engine, database *gorm.DB, middlewareController *middlewarecontroller.MiddlewareController, errorController *errorcontroller.ErrorController) *topicController {
	c := &topicController{
		config:               config,
		engine:               engine,
		database:             database,
		middlewareController: middlewareController,
		errorController:      errorController,
	}

	g := c.engine.Group("/topics")
	{
		g.GET("/new", c.middlewareController.IfAuthorized, c.getNew)
		g.POST("/", c.middlewareController.IfAuthorized, c.post)
	}

	return c
}
