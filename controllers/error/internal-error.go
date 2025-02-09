package errorcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *ErrorController) InternalServerError(ctx *gin.Context) {
	session, _ := c.middlewareController.SessionManager.GetSession(ctx)
	var errStr string
	err, ok := ctx.Get("error")
	if ok {
		errStr = err.(string)
	}
	ctx.HTML(http.StatusInternalServerError, "error/error", gin.H{
		"title":   "Внутренняя ошибка сервера",
		"code":    http.StatusInternalServerError,
		"error":   "Внутренняя ошибка сервера: " + errStr,
		"session": session,
	})
}
