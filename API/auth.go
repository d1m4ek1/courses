package api

import (
	"courses/models"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	csrf "github.com/utrack/gin-csrf"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

var (
	errBadCredentials = errors.New("email or password is incorrect")

	jwtSecretKey = []byte("very-secret-key")
)

func setSessionHasLoginToken(userAuth *models.UserAuthData, db *sqlx.DB, ctx *gin.Context) (bool, bool, int, error) {
	isAuth, status, err := userAuth.Login(db)
	if err != nil {
		return false, false, status, err
	}

	if isAuth {
		session := sessions.Default(ctx)

		session.Options(sessions.Options{MaxAge: 60 * 60 * 24, HttpOnly: true, Secure: true, Path: "/", Domain: "localhost:3000"})
		session.Set("token", userAuth.Token)

		if err := session.Save(); err != nil {
			return false, false, http.StatusInternalServerError, err
		}

		return true, true, http.StatusOK, err
	}

	return true, false, http.StatusOK, err
}

func LoginToAccount(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var userAuth models.UserAuthData

		if err := ctx.ShouldBindJSON(&userAuth); err != nil {
			fmt.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"successfully": false,
				"error":        err.Error(),
			})
			return
		}

		successfully, isLogin, status, err := setSessionHasLoginToken(&userAuth, db, ctx)
		if err != nil {
			fmt.Println(err.Error())
			ctx.JSON(status, gin.H{
				"successfully": false,
				"error":        err.Error(),
				"isLogin":      isLogin,
			})
			return
		}

		ctx.SetCookie("user_token", userAuth.Token, 60*60*24, "/", "localhost:3000", true, true)

		ctx.JSON(status, gin.H{
			"successfully": successfully,
			"isLogin":      isLogin,
		})
	})
}

func CreateAccountAndLoginToAccount(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var userAuth models.UserAuthData

		if err := ctx.ShouldBindJSON(&userAuth); err != nil {
			fmt.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"successfully": false,
				"error":        err.Error(),
			})
			return
		}

		status, err := userAuth.CreateUser(db)
		if err != nil {
			fmt.Println(err.Error())

			ctx.JSON(status, gin.H{
				"successfully": false,
				"error":        err.Error(),
				"isLogin":      false,
			})
			return
		}

		if status == http.StatusOK {
			successfully, isLogin, status, err := setSessionHasLoginToken(&userAuth, db, ctx)
			if err != nil {
				fmt.Println(err.Error())
				ctx.JSON(status, gin.H{
					"successfully": false,
					"error":        err.Error(),
					"isLogin":      isLogin,
				})
				return
			}

			ctx.SetCookie("user_token", userAuth.Token, 60*60*24, "/", "localhost:3000", true, true)

			ctx.JSON(status, gin.H{
				"successfully": successfully,
				"isLogin":      isLogin,
			})

			return
		}

		ctx.JSON(status, gin.H{
			"successfully": true,
			"isLogin":      false,
		})
	})
}

func LogoutFromAccount(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var userAuth models.UserAuthData
		var err error

		userAuth.Token, err = ctx.Cookie("user_token")
		if err != nil {
			fmt.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"successfully": false,
				"error":        err.Error(),
			})
			return
		}

		session := sessions.Default(ctx)

		session.Clear()

		if err := session.Save(); err != nil {
			fmt.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"successfully": false,
				"error":        err.Error(),
			})
			return
		}

		if status, err := userAuth.Logout(db); err != nil {
			fmt.Println(err.Error())
			ctx.JSON(status, gin.H{
				"successfully": false,
				"error":        err.Error(),
			})
			return
		}

		ctx.SetCookie("user_token", "", -1, "/", "localhost:3000", true, true)

		ctx.JSON(http.StatusOK, gin.H{
			"successfully": true,
		})
	})
}

func CheckUserAuth(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		ctx.Header("X-CSRF-TOKEN", csrf.GetToken(ctx))

		if session.Get("token") == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"successfully": true,
				"isAuth":       false,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"successfully": true,
				"isAuth":       true,
			})
		}
	})
}
