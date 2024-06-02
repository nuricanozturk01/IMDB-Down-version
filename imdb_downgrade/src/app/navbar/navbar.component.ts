import {Component} from '@angular/core';
import {Router} from "@angular/router";
import {SearchService} from "../services/search.service";
import {CelebrityDTO, MovieDTO, SearchDTO, TvShowDTO} from "../../dto/dtos";


@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent {
  isFoundResults = false;
  movies: MovieDTO[] = [];
  tvShows: TvShowDTO[] = [];
  celebs: CelebrityDTO[] = [];

  constructor(private router: Router, private service: SearchService) {
  }

  onSearch(query: string) {
    if (query.length >= 2) {
      this.service.search(query).subscribe((response: SearchDTO) => {

        if (response.movies && response.movies.length > 0) {
          this.movies = response.movies;
        }
        if (response.tv_shows && response.tv_shows.length > 0) {
          this.tvShows = response.tv_shows;
        }

        if (response.celebrities && response.celebrities.length > 0) {
          this.celebs = response.celebrities;
        }
      });

    } else {
      this.isFoundResults = false;
      this.movies = [];
      this.tvShows = [];
      this.celebs = [];
    }
  }

  handleClickCelebrity(celebrity: CelebrityDTO) {
    //localStorage.setItem("celebrity_id", celebrity.id);
    this.isFoundResults = false;
    this.movies = [];
    this.tvShows = [];
    this.celebs = [];

    this.service.findCelebrityById(celebrity.id).subscribe((response: CelebrityDTO) => {
      this.router.navigate(['/celebrity', {celebrity: btoa(JSON.stringify(response))}]);
    });


    /* this.router.navigate(['/celebrity']).then(success => {
       if (success) {
         console.log('Navigation successful!');
       } else {
         console.log('Navigation failed!');
         console.error('Navigation failed to /celebrity-details');
       }
     }).catch(err => {
       console.error('Navigation error:', err);
     });*/
  }


  handleClickTvShow(tvShow: TvShowDTO) {
    this.isFoundResults = false;
    this.movies = [];
    this.tvShows = [];
    this.celebs = [];
    this.router.navigate(['/tv-details', {tvShow: btoa(JSON.stringify(tvShow))}]);
  }

  handleClickMovie(movie: MovieDTO) {
    this.isFoundResults = false;
    this.movies = [];
    this.tvShows = [];
    this.celebs = [];
    this.router.navigate(['/details', {movie: btoa(JSON.stringify(movie))}]);
  }
}
