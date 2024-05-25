package app

import (
	"imdb_project/config"
	helper "imdb_project/data/dal"
	"log"
)

func Run() {
	config.Load()
	_, err := helper.InitDb()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	//router := routes.New()
	//log.Fatal(http.ListenAndServe(":8080", router))
	//imdbHelper := helper.New(db)

	/*movie := entity.Movie{
		Name: "Inception",
		Photos: []entity.Photo{
			{URL: "photo1.jpg"},
			{URL: "photo2.jpg"},
		},
		Trailers: []entity.Trailer{
			{URL: "trailer1.mp4"},
			{URL: "trailer2.mp4"},
		},
	}

	tvShow := entity.TVShow{
		Name: "Breaking Bad",
		Photos: []entity.Photo{
			{URL: "photo1.jpg"},
			{URL: "photo2.jpg"},
		},
		Trailers: []entity.Trailer{
			{URL: "trailer1.mp4"},
			{URL: "trailer2.mp4"},
		},
	}

	user := entity.User{
		Username: "John Doe",
		Photos: []entity.Photo{
			{URL: "user_photo1.jpg"},
			{URL: "user_photo2.jpg"},
		},
	}

	celebrity := entity.Celebrity{
		Name: "Jane Smith",
		Photos: []entity.Photo{
			{URL: "celeb_photo1.jpg"},
			{URL: "celeb_photo2.jpg"},
		},
	}

	db.Create(&movie)
	db.Create(&tvShow)
	db.Create(&user)
	db.Create(&celebrity)

	watchListItem1 := entity.WatchListItem{
		WatchListID: user.WatchList.ID,
		MediaID:     movie.ID,
		MediaType:   "movies",
	}
	watchListItem2 := entity.WatchListItem{
		WatchListID: user.WatchList.ID,
		MediaID:     tvShow.ID,
		MediaType:   "tv_shows",
	}

	db.Create(&watchListItem1)
	db.Create(&watchListItem2)

	like1 := entity.Like{
		UserID:    user.ID,
		MediaID:   movie.ID,
		MediaType: "movies",
	}
	db.Create(&like1)

	like2 := entity.Like{
		UserID:    user.ID,
		MediaID:   tvShow.ID,
		MediaType: "tv_shows",
	}
	db.Create(&like2)*/
}
