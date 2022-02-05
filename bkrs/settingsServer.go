package bkrs

import (
	"html/template"
	"log"
	"net/http"
	"reflect"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ts, err := template.ParseFiles("./resources/index.page.tmpl")
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	type Feature struct {
		Name  string
		Value bool
	}
	refl := reflect.ValueOf(globalConfig.Features)
	reflType := refl.Type()
	features := []Feature{}
	for i := 0; i < refl.NumField(); i++ {
		features = append(features, Feature{reflType.Field(i).Name, refl.Field(i).Interface().(bool)})
	}

	type Shortcut struct {
		Name  string
		Value string
	}
	refl = reflect.ValueOf(globalConfig.Shortcuts)
	reflType = refl.Type()
	shortcuts := []Shortcut{}
	for i := 0; i < refl.NumField(); i++ {
		shortcuts = append(shortcuts, Shortcut{reflType.Field(i).Name, refl.Field(i).Interface().(string)})
	}

	data := map[string]interface{}{
		"Features":  features,
		"Shortcuts": shortcuts,
		"Version":   version,
	}

	err = ts.Execute(w, data)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal Server Error", 500)
	}
}

func RunSettingsServer() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./resources"))))
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":7838", nil))
}
