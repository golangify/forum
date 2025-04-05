package middlewarecontroller

import (
	"forum/config"
	sessionmanager "forum/controllers/middleware/session-manager"
	"html/template"
	"net/http"
	"slices"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MiddlewareController struct {
	config         *config.Config
	engine         *gin.Engine
	database       *gorm.DB
	SessionManager *sessionmanager.SessionManager
}

func NewMiddlewareController(config *config.Config, engine *gin.Engine, database *gorm.DB) *MiddlewareController {
	return &MiddlewareController{
		config:         config,
		engine:         engine,
		database:       database,
		SessionManager: sessionmanager.NewSessionManager(config, engine, database),
	}
}

func (c *MiddlewareController) Identificate(ctx *gin.Context) {
	session, err := c.SessionManager.GetSession(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if session == nil {
		session, err = c.SessionManager.NewSession(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	if time.Now().After(session.ExpiresAt) {
		if err = c.database.Delete(session).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		session, err = c.SessionManager.NewSession(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	clientIP := ctx.ClientIP()
	if session.LastIP != clientIP {
		session.LastIP = clientIP
		if err = c.database.Updates(session).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	ctx.Set("session", session)
}

func (c *MiddlewareController) IfAuthorized(ctx *gin.Context) {
	session, err := c.SessionManager.GetSession(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	if session == nil || session.User == nil {
		ctx.HTML(http.StatusForbidden, "error/error", gin.H{
			"title":   "Недостаточно прав",
			"code":    http.StatusForbidden,
			"error":   template.HTML("необходима <a href='/users/login'>авторизация</a> для доступа к этой странице"),
			"session": session,
		})
		ctx.Abort()
	}
}

func (c *MiddlewareController) IfAdministrator(ctx *gin.Context) {
	session, err := c.SessionManager.GetSession(ctx)
	if err != nil {
		ctx.HTML(http.StatusForbidden, "error/error", gin.H{
			"title":   "Недостаточно прав",
			"code":    http.StatusForbidden,
			"error":   template.HTML("необходима <a href='/users/login'>авторизация</a> для доступа к этой странице"),
			"session": session,
		})
		ctx.Abort()
		return
	}
	if !slices.Contains(session.User.Roles, "admin") {
		ctx.HTML(http.StatusForbidden, "error/error", gin.H{
			"title":   "Недостаточно прав",
			"code":    http.StatusForbidden,
			"error":   template.HTML("необходим админ-аккаунт для доступа к этой странице"),
			"session": session,
		})
		ctx.Abort()
		return
	}
}
