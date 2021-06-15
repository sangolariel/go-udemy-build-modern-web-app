package config

import "html/template"

// AppConfig holds the application config
type AppConfig struct {
	TemplateCatche map[string]*template.Template
}
