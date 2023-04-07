package middlewares

import (
	generatetoken "courses/pkg/generateToken"
	"fmt"

	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func CreateConfigCSRFDefenderMiddleware() csrf.Options {
	secretKey, err := generatetoken.CreateToken(256)
	if err != nil {
		fmt.Println(err.Error())
	}

	if secretKey == "" {
		secretKey = "secret-csrf-dfgsn53434532-dsfgsdvoljj54bn6-vzxc787834jm2mcv"
	}

	return csrf.Options{
		Secret: secretKey,
		ErrorFunc: func(ctx *gin.Context) {
			ctx.String(400, "CSRF token mismatch")
			ctx.Abort()
		},
	}
}
