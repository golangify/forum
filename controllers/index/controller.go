package indexcontroller

import (
	"github.com/gin-gonic/gin"
)

type indexController struct {
	engine *gin.Engine
}

func NewIndexController(engine *gin.Engine) *indexController {
	c := &indexController{
		engine: engine,
	}

	engine.GET("/", c.get)

	return c
}
