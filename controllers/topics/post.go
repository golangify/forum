package topiccontroller

import (
	"fmt"
	"forum/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type newTopicRequest struct {
	SectionID uint   `form:"section_id"`
	Title     string `form:"title"`
	Body      string `form:"body"`
}

func (c *topicController) post(ctx *gin.Context) {
	session, _ := c.middlewareController.SessionManager.GetSession(ctx)
	var newTopicRequest newTopicRequest
	if err := ctx.ShouldBind(&newTopicRequest); err != nil {
		c.renderNewTopicPage(ctx, newTopicRequest.Title, newTopicRequest.Body, err)
		return
	}

	var section models.Section
	if err := c.database.First(&section, newTopicRequest.SectionID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.Set("error", fmt.Sprint("невозможно создать тему для раздела с id ", newTopicRequest.SectionID, ", т.к. раздела с таким id нет. Возможно раздел был только что удалён"))
			c.errorController.NotFound(ctx)
			return
		}
		ctx.Set("error", err.Error())
		c.errorController.InternalServerError(ctx)
		return
	}

	if err := c.validateNewTopicForm(newTopicRequest); err != nil {
		c.renderNewTopicPage(ctx, newTopicRequest.Title, newTopicRequest.Body, err)
		return
	}

	newTopic := models.Topic{
		SectionID: section.ID,
		Section:   &section,
		UserID:    session.UserID,
		User:      session.User,
		Title:     newTopicRequest.Title,
		Body:      newTopicRequest.Body,
	}

	if err := c.database.Create(&newTopic).Error; err != nil {
		c.renderNewTopicPage(ctx, newTopicRequest.Title, newTopicRequest.Body, err)
		return
	}

	ctx.Redirect(http.StatusFound, fmt.Sprint("/sections/", newTopic.ID))
}
