package router

import (
	"github.com/gorilla/mux"
	"github.com/is-ashish/mongo-api/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	// http://localhost:4000/api/movies
	router.HandleFunc("/api/movies", controller.GetAllMyMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateOneMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/deletAllMovies", controller.DeleteAllMovies).Methods("DELETE")
	return router
}
