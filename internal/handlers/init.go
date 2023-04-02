package handlers

import (
	"html/template"
	"strings"

	"github.com/gin-gonic/gin"
)

type InitServer interface {
	InitServer() error
}

type files interface {
	initTemplates() (*template.Template, error)
	updateTemplates()
}

type routes interface {
	routesInit()
}

type controllers interface {
	home() gin.HandlerFunc
	courses() gin.HandlerFunc
	order() gin.HandlerFunc
	cart() gin.HandlerFunc
	addOrder() gin.HandlerFunc
}

type Server struct {
	Router         *gin.Engine
	ProductionMode bool
}

func (s *Server) InitServer() error {
	s.Router.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})

	go s.updateTemplates()

	s.Router.Static("/assets/js", "./dist/assets/js")
	s.Router.Static("/assets/css", "./dist/assets/css")

	s.Router.Use(gin.Logger())
	s.Router.Use(gin.Recovery())

	s.routesInit()

	if err := s.Router.Run(":5050"); err != nil {
		return err
	}

	return nil
}
