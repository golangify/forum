package usercontroller

import (
	"forum/config"
	middlewarecontroller "forum/controllers/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userController struct {
	config     *config.Config
	engine     *gin.Engine
	database   *gorm.DB
	middleware *middlewarecontroller.MiddlewareController
}

func NewUserController(config *config.Config, engine *gin.Engine, database *gorm.DB, middleware *middlewarecontroller.MiddlewareController) *userController {
	c := &userController{
		config:     config,
		engine:     engine,
		database:   database,
		middleware: middleware,
	}

	g := c.engine.Group("/users")
	{
		g.GET("/login", c.getLogin)
		g.POST("/login", c.postLogin)

		g.GET("/register", c.getRegister)
		g.POST("/register", c.postRegister)

		g.POST("/logout", c.postLogout)
	}

	return c
}
