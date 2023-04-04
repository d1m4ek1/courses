package api

import (
	"courses/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func AddCourseController(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var data models.CourseData

		if err := json.NewDecoder(ctx.Request.Body).Decode(&data); err != nil {
			fmt.Println(err.Error())
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"successfully": false,
				"error":        err.Error(),
			})
			return
		}

		if status, err := data.AddCourse(db); err != nil {
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

func GetCourseAllController(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		response, status, err := models.GetAllCourses(db)
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

func GetCourseByIDController(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var data models.CourseData

		id, err := strconv.ParseInt(ctx.Query("id"), 10, 32)
		if err != nil {
			fmt.Println(err.Error())

			ctx.JSON(http.StatusBadRequest, gin.H{
				"successfully": false,
				"error":        err.Error(),
			})
			return
		}

		data.Id = id

		status, err := data.GetByIDCourse(db)
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
			"data":         data,
		})
	})
}
