package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	restful "github.com/emicklei/go-restful"
)

type Book struct {
  isbn string
  title string
  author string
  price float32
}

var db *sql.DB

func init() {
	var err error
  // connStr := "host=localhost port=5432 user=taehoonkim dbname=bookstore sslmode=disable"
  // db, err := sql.Open("postgres", connStr)
	db, err = sql.Open("postgres", "postgres://taehoonkim:pass@localhost/bookstore?sslmode=disable")

  if err != nil {
    log.Fatal(err)
  }

  if err = db.Ping(); err != nil {
    log.Fatal(err)
  }
}

func main() {
	http.HandleFunc("/books", booksIndex)
	http.HandleFunc("/books/show", booksShow)
	http.ListenAndServe(":3000", nil)
}

func booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	bks := make([]*Book, 0)
	for rows.Next() {
		bk := new(Book)
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
	}
}


func booksShow(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	isbn := r.FormValue("isbn")
	if isbn == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	row := db.QueryRow("SELECT * FROM books WHERE isbn = $1", isbn)

	bk := new(Book)
	err := row.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
}

func New() http.Handler {
	api := new(restful.WebService)
	
}