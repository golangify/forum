package usercontroller

import (
	"forum/config"
	errorcontroller "forum/controllers/error"
	middlewarecontroller "forum/controllers/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userController struct {
	config               *config.Config
	engine               *gin.Engine
	database             *gorm.DB
	middlewareController *middlewarecontroller.MiddlewareController
	errorController      *errorcontroller.ErrorController
}

func NewUserController(config *config.Config, engine *gin.Engine, database *gorm.DB, middlewareController *middlewarecontroller.MiddlewareController, errorController *errorcontroller.ErrorController) *userController {
	c := &userController{
		config:               config,
		engine:               engine,
		database:             database,
		middlewareController: middlewareController,
		errorController:      errorController,
	}

	g := c.engine.Group("/users")
	{
		g.GET("/login", c.getLogin)
		g.POST("/login", c.postLogin)

		g.GET("/:id", c.getByID)

		g.GET("/register", c.getRegister)
		g.POST("/register", c.postRegister)

		g.POST("/logout", c.postLogout)
	}

	return c
}
