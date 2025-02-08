package indexcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *indexController) get(ctx *gin.Context) {
	ctx.Redirect(http.StatusSeeOther, "/sections")
}
