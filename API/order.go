package api

import (
	"courses/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func AddToOrderController(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		token, _ := ctx.Cookie("user_token")
		order := models.DataForOrder{Token: token}

		if err := ctx.ShouldBindJSON(&order); err != nil {
			fmt.Println(err.Error())

			ctx.JSON(http.StatusBadRequest, gin.H{
				"successfully": false,
				"error":        err.Error(),
			})
			return
		}

		if status, err := order.AddToOrder(db); err != nil {
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

func GetUserOrdersControllers(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		token, _ := ctx.Cookie("user_token")
		order := models.DataForOrder{Token: token}

		response, status, err := order.GetUserOrders(db)
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

func DeleteOrderControllers(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		orderId, err := strconv.ParseInt(ctx.Query("orderId"), 10, 32)
		if err != nil {
			fmt.Println(err.Error())

			ctx.JSON(http.StatusBadRequest, gin.H{
				"successfully": false,
				"error":        err.Error(),
			})
			return
		}

		order := models.DataForOrder{Id: orderId}

		if status, err := order.DeleteOrder(db); err != nil {
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
