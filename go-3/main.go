package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

//PORT port to be used
const PORT = "8080"

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func connect() *sql.DB {
	db, err := sql.Open("sqlite3", "file:database.sqlite?cache=shared")
	checkError(err)
	return db
}

// Actor class
type Actor struct {
	Name  string `json:"name"`
	Quote string `json:"quote"`
}

func main() {
	r := mux.NewRouter()
	http.Handle("/", r)
	r.Handle("/v1/quote", quote()).Methods("GET", "OPTIONS")
	r.Handle("/v1/quote/{actor}", quoteByActor()).Methods("GET", "OPTIONS")
	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":" + PORT,
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func quote() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		db := connect()
		result, err := db.Query("SELECT actor, detail FROM scripts ORDER BY RANDOM() LIMIT 1")
		checkError(err)
		defer result.Close()

		var detail string
		var actor string

		for result.Next() {
			err = result.Scan(&actor, &detail)
			checkError(err)
			// fmt.Println(actor, detail)
		}

		// fmt.Println(actor)

		if err := json.NewEncoder(w).Encode(Actor{Name: actor, Quote: detail}); err != nil {
			panic(err)
		}
	})
}

func quoteByActor() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		vars := mux.Vars(r)
		actor := strings.Replace(vars["actor"], "+", " ", -1)

		db := connect()

		query := "SELECT detail FROM scripts where actor LIKE ? ORDER BY RANDOM() LIMIT 1;"
		result, err := db.Query(query, actor)
		checkError(err)
		defer result.Close()

		var quote string

		for result.Next() {
			err = result.Scan(&quote)
			checkError(err)
		}

		if err := json.NewEncoder(w).Encode(Actor{Name: actor, Quote: quote}); err != nil {
			panic(err)
		}
	})
}
