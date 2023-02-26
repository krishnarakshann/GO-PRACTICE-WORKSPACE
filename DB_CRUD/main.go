package main

import (
	"dbcrud/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	fmt.Println("Welcome to db  CRUD PRACTICE ")

	//creating Routes

	r := mux.NewRouter()
	r.HandleFunc("/get", handlers.GetRecords).Methods("GET")
	r.HandleFunc("/post", handlers.PostRecord).Methods("POST")
	r.HandleFunc("/get/{Id}", handlers.GetRecordById).Methods("GET")
	r.HandleFunc("/getquery", handlers.GetQueryParamMultiple).Methods("GET")
	//Similarly for Update and Delete

	http.ListenAndServe("localhost:8080", r)

}
