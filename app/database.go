package app

import (
	"database/sql"
	"go-restful-api/helper"
	"time"
)

func NewDb() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/go-restful-api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}
