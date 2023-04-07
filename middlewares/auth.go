package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		sessionID := session.Get("token")

		if sessionID == nil {
			ctx.Redirect(303, "/auth/login")
			ctx.Abort()
		}
	})
}
