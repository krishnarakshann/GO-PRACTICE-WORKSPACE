package db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

func GetDB() (*sql.DB, error) {

	config := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Addr:                 "localhost:3306",
		DBName:               "practice",
		AllowNativePasswords: true,
		Net:                  "tcp",
	}

	conf := config.FormatDSN()
	db, err := sql.Open("mysql", conf)

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return db, nil
}
