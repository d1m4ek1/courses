package handlers

import (
	api "courses/API"

	"github.com/jmoiron/sqlx"
)

func (s *Server) routesInitAPI(db *sqlx.DB) {
	apiroute := s.Router.Group("/api")
	{
		course := apiroute.Group("/course")
		{
			course.POST("/add", api.AddCourseController(db))

			course.POST("/buy", api.BuyCourseController(db))

			course.GET("/all", api.GetCourseAllController(db))

			course.GET("/get", api.GetCourseByIDController(db))

			course.PUT("/edit", api.EditCourseController(db))

			course.DELETE("/delete", api.DeleteCourseController(db))
		}

		cart := apiroute.Group("/cart")
		{
			cart.GET("/all", api.GetAllProductsController(db))

			cart.DELETE("/delete", api.DeleteFromCartController(db))
		}

		order := apiroute.Group("/order")
		{
			order.POST("/add", api.AddToOrderController(db))

			order.GET("/all", api.GetUserOrdersControllers(db))

			order.DELETE("/delete", api.DeleteOrderControllers(db))
		}
	}
}
