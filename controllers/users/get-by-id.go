package usercontroller

import (
	"forum/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (c *userController) getByID(ctx *gin.Context) {
	userID, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var user models.User
	if err := c.database.First(&user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.errorController.NotFound(ctx)
			return
		}
		ctx.Set("error", err.Error())
		c.errorController.InternalServerError(ctx)
		return
	}
	session, _ := c.middlewareController.SessionManager.GetSession(ctx)
	ctx.HTML(http.StatusOK, "users/profile", gin.H{
		"title":   "Профиль " + user.Username,
		"session": session,
		"user":    user,
	})
}
