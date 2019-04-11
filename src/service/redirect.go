package service

import (
	"config"
	"net/http"
	"strings"
	"time"
)


func redirectService() *http.Server {
	h := &httpHandler{}
	return &http.Server{
		Addr:         config.GetEnv().REDIRECT_SERVER_PORT,
		Handler:      h,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

type httpHandler struct {}

func (h httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	host := strings.Split(r.Host, ":")

	target := "https://" + host[0] + r.URL.Path

	http.Redirect(w, r, target, http.StatusMovedPermanently)
}
