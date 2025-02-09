package errorcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *ErrorController) NotFound(ctx *gin.Context) {
	session, _ := c.middlewareController.SessionManager.GetSession(ctx)

	var errStr = "запрашиваемая страница не найдена"
	if v, ok := ctx.Get("error"); ok {
		errStr = v.(string)
	}

	ctx.HTML(http.StatusNotFound, "error/error", gin.H{
		"title":   "Страница не найдена",
		"code":    http.StatusNotFound,
		"error":   errStr,
		"session": session,
	})
}
