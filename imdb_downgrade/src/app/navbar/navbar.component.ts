import {Component} from '@angular/core';
import {Router} from "@angular/router";
import {SearchService} from "../services/search.service";
import {CelebrityDTO, MovieDTO, SearchDTO, TvShowDTO} from "../../dto/dtos";
import {TranslateService} from "@ngx-translate/core";
import {AuthenticationService} from "../services/authentication.service";
import {MessageService} from "../services/message.service";

export const isLoggedIn = () => {
  return localStorage.getItem('email') !== null;
}

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
  options: string[] = [];

  constructor(private translate: TranslateService, private router: Router, private service: SearchService,
              private authService: AuthenticationService, private messageService: MessageService) {
    translate.addLangs(['en', 'tr']);
    translate.setDefaultLang('en');

    const savedLang = localStorage.getItem('language');
    const browserLang = savedLang || translate.getBrowserLang();
    translate.use(browserLang.match(/en|tr/) ? browserLang : 'en');
    this.translate.get('OPTIONS').subscribe(value => {
      this.options = [
        value['ALL'],
        value['TITLES'],
        value['TV_EPISODES'],
        value['CELEBS']
      ]
    })
  }


  getEmail() {
    try {
      return localStorage.getItem('email').substring(0, localStorage.getItem('email').indexOf('@'));
    } catch (e) {
      return '';
    }
  }

  switchLanguage(language: string) {
    this.translate.use(language);
    localStorage.setItem('language', language);
  }

  selectOption(option: string) {
    this.selectedOption = option;
    console.log(this.selectedOption);
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


      if (this.selectedOption === 'Titles' || this.selectedOption === 'Başlıklar') {// Movies
        this.celebs = [];
        this.tvShows = [];
      } else if (this.selectedOption === 'TV Episodes' || this.selectedOption === 'TV Bölümleri') { // TV Shows
        this.movies = [];
        this.celebs = [];
      } else if (this.selectedOption === 'Celebs' || this.selectedOption === 'Ünlüler') { // Celebs
        console.log('Celebs');
        this.movies = [];
        this.tvShows = [];
      }

    } else {
      console.log('Query length must be greater than 2');
      this.isFoundResults = false;
      this.movies = [];
      this.tvShows = [];
      this.celebs = [];
    }
  }

  handleClickCelebrity(celebrity: CelebrityDTO) {
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

  logout() {
    this.router.navigate(['/sign-in']).then(() => {
        this.authService.logout().subscribe((response: boolean) => {
          this.messageService.showSuccess("Logout", "Logout successful!")
        })
      }
    );
  }

  protected readonly isLoggedIn = isLoggedIn;

  handleClickWatchList() {
    if (isLoggedIn()) {
      this.router.navigate(['/watch-list']);
    } else {
      this.messageService.showWarning("Error", "You must be logged in to see your watchlist!");
      this.router.navigate(['/sign-in']);
    }
  }
}
