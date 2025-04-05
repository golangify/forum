package topiccontroller

import (
	"forum/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (c *topicController) getByID(ctx *gin.Context) {
	topicID, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var topic models.Topic
	if err := c.database.Preload("Section").First(&topic, topicID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.errorController.NotFound(ctx)
			return
		}
		ctx.Set("error", err.Error())
		c.errorController.InternalServerError(ctx)
		return
	}

	session, _ := c.middlewareController.SessionManager.GetSession(ctx)
	ctx.HTML(http.StatusOK, "topics/topic", gin.H{
		"title":   topic.Title,
		"section": topic.Section,
		"topic":   topic,
		"session": session,
	})
}
