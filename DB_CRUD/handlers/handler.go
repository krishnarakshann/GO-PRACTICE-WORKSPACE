package handlers

import (
	"context"
	"dbcrud/db"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Demo struct {
	Name string  `json:"name"`
	Id   int     `json:"id"`
	Cgpa float64 `json:"cgpa"`
}

var dslice []Demo

const tablename = "demo"

func GetRecords(res http.ResponseWriter, r *http.Request) {
	db, err := db.GetDB()

	if err != nil {
		log.Fatal(err.Error())
		res.WriteHeader(500)
		res.Write([]byte(`Unable to connect to db`))
		return
	}

	ctx := context.Background()
	query := "SELECT name,id,cgpa FROM " + tablename
	rows, err := db.QueryContext(ctx, query)

	for rows.Next() {
		var d Demo
		if err := rows.Scan(&d.Name, &d.Id, &d.Cgpa); err != nil {
			log.Fatal(err.Error())
			return
		}
		dslice = append(dslice, d)
	}

	res.Header().Set("content-Type", "application/json")
	b, err := json.Marshal(dslice)

	if err != nil {
		log.Fatal(err.Error())
		res.WriteHeader(400)
		res.Write([]byte(`Error while marshalling the data`))
	}
	res.Write([]byte(b))
}

type SuccessResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   Demo   `json:"Demo"`
}

func PostRecord(w http.ResponseWriter, r *http.Request) {

	req, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err.Error())
		w.WriteHeader(400)
		w.Write([]byte(`Unable to read data`))
	}

	var d Demo
	err = json.Unmarshal(req, &d)

	if err != nil {
		log.Fatal(err.Error())
		w.WriteHeader(400)
		w.Write([]byte(`Unable to unmarshal the data`))
	}

	db, err := db.GetDB()

	if err != nil {
		log.Fatal(err.Error())
		w.WriteHeader(500)
		w.Write([]byte(`Unable to connect to db`))
	}

	query := "INSERT INTO " + tablename + " VALUES(?,?,?)"
	var args []interface{}
	args = append(args, d.Name, d.Id, d.Cgpa)

	ctx := context.Background()
	res, err := db.ExecContext(ctx, query, args...)

	if err != nil {
		w.WriteHeader(500)
		e := err.Error()
		w.Write([]byte(e))
	}

	rows_affected, _ := res.RowsAffected()
	if rows_affected <= 0 {
		w.Write([]byte(`unable to create record `))
	}

	s := SuccessResponse{
		Code:   201,
		Status: "Success",
		Data:   d,
	}
	w.Header().Set("content-Type", "application/json")
	b, err := json.Marshal(s)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`Error while marshalling the data`))
	}
	w.Write([]byte(b))
}

func GetRecordById(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	slice := strings.Split(path, "/")
	id_str := slice[2]

	/*
		Another way
		m := mux.Vars(r)
		id_str := m["Id"]
	*/

	id, err := strconv.Atoi(id_str)

	if err != nil {
		log.Fatal(err.Error())
		w.WriteHeader(400)
		w.Write([]byte(`Provide a Valid Id `))
		return
	}

	db, err := db.GetDB()
	if err != nil {
		log.Fatal(err.Error())
		w.WriteHeader(500)
	}

	query := "SELECT NAME,ID,CGPA FROM " + tablename + " WHERE ID = ?"
	var args []interface{}
	args = append(args, id)
	ctx := context.Background()

	row := db.QueryRowContext(ctx, query, args...)

	var d Demo
	if err := row.Scan(&d.Name, &d.Id, &d.Cgpa); err != nil {
		w.WriteHeader(500)
		log.Fatal(err.Error())
		w.Write([]byte(`Row Scan Error`))
		return
	}

	w.Header().Set("content-Type", "application/json")
	s := SuccessResponse{
		Code:   201,
		Status: "Success",
		Data:   d,
	}
	b, err := json.Marshal(s)
	w.Write([]byte(b))
}

func GetQueryParamMultiple(w http.ResponseWriter, r *http.Request) {

	//One Way
	m := r.URL.Query()
	param, present := m["name"] // returns a slice of string values

	if !present {
		fmt.Println("Key Doesnt Exist")
		w.WriteHeader(400)
		w.Write([]byte(`Invalid QueryParam`))
		return
	}

	//Another way is using r.URL.Query.Get("name") works only for single query param
	var args []interface{}
	query := "SELECT NAME,ID,CGPA FROM " + tablename + " WHERE NAME IN(?,?)"
	args = append(args, param[0], param[1])

	db, err := db.GetDB()

	if err != nil {
		w.WriteHeader(500)
		s := err.Error()
		w.Write([]byte(s))
		return
	}

	ctx := context.Background()
	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		log.Fatal(err.Error())
		w.WriteHeader(500)
		w.Write([]byte("Row Scan Error"))
		return
	}

	for rows.Next() {
		var d Demo
		if err = rows.Scan(&d.Name, &d.Id, &d.Cgpa); err != nil {
			w.WriteHeader(400)
			log.Fatal(err.Error())
			return
		}
		dslice = append(dslice, d)
	}

	w.Header().Set("content-Type", "application/json")

	b, err := json.Marshal(dslice)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`Marshalling Error`))
		return
	}
	w.Write([]byte(b))

}
