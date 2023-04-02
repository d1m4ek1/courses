package handlers

import (
	"fmt"
	"html/template"
	"path/filepath"
	"time"
)

func (s *Server) initTemplates() (*template.Template, error) {
	var files []string

	var paths = []string{
		"./dist/*html",
		"./dist/templates/*.html"}

	for _, path := range paths {
		file, err := filepath.Glob(path)
		if err != nil {
			return nil, err
		}

		files = append(files, file...)
	}

	tmpls, err := template.ParseFiles(files...)
	if err != nil {
		return nil, err
	}

	return tmpls, nil
}

func (s *Server) updateTemplates() {
	const tickTimeMS = 2000
	ticker := time.NewTicker(tickTimeMS * time.Millisecond)

	templates, err := s.initTemplates()
	if err != nil {
		ticker.Stop()
		fmt.Printf("\nAn error occurred, template data was not updated!")
		return
	}

	s.Router.SetHTMLTemplate(templates)

	if !s.ProductionMode {
		fmt.Println("Template data has been updated")

		for range ticker.C {
			templates, err := s.initTemplates()
			if err != nil {
				ticker.Stop()
				fmt.Printf("\nAn error occurred, template data was not updated!")
				return
			}

			s.Router.SetHTMLTemplate(templates)
		}
	}
}
