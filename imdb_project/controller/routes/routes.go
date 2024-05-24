package routes

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	//movieController := controller.New()
	/*router.HandleFunc("/movies", movieController.GetMovies).Methods("GET")
	router.HandleFunc("/movies", movieController.CreateMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", movieController.GetMovieByID).Methods("GET")
	router.HandleFunc("/movies/{id}", movieController.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", movieController.DeleteMovie).Methods("DELETE")*/
	return router
}
