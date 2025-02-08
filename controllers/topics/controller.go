package topiccontroller

import (
	"forum/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type topicController struct {
	config   *config.Config
	engine   *gin.Engine
	database *gorm.DB
}

func NewTopicController(config *config.Config, engine *gin.Engine, database *gorm.DB) *topicController {
	c := &topicController{
		config:   config,
		engine:   engine,
		database: database,
	}

	return c
}
