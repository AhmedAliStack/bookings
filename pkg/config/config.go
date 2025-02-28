package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	UseCach 		bool
	TemplateCach 	map[string]*template.Template
	InProduction    bool
	AppSession      *scs.SessionManager 
}
