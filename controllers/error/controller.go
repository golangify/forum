package errorcontroller

import (
	"forum/config"
	middlewarecontroller "forum/controllers/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ErrorController struct {
	config               *config.Config
	engine               *gin.Engine
	database             *gorm.DB
	middlewareController *middlewarecontroller.MiddlewareController
}

func NewErrorController(config *config.Config, engine *gin.Engine, database *gorm.DB, middlewareController *middlewarecontroller.MiddlewareController) *ErrorController {
	c := &ErrorController{
		config:               config,
		engine:               engine,
		database:             database,
		middlewareController: middlewareController,
	}

	g := c.engine.Group("/error")
	{
		g.GET("/not-found", c.NotFound)
	}

	return c
}
