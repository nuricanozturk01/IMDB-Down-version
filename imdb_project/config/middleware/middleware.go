package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
)

type AuthMiddleware struct {
	Store *sessions.CookieStore
}

func NewAuthMiddleware(store *sessions.CookieStore) *AuthMiddleware {
	return &AuthMiddleware{Store: store}
}

// Middleware function to check if user is authenticated (Handled by session)
func (m *AuthMiddleware) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session, _ := m.Store.Get(ctx.Request, "imdb-session")
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
