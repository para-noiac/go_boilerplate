package main

import (
	"boilerplate/app"
)

func main() {
	app := &app.App{}
	app.Initialize()
	app.Run(":3333")
}

// 	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
// 		vars := mux.Vars(r)
// 		title := vars["title"]
// 		page := vars["page"]

// 		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
// 	})
