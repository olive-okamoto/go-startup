package config

import "text/template"

// Holds the application config
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
