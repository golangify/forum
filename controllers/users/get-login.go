package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *userController) renderLoginPage(ctx *gin.Context, username string, err error) {
	var errorStr string
	if err != nil {
		errorStr = err.Error()
	}
	ctx.HTML(http.StatusOK, "users/login", gin.H{
		"title":    "Авторизация",
		"username": username,
		"error":    errorStr,
	})
}

func (c *userController) getLogin(ctx *gin.Context) {
	c.renderLoginPage(ctx, "", nil)
}
