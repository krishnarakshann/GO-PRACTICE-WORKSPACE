package main

import (
	"fmt"
	"net/http"
)

func demo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`Hello Welcome to basics handlers server`))
}
func main() {

	fmt.Println("Implementing Basic Http Server without mux ")

	http.HandleFunc("/hello", demo)
	http.ListenAndServe("localhost:8082", nil)

}
