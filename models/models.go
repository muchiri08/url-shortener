package models

import (
	"database/sql"
	"log"
)

func InitDB() (*sql.DB, error) {
	var err error
	connStr := "user=root dbname=urls password=KiNuThiaPro$2 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	} else {
		//create model for our URL service
		stmt, err := db.Prepare("CREATE TABLE web_url(id SERIAL PRIMARY KEY, url TEXT NOT NULL);")
		if err != nil {
			log.Println(err)
			return nil, err
		}
		res, err := stmt.Exec()
		log.Println(res)
	}

	return db, nil
}
