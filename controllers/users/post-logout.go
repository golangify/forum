package usercontroller

import (
	"forum/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *userController) postLogout(ctx *gin.Context) {
	v, _ := ctx.Get("session")
	session := v.(*models.Session)
	if session.User == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
		})
		return
	}
	c.middleware.SessionManager.DeleteSession(ctx)
}
