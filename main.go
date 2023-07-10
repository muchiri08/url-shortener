package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	base62 "github.com/muchiri08/urlshortener/utils"
	"net/http"
)

// DBClient stores the database session information. Needs to be initialized once
type DBClient struct {
	db *sql.DB
}

type Record struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

func (driver *DBClient) GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	var url string
	vars := mux.Vars(r)

	//Get ID from base62 url
	id := base62.ToBase10(vars["encoded_string"])
	err := driver.db.QueryRow("SELECT url FROM web_url WHERE id = $1", id).Scan(&url)
	//Handles response details
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		responseMap := map[string]interface{}{"url": url}
		response, _ := json.Marshal(responseMap)
		w.Write(response)
	}
}

func main() {

}
