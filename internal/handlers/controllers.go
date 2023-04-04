package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) home() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Главная",
		})
	})
}

func (s *Server) courses() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Курсы",
		})
	})
}

func (s *Server) order() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Заказы",
		})
	})
}

func (s *Server) cart() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Корзина",
		})
	})
}

func (s *Server) addOrder() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Корзина",
		})
	})
}

func (s *Server) previewCourse() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Курс",
		})
	})
}
