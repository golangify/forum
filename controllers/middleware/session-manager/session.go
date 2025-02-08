package sessionmanager

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

var sessionIdentificatorSymbols = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-1234567890")

func (m *SessionManager) generateSessionIdentificator() string {
	b := make([]byte, m.config.Session.IdentificatorLength)
	for i := range b {
		b[i] = sessionIdentificatorSymbols[rand.Intn(len(sessionIdentificatorSymbols))]
	}
	return string(b)
}

func (c *SessionManager) setSession(ctx *gin.Context, sessionID string, maxAge int) {
	secure := ctx.Request.TLS != nil
	cookie := &http.Cookie{
		Name:     "session",
		Value:    sessionID,
		Path:     "/",
		Domain:   "",
		MaxAge:   maxAge,
		Secure:   secure,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(ctx.Writer, cookie)
}
