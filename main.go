package main

import (
	"net/http"

	"github.com/zaileon/chitchat/data"
)

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err == nil {
		_, e := session(w, r)
		if e != nil {
			generateHTML(w, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(w, threads, "layout", "private.navbar", "index")
		}
	}
}

func main() {

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", index)
	// mux.HandleFunc("/err", err)

	// mux.HandleFunc("/login", login)
	// mux.HandleFunc("/logout", logout)
	// mux.HandleFunc("/signup", signup)
	// mux.HandleFunc("/signup_account", signupAccount)
	// mux.HandleFunc("/authenticate", authenticate)

	// mux.HandleFunc("/thread/new", newThread)
	// mux.HandleFunc("/thread/create", createThread)
	// mux.HandleFunc("/thread/post", postThread)
	// mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
