package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	base62 "github.com/muchiri08/urlshortener/utils"
	"io/ioutil"
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

// GetOriginalURL fetches the original URL for the given encoded(short)
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

// GenerateShortURL adds URL to DB and gives back shortened string
func (driver *DBClient) GenerateShortURL(w http.ResponseWriter, r *http.Request) {
	var id int
	var record Record
	postBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(postBody, &record)
	err := driver.db.QueryRow("INSERT INTO web_url(url) VALUES($1) RETURNING id", record.URL).Scan(&id)
	responseMap := map[string]interface{}{"encoded_string": base62.ToBase62(id)}
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(responseMap)
		w.Write(response)
	}
}

func main() {

}
