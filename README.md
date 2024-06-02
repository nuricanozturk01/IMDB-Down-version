# IMDB-Down-version Project

**for SE-3355 Project**

## Description

Create a slimmed down version of IMDB web application that will have following use cases/screens

- Generic Search will search movie title, summaries and actors.
    - Add a drop-down search like above to be specific on the search.
    - On first three letters, show maximum of 3 top items that satisfy the search text
- Show a list of at least 10 movies in the opening page (with a slider) with IMDB scores, image and a trailer
- If not logged in, login will be visible at the header.
- From top right, user can change language to English/Turkish. Only Home page should be
  in 2 languages EN/TR. Default should be browser language
- Login with email/password or Google will be supported.
- User can register with just email, password (at least 8 characters, 1 number and 1 none
  alphanumeric character) and country, city
- If google authenticated or after manual registration, user will be directed to HOME
  PAGE.
    - Header will have the username
    - Users will be able to add movies to watch list and rate/comment on movies (when
      not logged
    - When NOT logged in, users clicking on Add to Watchlist, Rate will be directed to
      login page
- **In Detail Page:**
- One image and one clip of the movie
- Rating of the movie
- Popularity of the movie. This is a business logic you will need to create based on research
  and some assumptions. Generally popularity is a mix of ratings, comments, page views etc. Also need to show ranking in
  popularity

## Technologies, Languages, and Frameworks

- Generic Repository Pattern
- Layered Architecture
- Monolithic Architecture
- GoLang (1.22)
    - Gin Gonic Framework
    - Gorm
    - Gorilla Sessions
    - BigCache (caching (country-city))
    - GoDotEnv
    - unit testing
    - goroutines
    - OAuth2
    - Crypto (bcrypt)
    - Scheduled Tasks (GoCron)
    - Validation (validator) (go-playground)
    - Cors
- MySQL
- SQLite3 (for testing)
- Amazon Simple Queue Service (SQS) (for popularity calculation)

## Assumptions

- Each user will have one watchlist.
- Users can only add movies to their own watchlists.
- Users can only delete movies from their own watchlists.
- Only registered users can rate movies.
- Only registered users can view movie details.
- Only registered users can search for movies.
- Each movie can have multiple trailers and images, but in the application, only one trailer and one image will be
  shown.
- Each movie and TV show can only have one rating.
- Popularity is calculated based on the number of the click count to media, rating and adding to watchlist. In
  popularity calculation, I use Amazon SQS because it is a good way to handle the popularity calculation in the
  background. Also,
  decreased the unit time write operations to the database.
- The popularity of the media is calculated every 30 minutes from Queue. (Every 30 minutes, the popularity of the media
  is calculated and written to the database.)
- Localized languages are English and Turkish. The default language is the browser language. If the browser language is
  not English or Turkish, the default language is English.

## Architecture Diagram

<img src="/images/arc.png" width="100%">

## Database Schema

<img src="/images/db.png" width="100%">
