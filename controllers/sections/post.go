package sectioncontroller

import (
	"fmt"
	"forum/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type newSectionRequest struct {
	Title string `form:"title"`
	Body  string `form:"body"`
}

func (c *sectionController) post(ctx *gin.Context) {
	v, _ := ctx.Get("session")
	session := v.(*models.Session)
	var newSectionRequest newSectionRequest
	if err := ctx.ShouldBind(&newSectionRequest); err != nil {
		c.renderNewSectionPage(ctx, newSectionRequest.Title, newSectionRequest.Body, err)
		return
	}

	if err := c.validateNewSectionForm(newSectionRequest); err != nil {
		c.renderNewSectionPage(ctx, newSectionRequest.Title, newSectionRequest.Body, err)
		return
	}

	newSection := models.Section{
		UserID: session.UserID,
		User:   session.User,
		Title:  newSectionRequest.Title,
		Body:   newSectionRequest.Body,
	}

	if err := c.database.Create(&newSection).Error; err != nil {
		c.renderNewSectionPage(ctx, newSectionRequest.Title, newSectionRequest.Body, err)
		return
	}

	ctx.Redirect(http.StatusFound, fmt.Sprint("/sections/", newSection.ID))
}
