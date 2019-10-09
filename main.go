package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// type Product struct {
// 	gorm.Model
// 	Code  string
// 	Price uint
// }

// func main() {
// 	db, err := gorm.Open("postgres", "host=localhost port=5432 user=leekahwei dbname=test password=1 sslmode=disable")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()
// }

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	http.ListenAndServe(":80", r)
}
