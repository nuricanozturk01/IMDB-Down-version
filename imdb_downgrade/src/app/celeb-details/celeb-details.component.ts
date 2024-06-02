import {Component, OnDestroy, OnInit} from '@angular/core';
import {CelebrityDTO, MovieDTO, TvShowDTO} from "../../dto/dtos";
import {ActivatedRoute} from "@angular/router";
import {DomSanitizer, SafeResourceUrl} from "@angular/platform-browser";
import {SearchService} from "../services/search.service";

@Component({
  selector: 'app-celeb-details',
  templateUrl: './celeb-details.component.html',
  styleUrls: ['./celeb-details.component.css']
})
export class CelebDetailsComponent implements OnInit, OnDestroy {
  celebrity: CelebrityDTO;
  movies: MovieDTO[] = [];
  tvShows: TvShowDTO[] = [];
  movieSafeUrlIfExists: SafeResourceUrl;
  tvShowSafeUrlIfExists: SafeResourceUrl;

  constructor(private route: ActivatedRoute, private sanitizer: DomSanitizer, private service: SearchService) {

  }

  ngOnDestroy() {
    localStorage.removeItem("celebrity_id");
    this.celebrity = null;
    this.movies = [];
    this.tvShows = [];
  }

  ngOnInit(): void {
    console.log('CelebDetailsComponent initialized');

    this.route.params.subscribe(params => {
      if (params['celebrity']) {
        this.celebrity = JSON.parse(atob(params['celebrity']));
        this.movies = this.celebrity.movies;
        this.tvShows = this.celebrity.tv_shows;


        this.movieSafeUrlIfExists = this.movies ? this.sanitizer.bypassSecurityTrustResourceUrl(this.movies[0].trailers[0].url) : null;
        this.tvShowSafeUrlIfExists = this.tvShows ? this.sanitizer.bypassSecurityTrustResourceUrl(this.tvShows[0].trailers[0].url) : null;
      }
    });
  }

}
