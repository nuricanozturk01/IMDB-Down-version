import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {catchError, map, Observable} from "rxjs";
import {REQUEST_SEARCH} from "./connection";
import {Company, Like, MovieDTO, Photo, SearchDTO, Trailer} from "../../dto/dtos";

@Injectable({
  providedIn: 'root'
})
export class SearchService {

  constructor(private http: HttpClient) {
  }

  mapToMovies(dto: SearchDTO, movies: any): SearchDTO {
    dto.movies = movies.map((movie: any) => {
      const dto = new MovieDTO()
      dto.id = movie.id
      dto.name = movie.name
      dto.year = movie.year
      dto.average_rate = movie.average_rate
      dto.popularity = movie.popularity
      dto.click_count = movie.click_count
      dto.trailers = movies.map((trailer: any) => {
        const trailerDTO = new Trailer()
        trailerDTO.id = trailer.id
        trailerDTO.url = trailer.url
        trailerDTO.media_type = trailer.media_type
        trailerDTO.media_id = trailer.media_id
        return trailerDTO
      })

      dto.likes = movies.map((like: any) => {
        const likeDTO = new Like()
        likeDTO.id = like.id
        likeDTO.media_id = like.media_id
        likeDTO.media_type = like.media_type
        likeDTO.user_id = like.user_id
        return likeDTO
      })

      dto.photos = movies.map((photo: any) => {
        const photoDTO = new Photo()
        photoDTO.id = photo.id
        photoDTO.url = photo.url
        photoDTO.media_type = photo.media_type
        photoDTO.media_id = photo.media_id
        return photoDTO
      })

      dto.companies = movies.map((company: any) => {
        const companyDTO = new Company()
        companyDTO.id = company.id
        companyDTO.media_id = company.media_id
        companyDTO.media_type = company.media_type
        companyDTO.name = company.name
        return companyDTO
      })


    })
    return dto;
  }

  search(keyword: string): Observable<SearchDTO> {
    return this.http.get(REQUEST_SEARCH(keyword), {withCredentials: true})
      .pipe(map((response: any) => {
          console.log(response)
          if (response.status_code === 200) {
            let dto = new SearchDTO()
            /*if (response.movies && response.movies.length > 0) {
              dto = this.mapToMovies(dto, response.movies)
            }
            if (response.tv_shows && response.tv_shows.length > 0) {
              return response.tv_shows;
            }
            if (response.celebs && response.celebs.length > 0) {
              return response.celebs;
            }*/
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
}
