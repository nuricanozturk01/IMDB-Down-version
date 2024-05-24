package app

import (
	"imdb_project/config"
	helper "imdb_project/data/dal"
	"log"
	"net/http"
)

func Run() {
	config.Load()
	if _, err := helper.InitDb(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	router := routes.New()
	log.Fatal(http.ListenAndServe(":8080", router))
}
