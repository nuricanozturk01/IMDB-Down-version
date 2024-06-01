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
    filteredResults = [];
    movies: MovieDTO[] = [];
    tvShows: TvShowDTO[] = [];
    celebs: CelebrityDTO[] = [];

    constructor(private router: Router, private service: SearchService) {
    }

    handleSignInClick() {
        this.router.navigate(['/sign-in']);
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
}
