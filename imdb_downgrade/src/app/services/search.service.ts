import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {catchError, map, Observable} from "rxjs";
import {
  ADD_WATCH_LIST_MOVIE,
  ADD_WATCH_LIST_TV,
  LIKE_MOVIE,
  REMOVE_WATCH_LIST_MOVIE,
  REMOVE_WATCH_LIST_TV,
  REQUEST_ALL_COUNTRIES,
  REQUEST_ALL_MOVIES,
  REQUEST_ALL_TV_SHOW,
  REQUEST_CELEBRITY_DETAILS,
  REQUEST_CITIES_BY_COUNTRY,
  REQUEST_MOVIE_DETAILS,
  REQUEST_RATE_MOVIE,
  REQUEST_RATE_TV_SHOW,
  REQUEST_SEARCH,
  REQUEST_TV_SHOW_DETAILS,
  REQUEST_USER_INFO,
  REQUEST_WATCH_LIST
} from "./connection";
import {
  Celebrity,
  CelebrityDTO,
  CityDTO,
  Company,
  CountryDTO,
  Like,
  MovieDTO,
  Photo,
  SearchDTO,
  Trailer,
  TvShowDTO,
  WatchListDTO
} from "../../dto/dtos";

@Injectable({
  providedIn: 'root'
})
export class SearchService {

  constructor(private http: HttpClient) {
  }


  search(keyword: string): Observable<SearchDTO> {
    return this.http.get(REQUEST_SEARCH(keyword), {withCredentials: true})
      .pipe(map((response: any) => {
          console.log(response)
          if (response.status_code === 200) {
            let dto = new SearchDTO()
            dto.keyword = response.data.keyword
            dto.movies = response.data.movies
            dto.tv_shows = response.data.tv_shows
            dto.celebrities = response.data.celebs
            return dto;
          }
        }),
        catchError((error: any) => {
            throw error;
          }
        ));
  }

  findAllMovies(): Observable<MovieDTO[]> {
    return this.http.get(REQUEST_ALL_MOVIES, {withCredentials: true})
      .pipe(map((response: any) => {
          if (response.status_code === 200) {
            return response.data.map((movie: any) => {
              const dto = new MovieDTO()
              dto.id = movie.id
              dto.name = movie.name
              dto.year = movie.year
              dto.average_rate = movie.average_rate
              dto.popularity = movie.popularity
              dto.click_count = movie.click_count
              dto.description = movie.description


              dto.trailers = movie.trailers.map((trailer: any) => {
                const trailerDTO = new Trailer()
                trailerDTO.id = trailer.id
                trailerDTO.url = trailer.url
                trailerDTO.media_type = trailer.media_type
                trailerDTO.media_id = trailer.media_id
                return trailerDTO
              })

              dto.likes = movie.likes.map((like: any) => {
                const likeDTO = new Like()
                likeDTO.id = like.id
                likeDTO.media_id = like.media_id
                likeDTO.media_type = like.media_type
                likeDTO.user_id = like.user_id
                return likeDTO
              })

              dto.photos = movie.photos.map((photo: any) => {
                const photoDTO = new Photo()
                photoDTO.id = photo.id
                photoDTO.url = photo.url
                photoDTO.media_type = photo.media_type
                photoDTO.media_id = photo.media_id
                return photoDTO
              })

              dto.companies = movie.companies.map((company: any) => {
                const companyDTO = new Company()
                companyDTO.id = company.id
                companyDTO.media_id = company.media_id
                companyDTO.media_type = company.media_type
                companyDTO.name = company.name
                return companyDTO
              })

              dto.celebs = movie.celebs.map((celebrity: any) => {
                const celebrityDTO = new Celebrity()
                celebrityDTO.id = celebrity.id
                celebrityDTO.name = celebrity.name
                celebrityDTO.movies = celebrity.movies
                celebrityDTO.tv_shows = celebrity.tv_shows
                celebrityDTO.photos = celebrity.photos
                return celebrityDTO
              })

              return dto;

            })
          }
        }),
        catchError((error: any) => {
            throw error;
          }
        ));
  }

  addOnWatchList(id: string, type: string = "movie"): Observable<string> {
    let url = ADD_WATCH_LIST_MOVIE(id)
    if (type === "tv_show") {
      url = ADD_WATCH_LIST_TV(id)
    }
    return this.http.post(url, {}, {withCredentials: true}).pipe(
      map((response: any) => {
        if (response.status_code === 200) {
          return response.message;
        }

        return response.message;
      }),
      catchError((error: any) => {
          return ["Already added to watch list!"]
        }
      ));
  }


  removeOnWatchList(id: string, type: string = "movie"): Observable<string> {
    let url = REMOVE_WATCH_LIST_MOVIE(id)
    if (type === "tv_show") {
      url = REMOVE_WATCH_LIST_TV(id)
    }
    return this.http.delete(url, {withCredentials: true}).pipe(
      map((response: any) => {
        console.log("RESPONSE: ", response)
        if (response.status_code === 200) {
          return response.message;
        }
      }),
      catchError((error: any) => {
          return ["Already removed from watch list!"]
        }
      ));
  }

  likeMovie(movie_id: string) {
    return this.http.post(LIKE_MOVIE(movie_id), {}, {withCredentials: true}).pipe(
      map((response: any) => {
        if (response.status_code === 200) {
          return response.message;
        }
      }),
      catchError((error: any) => {
          return ["Already liked!"]
        }
      ));
  }


