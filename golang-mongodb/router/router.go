package router

import (
	"github.com/SrivathsaAcharya7/golang-mongodb/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	r.HandleFunc("/api/movie", controller.CreateMovies).Methods("POST")
	r.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	r.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")
	r.HandleFunc("/api/deleteallmovies", controller.DeleteAllMovies).Methods("DELETE")

	return r
}
