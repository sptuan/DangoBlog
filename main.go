package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

func main() {
	// Init Information
	logger.Infof("DangoBlog Start@%s\n", config.Address)

	// http server mux
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// starting http server
	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeOut * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeOut * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
