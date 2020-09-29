package main

import (
	"DangoBlog/data"
	"net/http"
)

// return /err?msg = content
// as a error msg
func err(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(w, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err != nil {
		error_message(w, r, "Cannot get threads")
	} else {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(w, threads, "layout", "private.navbar", "index")
		}
	}
}
