package handlers

import (
	api "courses/API"
	"courses/middlewares"

	"github.com/jmoiron/sqlx"
)

func (s *Server) routesInitAPI(db *sqlx.DB) {
	apiroute := s.Router.Group("/api")
	{
		courseAuth := apiroute.Group("/course")
		courseAuth.Use(middlewares.Auth())
		{
			courseAuth.POST("/add", api.AddCourseController(db))

			courseAuth.POST("/buy", api.BuyCourseController(db))

			courseAuth.PUT("/edit", api.EditCourseController(db))

			courseAuth.DELETE("/delete", api.DeleteCourseController(db))
		}

		course := apiroute.Group("/course")
		{
			course.GET("/all", api.GetCourseAllController(db))

			course.GET("/get", api.GetCourseByIDController(db))
		}

		cart := apiroute.Group("/cart")
		cart.Use(middlewares.Auth())
		{
			cart.GET("/all", api.GetAllProductsController(db))

			cart.DELETE("/delete", api.DeleteFromCartController(db))
		}

		order := apiroute.Group("/order")
		order.Use(middlewares.Auth())
		{
			order.POST("/add", api.AddToOrderController(db))

			order.GET("/all", api.GetUserOrdersControllers(db))

			order.DELETE("/delete", api.DeleteOrderControllers(db))
		}

		auth := apiroute.Group("/auth")
		{
			auth.POST("/login", api.LoginToAccount(db))

			auth.POST("/register", api.CreateAccountAndLoginToAccount(db))

			auth.POST("/logout", api.LogoutFromAccount(db))

			auth.GET("/check", api.CheckUserAuth(db))
		}
	}
}
