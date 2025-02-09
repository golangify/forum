package sectioncontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *sectionController) renderNewSectionPage(ctx *gin.Context, title string, body string, err error) {
	session, _ := c.middlewareController.SessionManager.GetSession(ctx)
	var errorStr string
	if err != nil {
		errorStr = err.Error()
	}
	ctx.HTML(http.StatusOK, "sections/new", gin.H{
		"title":         "Создание нового раздела",
		"section_title": title,
		"section_body":  body,
		"error":         errorStr,
		"session":       session,
	})
}

func (c *sectionController) getNew(ctx *gin.Context) {
	c.renderNewSectionPage(ctx, "", "", nil)
}
