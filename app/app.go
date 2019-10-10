package app

import (
	"go_boilerplate/app/handler"
	"go_boilerplate/app/model"
	"go_boilerplate/config"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type App struct {
	Router *mux.Router
}

func (app *App) setRouters() {
	handler := &handler.Handler{}
	handler.DB = model.DBMigrate(config.InitDB())
	app.Router.HandleFunc("/product", handler.CreateProduct).Methods("POST")
	app.Router.HandleFunc("/product", handler.GetProducts).Methods("GET")
	app.Router.HandleFunc("/test", handler.Test).Methods("GET")
}

func (app *App) run(host string) {
	log.Fatal(http.ListenAndServe(host, app.Router))
}

func (app *App) Start() {
	app.Router = mux.NewRouter()
	app.setRouters()
	app.run(":" + os.Getenv("PORT"))
}
