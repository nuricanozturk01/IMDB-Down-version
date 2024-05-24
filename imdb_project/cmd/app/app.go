package app

import (
	"imdb_project/config"
	"imdb_project/database/database"
)

func Run() {
	config.Load()
	database.Init()
}
