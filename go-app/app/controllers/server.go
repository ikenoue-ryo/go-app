package controllers

import (
	"go-project/go-app/config"
	"net/http"
)

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
