package usercontroller

import (
	"errors"
	"forum/models"
	"gorm.io/gorm"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type registerRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func (c *userController) postRegister(ctx *gin.Context) {
	var registerRequest registerRequest
	err := ctx.ShouldBind(&registerRequest)
	if err != nil {
		c.renderRegisterPage(ctx, registerRequest.Username, err)
		return
	}

	err = c.validateRegisterForm(registerRequest)
	if err != nil {
		c.renderRegisterPage(ctx, registerRequest.Username, err)
		return
	}

	var registeredUser models.User
	if err = c.database.First(&registeredUser, "username = ?", registerRequest.Username).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			c.renderRegisterPage(ctx, registerRequest.Username, err)
			return
		}
	}
	if registeredUser.ID != 0 {
		c.renderRegisterPage(ctx, registerRequest.Username, errors.New("пользователь с таким именем уже существует"))
		return
	}

	hashedPassword, err := c.hashPassword(registerRequest.Password)
	if err != nil {
		c.renderRegisterPage(ctx, registerRequest.Username, err)
		return
	}

	registeredUser = models.User{
		Username:     registerRequest.Username,
		PasswordHash: hashedPassword,
	}
	err = c.database.Create(&registeredUser).Error
	if err != nil {
		c.renderRegisterPage(ctx, registerRequest.Username, err)
		return
	}

	err = c.middlewareController.SessionManager.BindSessionToUser(ctx, &registeredUser)
	if err != nil {
		c.renderRegisterPage(ctx, registerRequest.Username, err)
		return
	}

	ctx.Redirect(http.StatusFound, "/")
}

func (c *userController) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
