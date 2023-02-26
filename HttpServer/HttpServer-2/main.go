package main

import (
	"fmt"
	"net/http"
)

type demohandler struct{}

func (d *demohandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`Welcome to Http Server Implementation without using mux where handler is a type `))
}

type practicehandler struct{}

func (p *practicehandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`Welcome to Another endpoint Http Server Implementation without using mux where handler is a type `))
}

func main() {
	fmt.Println("Implenting Http Server when the handler is a type ")

	d := demohandler{}
	http.Handle("/hello", &d)

	p := practicehandler{}
	http.Handle("/hi", &p)

	http.ListenAndServe("localhost:8080", nil)

}
