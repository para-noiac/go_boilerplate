package app

import (
	"boilerplate/app/handler"
	"boilerplate/app/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (app *App) Initialize() {
	app.DB = model.DBMigrate()
	app.Router = mux.NewRouter()
	// app.setRouters()
	app.Router.HandleFunc("/product", handler.CreateProduct).Methods("GET")

}

// 	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
// 		vars := mux.Vars(r)
// 		title := vars["title"]
// 		page := vars["page"]

// 		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
// 	})

// func (app *App) CreateProduct(w http.ResponseWriter, r *http.Request) {
// 	handler.CreateProduct(app.DB, w, r)
// }

// func (app *App) GetProducts(w http.ResponseWriter, r *http.Request) {
// 	handler.GetProducts(app.DB, w, r)
// }

func (app *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, app.Router))
}
