package routes

import (
	"html/template"
	"net/http"
)

func WebIndex(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("main.html")
	err := t.Execute(w, nil)
	if err != nil {
		return
	}
}
