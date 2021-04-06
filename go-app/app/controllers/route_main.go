package controllers

import (
	"log"
	"net/http"
	"text/template"
)

func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("go-app/app/views/templates/top.html")
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(w, nil)
}
