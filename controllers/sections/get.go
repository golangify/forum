package sectioncontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *sectionController) get(ctx *gin.Context) {
	session, _ := c.middlewareController.SessionManager.GetSession(ctx)
	ctx.HTML(http.StatusOK, "sections/sections", gin.H{
		"title":   "Разделы",
		"session": session,
	})
}
