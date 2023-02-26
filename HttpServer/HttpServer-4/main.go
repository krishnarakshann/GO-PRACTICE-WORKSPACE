package main

import (
	"fmt"
	"net/http"
)

func demo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`Welcome to DEMO function  Http Server Imeplementation using handlers.handlerFunc`))
}

func stunt(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`Welcome to STUNT FUNCTION Http Server Imeplementation using handlers.handlerFunc`))
}

func main() {
	fmt.Println("Implementing Http Server in another way without using the mux")

	handler := http.HandlerFunc(demo) // converts the function demo to handler function and returns us a type handler
	http.Handle("/hello", handler)    //maps the given pattern to the handler

	handler1 := http.HandlerFunc(stunt)
	http.Handle("/hi", handler1)

	http.ListenAndServe("localhost:8083", nil)

}
