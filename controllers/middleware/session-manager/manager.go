package sessionmanager

import (
	"errors"
	"forum/config"
	"forum/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SessionManager struct {
	config   *config.Config
	engine   *gin.Engine
	database *gorm.DB
}

func NewSessionManager(config *config.Config, engine *gin.Engine, database *gorm.DB) *SessionManager {
	m := &SessionManager{
		config:   config,
		engine:   engine,
		database: database,
	}

	return m
}

func (m *SessionManager) GetSession(ctx *gin.Context) (*models.Session, error) {
	identificator, err := ctx.Cookie("session")
	if err != nil {
		if strings.Contains(err.Error(), "named cookie not present") {
			return nil, nil
		}
		return nil, err
	}

	var session models.Session
	if err = m.database.Preload("User").First(&session, "identificator = ?", identificator).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &session, nil
}

func (m *SessionManager) NewSession(ctx *gin.Context) (*models.Session, error) {
	session := &models.Session{
		Identificator: m.generateSessionIdentificator(),
		FirstIP:       ctx.ClientIP(),
		LastIP:        ctx.ClientIP(),
		ExpiresAt:     time.Now().Add(time.Duration(m.config.Session.DaysTimeout) * time.Hour * 24),
	}
	err := m.database.Create(session).Error
	if err != nil {
		return nil, err
	}
	m.setSession(ctx, session.Identificator, int(time.Duration(int(m.config.Session.DaysTimeout))*time.Hour*24))
	ctx.Set("session", session)
	return session, nil
}

func (m *SessionManager) BindSessionToUser(ctx *gin.Context, user *models.User) error {
	v, ok := ctx.Get("session")
	if !ok {
		return errors.New("ошибка биндинга сессии")
	}
	session := v.(*models.Session)
	session.UserID = user.ID
	session.User = user
	if err := m.database.Updates(session).Error; err != nil {
		return err
	}
	return nil
}

func (m *SessionManager) DeleteSession(ctx *gin.Context) error {
	v, _ := ctx.Get("session")
	session := v.(*models.Session)
	err := m.database.Delete(session).Error
	if err != nil {
		return err
	}
	ctx.SetCookie("session", "", -1, "/", "", false, true)
	return nil
}
