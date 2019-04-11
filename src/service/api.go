package service

import (
	"config"
	"net/http"
	"routers"
	"time"
)

func apiService() *http.Server {
	handle := routers.InitRouter()
	return &http.Server{
		Addr:         config.GetEnv().API_SERVER_PORT,
		Handler:      handle,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
