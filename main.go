package main

import (
	"log"
	"net/http"
	"os"

	"github.com/srimantasahu/go-rest-api/routes"
)

func main() {

	router := routes.MovieRoutes()

	http.Handle("/api", router)

	log.Println("Listening on PORT : " + os.Getenv("PORT"))

	log.Println(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
