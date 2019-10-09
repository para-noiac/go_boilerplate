package app

import (
	"go_boilerplate/app/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type App struct {
	Router *mux.Router
}

func (app *App) Routing() {
	// app.DB = model.DBMigrate()
	app.Router = mux.NewRouter()
	// app.setRouters()
	app.Router.HandleFunc("/product", handler.GetProducts).Methods("GET")
	app.Router.HandleFunc("/test", handler.Test).Methods("GET")
}

// r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	title := vars["title"]
// 	page := vars["page"]

// 	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
// })

// func (app *App) CreateProduct(w http.ResponseWriter, r *http.Request) {
// 	handler.CreateProduct(app.DB, w, r)
// }

// func (app *App) GetProducts(w http.ResponseWriter, r *http.Request) {
// 	handler.GetProducts(app.DB, w, r)
// }

func (app *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, app.Router))
}
