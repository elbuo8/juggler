package main

import (
	"github.com/elbuo8/juggler/app"
	"github.com/elbuo8/juggler/app/controllers"
	"github.com/elbuo8/juggler/app/middlewares"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"os"
)

func BuildServer() http.Handler {
	app, err := app.NewApp()
	if err != nil {
		app.Logger.Fatal(err)
	}
	r := mux.NewRouter()
	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(middlewares.SetDBSession(app))
	r.HandleFunc("/api/services", controllers.GetServices(app)).Methods("GET")
	r.HandleFunc("/api/services/{name}", controllers.GetService(app)).Methods("GET")
	r.HandleFunc("/api/services", controllers.CreateService(app)).Methods("POST")
	r.HandleFunc("/api/services/{name}", controllers.DeleteService(app)).Methods("DELETE")
	n.UseHandler(r)
	return n
}

func main() {
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), BuildServer()))
}
