package sectioncontroller

import (
	"forum/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *sectionController) get(ctx *gin.Context) {
	v, _ := ctx.Get("session")
	session := v.(*models.Session)
	ctx.HTML(http.StatusOK, "sections/sections", gin.H{
		"title":   "Разделы",
		"session": session,
	})
}
