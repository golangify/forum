package usercontroller

import (
	"errors"
	"forum/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type loginRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func (c *userController) postLogin(ctx *gin.Context) {
	var loginRequest loginRequest
	if err := ctx.ShouldBind(&loginRequest); err != nil {
		c.renderLoginPage(ctx, loginRequest.Username, err)
		return
	}

	if err := c.validateLoginForm(loginRequest); err != nil {
		c.renderLoginPage(ctx, loginRequest.Username, err)
		return
	}

	var user models.User
	if err := c.database.First(&user, "username = ?", loginRequest.Username).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.renderLoginPage(ctx, loginRequest.Username, errors.New("неправильное имя или пароль"))
			return
		}
		c.renderLoginPage(ctx, loginRequest.Username, err)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginRequest.Password)); err != nil {
		c.renderLoginPage(ctx, loginRequest.Username, errors.New("неправильное имя или пароль"))
		return
	}

	if err := c.middleware.SessionManager.BindSessionToUser(ctx, &user); err != nil {
		c.renderLoginPage(ctx, loginRequest.Username, err)
		return
	}

	ctx.Redirect(http.StatusFound, "/")
}
