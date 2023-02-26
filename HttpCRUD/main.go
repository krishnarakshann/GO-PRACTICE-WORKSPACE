package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strings"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var s []Student

func print(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json") //displaying the data in json format in response
	b, err := json.Marshal(s)                          //returns the slice of bytes
	if err != nil {
		fmt.Println("There is an error ")
	}
	w.Write([]byte(b))

	// Another way of marshaling
	err = json.NewEncoder(w).Encode(s) // returns an error if there is an error

	if err != nil {
		fmt.Println("There is an error while marshaling")
	}

	// Another way
	a, err := json.Marshal(s)

	if err != nil {
		fmt.Println("There is an error while marshalling")
	}
	w.Write([]byte(a))
}

func post(w http.ResponseWriter, r *http.Request) {

	var student Student
	rbody, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(400)
		fmt.Println("error while reading data")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.Unmarshal(rbody, &student)

	if err != nil {
		w.WriteHeader(400)
		fmt.Println("There is an error")
		return
	}
	s = append(s, student)
	fmt.Println("After unmarshalling the data ", s)
}

func queryMultiple(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name, present := query["name"]
	age, present := query["age"]

	if !present || len(name) == 0 {
		fmt.Println("Name not present")
	}
	fmt.Println("query params name are ", name)
	fmt.Println("query params age are ", age)

	/*

		Another way of queryparams

		r.ParseForm()
		name , present := r.Form["name"]

	*/

}
func querySingle(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Println("Name is ", name)
}

func pathParam(w http.ResponseWriter, r *http.Request) {

	n := mux.Vars(r)
	name := n["name"]

	fmt.Println("Name is ", name)

	//Another way of path param
	path := r.URL.Path
	fmt.Println(path)
	slice := strings.Split(path, "/")
	fmt.Println(slice)
	fmt.Println(slice[2], "is the path param")

}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/display", print).Methods("GET")
	r.HandleFunc("/post", post).Methods("POST")

	// for query params
	r.HandleFunc("/query", queryMultiple).Methods("GET")
	r.HandleFunc("/querysingle", querySingle).Methods("GET")
	r.HandleFunc("/pathparam/{name}", pathParam).Methods("GET")

	http.ListenAndServe("localhost:8080", r)

}
