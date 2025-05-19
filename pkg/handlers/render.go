package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/AhmedAliStack/bookings/pkg/config"
	"github.com/AhmedAliStack/bookings/pkg/models"
	"github.com/justinas/nosurf"
)

var appConfig *config.AppConfig
func SetConfig(value *config.AppConfig){
	appConfig = value
}

func AddToDefault(td *models.TemplateData, r *http.Request) *models.TemplateData{
	td.CSRFToken = nosurf.Token(r)
	return td
}

func renderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData, r *http.Request) {
	var tc map[string]*template.Template
	if appConfig.UseCach{
		tc = appConfig.TemplateCach
	} else {
		tc = CreateTemplateCach()
	}
	
	t, inMap := tc[tmpl]
	if !inMap{
		log.Fatal("Page is not found")
	}

	buf := new(bytes.Buffer)
	td = AddToDefault(td, r)
	err := t.Execute(buf, td)

	if err != nil {
		fmt.Println("error while parsing file", tmpl, ":", err)
		return
	}

	_,err = buf.WriteTo(w)

	if err != nil{
		fmt.Println("error while parsing file", tmpl, ":", err)
		return
	}
}

func CreateTemplateCach() (map[string]*template.Template){
	myCach := map[string]*template.Template{}

	pages,err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCach
	}

	for _,page := range pages {
		name := filepath.Base(page)
		ts,err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCach
		}

		matchs,err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCach
		}
		if len(matchs) > 0 {
			ts,err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCach
			}
		}

		myCach[name] = ts
	}

	return myCach
}
