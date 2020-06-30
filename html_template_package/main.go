package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// Entry is structure
type Entry struct {
	Number, Sqaure, Double int
}

// DATA wil keep data
var DATA []Entry

// HTMLFILE will have the html file
const HTMLFILE = "file.html"

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Host: %s Url: %s\n", r.Host, r.URL)
	myT := template.Must(template.ParseGlob(HTMLFILE))
	myT.ExecuteTemplate(w, HTMLFILE, DATA)
}

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		log.Fatalf("Error while open file: %v", err)
	}

	_, err = db.Exec("DELETE FROM data")
	if err != nil {
		log.Fatalf("Error while deleting data from data table %v", err)
	}

	query, err := db.Prepare("INSERT INTO data(number, double, sqaure) VALUES(?,?,?)")
	if err != nil {
		log.Fatalf("Error while preparing query: %v", err)
	}
	for i := 10; i < 200; i++ {
		_, _ = query.Exec(i, i*2, i*i)
	}
	query.Close()

	rows, err := db.Query("SELECT * FROM data")
	if err != nil {
		log.Fatalf("Error while selecting data from data table %v", err)
	}

	var number, double, sqaure int
	for rows.Next() {
		_ = rows.Scan(&number, &double, &sqaure)
		DATA = append(DATA, Entry{number, sqaure, double})
	}

	http.HandleFunc("/", myHandler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error while listne to the port :%v", err)
	}
}
