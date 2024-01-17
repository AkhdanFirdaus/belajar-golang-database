package belajargolangdatabase

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	dataSource := "root:Passw0rd@!@tcp(127.0.0.1:3306)/belajar_golang_database?parseTime=true"
	db, err := sql.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
