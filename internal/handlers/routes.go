package handlers

func (s *Server) routesInit() {
	home := s.Router.Group("/")
	{
		home.GET("", s.home())
	}

	courses := s.Router.Group("/courses")
	{
		courses.GET("", s.courses())

		courses.GET("/add", s.addCourse())

		courses.GET("/:id/edit", s.editCourse())

		courses.GET("/:id", s.previewCourse())
	}

	order := s.Router.Group("/orders")
	{
		order.GET("", s.order())
	}

	cart := s.Router.Group("/cart")
	{
		cart.GET("", s.cart())
	}
}
