package app

import (
	"imdb_project/config"
	helper "imdb_project/data/dal"
	"log"
)

func Run() {
	config.Load()
	db, err := helper.InitDb()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	//router := routes.New()
	//log.Fatal(http.ListenAndServe(":8080", router))

	myHelper := helper.New(db)

	movies := myHelper.Search("Red")

	if len(movies.Movies) != 0 {
		for _, movie := range movies.Movies {
			log.Println(movie.Name)
		}
	}
}
