package topiccontroller

import (
	"fmt"
	"forum/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (c *topicController) renderNewTopicPage(ctx *gin.Context, title string, body string, err error) {
	session, _ := c.middlewareController.SessionManager.GetSession(ctx)
	var errorStr string
	if err != nil {
		errorStr = err.Error()
	}

	sectionID, _ := strconv.ParseUint(ctx.Query("section_id"), 10, 32)
	var section models.Section
	if err = c.database.First(&section, sectionID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.Set("error", fmt.Sprint("невозможно создать тему для раздела с id ", sectionID, ", т.к. раздела с таким id нет"))
			c.errorController.NotFound(ctx)
			return
		}
	}

	ctx.HTML(http.StatusOK, "topics/new", gin.H{
		"title":       fmt.Sprint("Создание новой темы в разделе «", section.Title, "»"),
		"section":     section,
		"topic_title": title,
		"topic_body":  body,
		"error":       errorStr,
		"session":     session,
	})
}

func (c *topicController) getNew(ctx *gin.Context) {
	c.renderNewTopicPage(ctx, "", "", nil)
}
