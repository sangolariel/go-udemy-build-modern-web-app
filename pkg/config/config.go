package config

import (
	"html/template"
	"log"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCatche      bool
	TemplateCatche map[string]*template.Template
	Info           *log.Logger
}
