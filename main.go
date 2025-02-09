package main

import (
	"forum/config"
	errorcontroller "forum/controllers/error"
	indexcontroller "forum/controllers/index"
	middlewarecontroller "forum/controllers/middleware"
	sectioncontroller "forum/controllers/sections"
	topiccontroller "forum/controllers/topics"
	usercontroller "forum/controllers/users"
	"forum/models"
	templatesutils "forum/utils/templates"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	engine := gin.Default()

	config, err := config.LoadFromJsonFile("config.json")
	if err != nil {
		log.Fatalln(err)
	}

	database, err := gorm.Open(sqlite.Open(config.Database), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	if err = database.AutoMigrate(&models.User{}, &models.Session{}, &models.Section{}); err != nil {
		log.Fatalln(err)
	}

	err = templatesutils.Init(config, engine)
	if err != nil {
		log.Fatalln(err)
	}

	engine.Static("/scripts", "./scripts")
	engine.Static("/styles", "./styles")
	engine.Static("/assets", "./assets")

	middlewareController := middlewarecontroller.NewMiddlewareController(config, engine, database)
	engine.Use(middlewareController.Identificate)

	errorController := errorcontroller.NewErrorController(config, engine, database, middlewareController)
	engine.NoRoute(errorController.NotFound)

	indexcontroller.NewIndexController(config, engine, database, middlewareController, errorController)
	sectioncontroller.NewSectionController(config, engine, database, middlewareController, errorController)
	topiccontroller.NewTopicController(config, engine, database, middlewareController, errorController)
	usercontroller.NewUserController(config, engine, database, middlewareController, errorController)

	engine.Run(":8080")
}
