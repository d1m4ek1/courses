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

			course.GET("/all", api.GetCourseAllController(db))

			course.GET("/byid", api.GetCourseByIDController(db))
		}
	}
}
