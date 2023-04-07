package handlers

import "courses/middlewares"

func (s *Server) routesInit() {
	home := s.Router.Group("/")
	{
		home.GET("", s.home())
	}

	courses := s.Router.Group("/courses")
	{
		courses.GET("", s.courses())

		courses.Use(middlewares.Auth()).GET("/add", s.addCourse())

		courses.Use(middlewares.Auth()).GET("/:id/edit", s.editCourse())

		courses.GET("/:id", s.previewCourse())
	}

	order := s.Router.Group("/orders")
	order.Use(middlewares.Auth())
	{
		order.GET("", s.order())
	}

	cart := s.Router.Group("/cart")
	cart.Use(middlewares.Auth())
	{
		cart.GET("", s.cart())
	}

	login := s.Router.Group("/auth")
	{
		login.GET("/login", s.login())

		login.GET("/register", s.register())
	}
}
