package sectioncontroller

import (
	"forum/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type sectionController struct {
	config   *config.Config
	engine   *gin.Engine
	database *gorm.DB
}

func NewSectionController(config *config.Config, engine *gin.Engine, database *gorm.DB) *sectionController {
	c := &sectionController{
		config:   config,
		engine:   engine,
		database: database,
	}

	g := c.engine.Group("/sections")
	{
		g.GET("/", c.get)

		g.GET("/new", c.getNew)
		g.POST("/", c.post)
	}

	return c
}