  findMovieDetails(id: string): Observable<MovieDTO> {
    return this.http.get(REQUEST_MOVIE_DETAILS(id), {withCredentials: true}).pipe(
      map((response: any) => {
        if (response.status_code === 200) {
          return response.data;
        }
      }),
      catchError((error: any) => {
          return null
        }
      ));
  }

  getWatchList(): Observable<WatchListDTO> {
    return this.http.get(REQUEST_WATCH_LIST, {withCredentials: true}).pipe(
      map((response: any) => {
        if (response.status_code === 200) {
          console.log("RESPONSE: ", response)
          return response.data;
        }
      }),
      catchError((error: any) => {
          return null
        }
      ));
  }

  findTvShowDetails(id: string) {
    return this.http.get(REQUEST_TV_SHOW_DETAILS(id), {withCredentials: true}).pipe(
      map((response: any) => {
        if (response.status_code === 200) {
          return response.data;
        }
      }),
      catchError((error: any) => {
          return null
        }
      ));
  }

  findAllTvShows(): Observable<TvShowDTO[]> {
    return this.http.get(REQUEST_ALL_TV_SHOW, {withCredentials: true})
      .pipe(map((response: any) => {

          if (response.status_code === 200) {
            return response.data.map((tv: any) => {
              const dto = new TvShowDTO()
              dto.id = tv.id
              dto.name = tv.name
              dto.year = tv.year
              dto.average_rate = tv.average_rate
              dto.popularity = tv.popularity
              dto.click_count = tv.click_count
              dto.episode_count = tv.episode_count
              dto.season_count = tv.season_count
              dto.description = tv.description

              dto.trailers = tv.trailers.map((trailer: any) => {
                const trailerDTO = new Trailer()
                trailerDTO.id = trailer.id
                trailerDTO.url = trailer.url
                trailerDTO.media_type = trailer.media_type
                trailerDTO.media_id = trailer.media_id
                return trailerDTO
              })

              if (tv.likes !== null) {
                dto.likes = tv.likes.map((like: any) => {
                  const likeDTO = new Like()
                  likeDTO.id = like.id
                  likeDTO.media_id = like.media_id
                  likeDTO.media_type = like.media_type
                  likeDTO.user_id = like.user_id
                  return likeDTO
                })
              }


              dto.photos = tv.photos.map((photo: any) => {
                const photoDTO = new Photo()
                photoDTO.id = photo.id
                photoDTO.url = photo.url
                photoDTO.media_type = photo.media_type
                photoDTO.media_id = photo.media_id
                return photoDTO
              })

              dto.companies = tv.companies.map((company: any) => {
                const companyDTO = new Company()
                companyDTO.id = company.id
                companyDTO.media_id = company.media_id
                companyDTO.media_type = company.media_type
                companyDTO.name = company.name
                return companyDTO
              })

              dto.celebs = tv.celebs.map((celebrity: any) => {
                const celebrityDTO = new Celebrity()
                celebrityDTO.id = celebrity.id
                celebrityDTO.name = celebrity.name
                celebrityDTO.movies = celebrity.movies
                celebrityDTO.tv_shows = celebrity.tv_shows
                celebrityDTO.photos = celebrity.photos
                return celebrityDTO
              })

              return dto;
            })
          }
        }),
        catchError((error: any) => {
            throw error;
          }
        ));
  }

  rateMovie(id: string, rate: number) {
    return this.http.post(REQUEST_RATE_MOVIE(id, rate), {}, {withCredentials: true}).pipe(
      map((response: any) => {
        if (response.status_code === 200) {
          return response.message;
        }
      }),
      catchError((error: any) => {
          return ["Already rated!"]
        }
      ));
  }

  rateTvShow(id: string, rate: number) {
    return this.http.post(REQUEST_RATE_TV_SHOW(id, rate), {}, {withCredentials: true}).pipe(
      map((response: any) => {
        if (response.status_code === 200) {
          return response.message;
        }
      }),
      catchError((error: any) => {
          return ["Already rated!"]
        }
      ));
  }


  findCelebrityById(id: string): Observable<CelebrityDTO> {
    return this.http.get(REQUEST_CELEBRITY_DETAILS(id), {withCredentials: true}).pipe(
      map((response: any) => {
        if (response.status_code === 200) {
          console.log("RESPONSE: ", response)
          return response.data;
        }
      }),
      catchError((error: any) => {
          return null
        }
      ));
  }

  getCountries(): Observable<CountryDTO[]> {
    return this.http.get(REQUEST_ALL_COUNTRIES).pipe(
      map((response: any) => {
        return response.countries;
      }),
      catchError((error: any) => {
          return null
        }
      ));
  }

  getCitiesByCountryName(countryName: string): Observable<CityDTO[]> {
    return this.http.get(REQUEST_CITIES_BY_COUNTRY(countryName)).pipe(
      map((response: any) => {
        console.log("RESPONSE: ", response)
        return response.cities;
      }),
      catchError((error: any) => {
          return null
        }
      ));
  }

  getUserInfo() : Observable<any> {
    return this.http.get(REQUEST_USER_INFO, {withCredentials: true}).pipe(
      map((response: any) => {
        console.log("RESPONSE: ", response)
        return response.data;
      }),
      catchError((error: any) => {
          return null
        }
      ));
  }
}
