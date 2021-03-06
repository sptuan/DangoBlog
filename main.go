package main

import (
	//"github.com/julienschmidt/httprouter"
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

	// set route here
	mux.HandleFunc("/", index)
	// error
	mux.HandleFunc("/err", err)

	// login functions
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	// post and threads
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	// starting http server
	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeOut * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeOut * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
