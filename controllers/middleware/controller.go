package middlewarecontroller

import (
	"forum/config"
	sessionmanager "forum/controllers/middleware/session-manager"
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
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	if session == nil {
		session, err = c.SessionManager.NewSession(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}
	}
	if time.Now().After(session.ExpiresAt) {
		if err = c.database.Delete(session).Error; err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}
		session, err = c.SessionManager.NewSession(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}
	}
	clientIP := ctx.ClientIP()
	if session.LastIP != clientIP {
		session.LastIP = clientIP
		if err = c.database.Updates(session).Error; err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}
	}
	ctx.Set("session", session)
}
