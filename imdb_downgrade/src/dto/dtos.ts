export class MovieDTO {
  public id: string;
  public name: string;
  public average_rate: string
  public year: number
  public popularity: number
  public click_count: number
  public trailers: Trailer[]
  public photos: Photo[]
  public companies: Company[]
  public likes: Like[]
  public celebrities: Celebrity[]
}

export class TvShowDTO {
  public id: string;
  public name: string;
  public year: number
  public average_rate: string
  public popularity: number
  public click_count: number
  public episode_count: number
  public season_count: number
  public trailers: Trailer[]
  public photos: Photo[]
  public companies: Company[]
  public likes: Like[]
  public celebrities: Celebrity[]

}


export class Trailer {
  public id: string;
  public media_id: string;
  public media_type: string;
  public url: string;
}

export class Photo {
  public id: string;
  public media_id: string;
  public media_type: string;
  public url: string;
}

export class Company {
  public id: string;
  public media_id: string;
  public media_type: string;
  public name: string;
}

export class Like {
  public id: string;
  public media_id: string;
  public media_type: string;
  public user_id: string;
}

export class Movie {
  public id: string;
  public name: string;
  public average_rate: string
  public year: number
  public popularity: number
  public click_count: number
  public trailers: Trailer[]
  public photos: Photo[]
  public companies: Company[]
  public likes: Like[]
  public celebs: Celebrity[]
}

export class TvShow {
  public id: string;
  public name: string;
  public year: number
  public average_rate: string
  public popularity: number
  public click_count: number
  public episode_count: number
  public season_count: number
  public trailers: Trailer[]
  public photos: Photo[]
  public companies: Company[]
  public likes: Like[]
  public celebs: Celebrity[]
}

export class Celebrity {
  public id: string;
  public name: string;
  public movies: Movie[]
  public photos: Photo[]
  public tv_shows: TvShow[]
}


export class SearchDTO {
  public keyword: string
  public movies: MovieDTO[]
  public tv_shows: TvShowDTO[]
  public celebrities: CelebrityDTO[]
}

export class CelebrityDTO {
  public id: string;
  public name: string;
  public movies: MovieDTO[]
  public photos: Photo[]
  public tv_shows: TvShowDTO[]
}
