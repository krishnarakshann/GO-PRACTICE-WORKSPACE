package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
	"threelayer/internal/handlers"
	services "threelayer/internal/services/student"
	stores "threelayer/internal/stores/student"
)

func main() {
	fmt.Println("Welcome to three Layer Architecture ")

	db, err := getDB()

	if err != nil {
		fmt.Println("Failed to connect to db due to ", err)
		return
	} else {
		fmt.Println("Succesfully Connected to DB ")
	}

	store := stores.New(db)
	svc := services.New(store)
	handler := handlers.New(svc)

	r := mux.NewRouter()
	r.HandleFunc("/getall", handler.GetStudents).Methods("GET")
	r.HandleFunc("/post", handler.PostStudent).Methods("POST")
	r.HandleFunc("/get/{id}", handler.GetStudentById).Methods("GET")

	http.ListenAndServe("localhost:8080", r)

}

func getDB() (*sql.DB, error) {
	config := mysql.Config{
		User:                 "root",
		Passwd:               "password",
		Addr:                 "localhost:3306",
		DBName:               "practice",
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	conf := config.FormatDSN()
	db, err := sql.Open("mysql", conf)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		fmt.Println("Falied to ping to db")
		return nil, err
	}

	return db, nil
}
