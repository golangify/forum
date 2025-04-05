package sectioncontroller

import (
	"forum/config"
	errorcontroller "forum/controllers/error"
	middlewarecontroller "forum/controllers/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type sectionController struct {
	config               *config.Config
	engine               *gin.Engine
	database             *gorm.DB
	middlewareController *middlewarecontroller.MiddlewareController
	errorController      *errorcontroller.ErrorController
}

func NewSectionController(config *config.Config, engine *gin.Engine, database *gorm.DB, middlewareController *middlewarecontroller.MiddlewareController, errorController *errorcontroller.ErrorController) *sectionController {
	c := &sectionController{
		config:               config,
		engine:               engine,
		database:             database,
		middlewareController: middlewareController,
		errorController:      errorController,
	}

	g := c.engine.Group("/sections")
	{
		g.GET("/", c.get)

		g.GET("/new", c.middlewareController.IfAuthorized, c.middlewareController.IfAdministrator, c.getNew)
		g.POST("/", c.middlewareController.IfAuthorized, c.post)
		g.GET("/:id", c.getByID)
	}

	return c
}
