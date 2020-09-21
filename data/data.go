package data

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	// note: postgres has registed in drivers
	Db, err = sql.Open("postgres",
		"user=dango_admin dbname=dango_blog_dev"+
			" password=dange_admin_passwd ssl_mode-disable")
	if err != nil {
		// TODO: change logger same to package main
		log.Fatal(err)
	}
	return
}

// TODO: impl UUID here
