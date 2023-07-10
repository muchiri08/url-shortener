package main

import "database/sql"

// DBClient stores the database session information. Needs to be initialized once
type DBClient struct {
	db *sql.DB
}

type Record struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

func main() {

}
