package sectioncontroller

import (
	"fmt"
	"forum/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (c *sectionController) getByID(ctx *gin.Context) {
	sectionID, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var section models.Section
	err := c.database.Preload("User").First(&section, sectionID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.errorController.NotFound(ctx)
			return
		}
		ctx.Set("error", err.Error())
		c.errorController.InternalServerError(ctx)
		return
	}

	session, _ := c.middlewareController.SessionManager.GetSession(ctx)
	ctx.HTML(http.StatusOK, "sections/section", gin.H{
		"title":   fmt.Sprint("Раздел «", section.Title, "»"),
		"section": section,
		"session": session,
	})
}
