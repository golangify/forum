package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *userController) postLogout(ctx *gin.Context) {
	session, err := c.middlewareController.SessionManager.GetSession(ctx)
	if err != nil {
		ctx.Set("error", err.Error())
		c.errorController.InternalServerError(ctx)
		return
	}
	if session.User == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
		})
		return
	}
	if err = c.middlewareController.SessionManager.DeleteSession(ctx); err != nil {
		ctx.Set("error", err.Error())
		c.errorController.InternalServerError(ctx)
		return
	}
	if _, err = c.middlewareController.SessionManager.NewSession(ctx); err != nil {
		ctx.Set("error", err.Error())
		c.errorController.InternalServerError(ctx)
		return
	}
}
