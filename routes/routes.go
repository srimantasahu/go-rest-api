package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/srimantasahu/go-rest-api/controller"
	"github.com/srimantasahu/go-rest-api/models"
)

func MovieRoutes() *mux.Router {
	var router = mux.NewRouter()
	router = mux.NewRouter().StrictSlash(true)

	//Home Route
	router.HandleFunc("/api/", func(rw http.ResponseWriter, r *http.Request) {
		message := models.Message{
			Message: "Movie API!!!!",
		}
		json.NewEncoder(rw).Encode(message)
	})

	//Other Routes
	router.HandleFunc("/api/movies", controller.AddMovie).Methods(http.MethodPost)
	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods(http.MethodGet)
	router.HandleFunc("/api/movies/{id}", controller.GetMovieById).Methods(http.MethodGet)
	router.HandleFunc("/api/movies/{id}", controller.DeleteMovieById).Methods(http.MethodDelete)
	router.HandleFunc("/api/movies", controller.UpdateMovie).Methods(http.MethodPut)

	return router
}
