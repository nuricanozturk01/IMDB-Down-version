import {Component, OnInit} from '@angular/core';
import {Router} from "@angular/router";
import {SearchService} from "../services/search.service";
import {MovieDTO, TvShowDTO} from "../../dto/dtos";
import {NgbModal, NgbNavConfig} from "@ng-bootstrap/ng-bootstrap";
import {RateComponent} from "../rate/rate.component";
import {MessageService} from "../services/message.service";

@Component({
  selector: 'app-main-page',
  templateUrl: './main-page.component.html',
  styleUrls: ['./main-page.component.css']
})
export class MainPageComponent implements OnInit {
  movies: MovieDTO[] = [];
  movieSlides: MovieDTO[][] = [];
  tvShowSlides: TvShowDTO[][] = [];
  tvShows: TvShowDTO[] = [];

  constructor(private route: Router,
              private messageService: MessageService,
              private service: SearchService,
              config: NgbNavConfig,
              private modal: NgbModal) {
    config.destroyOnHide = false;
    config.roles = false;
  }

  ngOnInit(): void {
    this.service.getUserInfo().subscribe((response: any) => {
      if (response) {
        if (localStorage.getItem("email") === null || localStorage.getItem("email") === undefined) {
          localStorage.setItem("email", response.email)
          localStorage.setItem("first_name", response.first_name)
          localStorage.setItem("last_name", response.last_name)
          localStorage.setItem("id", response.id)
        }
      }
    })
    this.service.findAllMovies().subscribe((response: MovieDTO[]) => {
      this.movies = response;
      this.movieSlides = this.chunkArray(this.movies, 4);
    });

    this.service.findAllTvShows().subscribe((response: TvShowDTO[]) => {
      this.tvShows = response;
      this.tvShowSlides = this.chunkArray(this.tvShows, 4);
    });
  }

  private chunkArray(myArray, chunk_size) {
    let results = [];

    for (let i = 0; i < myArray.length; i += chunk_size)
      results.push(myArray.slice(i, i + chunk_size));

    return results;
  }

  clickMovieItem(movie: MovieDTO) {
    this.service.findMovieDetails(movie.id).subscribe((response: MovieDTO) => {
      this.route.navigate(['/details', {movie: btoa(JSON.stringify(response))}]);
    });
  }

  clickTvItem(movie: TvShowDTO) {
    this.service.findTvShowDetails(movie.id).subscribe((response: MovieDTO) => {
      this.route.navigate(['/tv-details', {tvShow: btoa(JSON.stringify(response))}]);
    });
  }

  clickAddOnWatchList(movie: MovieDTO) {
    this.service.addOnWatchList(movie.id).subscribe((response: string) => {
      this.messageService.showSuccess("Success", response);
    })
  }

  clickAddOnWatchListTv(tv: TvShowDTO) {
    this.service.addOnWatchList(tv.id, "tv_show").subscribe((response: string) => {
      this.messageService.showSuccess("Success", response);
    })
  }

  clickRateMovie(movie: MovieDTO) {
    const modalRef = this.modal.open(RateComponent, {
      size: 'lg',
      backdrop: 'static',
    });
    modalRef.componentInstance.movie = movie;
    modalRef.componentInstance.mediaType = "movie";
    modalRef.componentInstance.showModal = true;
  }

  clickRateTv(tvShow: TvShowDTO) {
    const modalRef = this.modal.open(RateComponent, {
      size: 'lg',
      backdrop: 'static',
    });
    modalRef.componentInstance.tvShow = tvShow;
    modalRef.componentInstance.mediaType = "tv";
    modalRef.componentInstance.showModal = true;
  }
}
