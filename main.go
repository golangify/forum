package main

import (
	"forum/config"
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

	indexcontroller.NewIndexController(engine)
	sectioncontroller.NewSectionController(config, engine, database)
	topiccontroller.NewTopicController(config, engine, database)
	usercontroller.NewUserController(config, engine, database, middlewareController)

	engine.Run(":8080")
}
