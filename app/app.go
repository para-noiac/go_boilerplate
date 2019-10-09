package app

import (
	"boilerplate/app/handler"
	"boilerplate/app/model"
	"boilerplate/config"
	"fmt"
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

func (app *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DB.Host,
		config.DB.Port,
		config.DB.User,
		config.DB.Password,
		config.DB.Dbname,
	)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		panic(err)
	}

	app.DB = model.DBMigrate(db)
	app.Router = mux.NewRouter()
	app.setRouters()
}

func (app *App) setRouters() {
	app.Post("/product", app.CreateProduct)
	app.Router.HandleFunc("/product", app.GetProducts).Methods("GET")
}

func (app *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("POST")
}

func (app *App) CreateProduct(w http.ResponseWriter, r *http.Request) {
	handler.CreateProduct(app.DB, w, r)
}

func (app *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, app.Router))
}

func (app *App) GetProducts(w http.ResponseWriter, r *http.Request) {
	handler.GetProducts(app.DB, w, r)
}
