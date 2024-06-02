import {Component} from '@angular/core';
import {Router} from "@angular/router";
import {SearchService} from "../services/search.service";
import {CelebrityDTO, MovieDTO, SearchDTO, TvShowDTO} from "../../dto/dtos";
import {TranslateService} from "@ngx-translate/core";


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
  selectedOption: string = 'All';
  options: string[] = ['All', 'Titles', 'TV Episodes', 'Celebs', 'Companies'];


  constructor(private translate: TranslateService, private router: Router, private service: SearchService) {
    translate.addLangs(['en', 'tr']);
    translate.setDefaultLang('en');

    const savedLang = localStorage.getItem('language');
    const browserLang = savedLang || translate.getBrowserLang();
    translate.use(browserLang.match(/en|tr/) ? browserLang : 'en');
  }


  switchLanguage(language: string) {
    this.translate.use(language);
    localStorage.setItem('language', language);
  }

  selectOption(option: string) {
    this.selectedOption = option;
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

      if (this.selectedOption === 'Titles') {
        this.celebs = [];
      }

      if (this.selectedOption === 'TV Episodes') {
        this.movies = [];
        this.celebs = [];
      }

      if (this.selectedOption === 'Celebs') {
        this.movies = [];
        this.tvShows = [];
      }

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
