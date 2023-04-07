package api

import (
	"courses/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func GetAllProductsController(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var data models.ProductData

		data.Token, _ = ctx.Cookie("user_token")

		response, status, err := data.GetAllProducts(db)
		if err != nil {
			fmt.Println(err.Error())
			ctx.JSON(status, gin.H{
				"successfully": false,
				"error":        err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"successfully": true,
			"data":         response,
		})
	})
}

func DeleteFromCartController(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var data models.ProductData
		var err error

		data.Id, err = strconv.ParseInt(ctx.Query("id"), 10, 32)
		if err != nil {
			fmt.Println(err.Error())

			ctx.JSON(http.StatusBadRequest, gin.H{
				"successfully": false,
				"error":        err.Error(),
			})
			return
		}

		if status, err := data.DeleteItemFromCart(db); err != nil {
			fmt.Println(err.Error())

			ctx.JSON(status, gin.H{
				"successfully": false,
				"error":        err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"successfully": true,
		})
	})
}
