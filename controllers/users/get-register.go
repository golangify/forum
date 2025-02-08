package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *userController) renderRegisterPage(ctx *gin.Context, username string, err error) {
	var errorStr string
	if err != nil {
		errorStr = err.Error()
	}
	ctx.HTML(http.StatusOK, "users/register", gin.H{
		"title":    "Регистрация",
		"username": username,
		"error":    errorStr,
	})
}

func (c *userController) getRegister(ctx *gin.Context) {
	c.renderRegisterPage(ctx, "", nil)
}
