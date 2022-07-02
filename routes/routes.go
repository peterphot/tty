package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

var RegisterRoutes = func(router *mux.Router) {
	fs := http.FileServer(http.Dir("static"))
	router.Handle("/", fs)

	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("static/js/"))))
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("static/css/"))))

	//router.HandleFunc("/login", callback).Methods("GET")

}
