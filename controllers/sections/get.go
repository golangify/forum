package sectioncontroller

import (
	"forum/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *sectionController) get(ctx *gin.Context) {
	session, _ := c.middlewareController.SessionManager.GetSession(ctx)

	var sections []models.Section
	err := c.database.Order("updated_at DESC").Preload("User").Find(&sections).Error
	if err != nil {
		ctx.Set("error", err.Error())
		c.errorController.InternalServerError(ctx)
		return
	}

	ctx.HTML(http.StatusOK, "sections/sections", gin.H{
		"title":    "Разделы",
		"sections": sections,
		"session":  session,
	})
}
