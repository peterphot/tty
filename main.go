package main

import (
	"github.com/gorilla/mux"
	"github.com/peterphot/tty/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	log.Fatal(http.ListenAndServe(":9000", r))
}
