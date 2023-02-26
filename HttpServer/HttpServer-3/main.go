package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`Welcome to Hello function bro`))
}

// creating a struct to implement servehttp method
type demoHandler struct{}

func (d demoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`Hi,welcome to handlers server session`))
}

func main() {
	fmt.Println("Implementing Http Server using the default Builtin Mux ")

	mux := http.NewServeMux()       //create a mux
	mux.HandleFunc("/hello", hello) //mapping the handler function to the pattern

	d := demoHandler{}
	mux.Handle("/hi", d) // maps the handler type  to the given pattern

	//handlers.ListenAndServe("localhost:8080", mux)

	//Another way of ListenAndServe
	start := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	start.ListenAndServe()
}
