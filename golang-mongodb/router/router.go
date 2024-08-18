package router

import (
	"github.com/SrivathsaAcharya7/golang-mongodb/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")

	r.HandleFunc("/api/movies", controller.CreateMovies).Methods("POST")

	r.HandleFunc("/api/movies/{id}", controller.MarkAsWatched).Methods("PUT")

	r.HandleFunc("/api/movies/{id}", controller.DeleteAMovie).Methods("DELETE")

	r.HandleFunc("/api/movies", controller.DeleteAllMovies).Methods("DELETE")

	return r
}
