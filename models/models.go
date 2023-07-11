package models

import (
	"database/sql"
	"log"
)

func InitDB() (*sql.DB, error) {
	log.Println("InitDB called")
	var err error
	connStr := "user=root dbname=urls password=KiNuThiaPro$2 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Error occurred when connecting to db")
		return nil, err
	} else {
		//create model for our URL service
		stmt, err := db.Prepare("CREATE TABLE web_url(id SERIAL PRIMARY KEY, url TEXT NOT NULL);")
		if err != nil {
			log.Println(err)
			return nil, err
		}
		res, err := stmt.Exec()
		log.Println("Res: ", res)
	}

	return db, nil
}
