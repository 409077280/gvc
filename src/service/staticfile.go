package service

import (
	"config"
	"net/http"
	"time"
)


func staticService() *http.Server {
	http.Handle("/", http.FileServer(http.Dir(config.GetEnv().TEMPLATE_PATH)))
	return &http.Server{
		Addr:         config.GetEnv().STATIC_SERVER_PORT,
		Handler:      nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}