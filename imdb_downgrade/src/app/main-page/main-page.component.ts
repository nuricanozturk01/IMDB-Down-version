import {Component, OnInit} from '@angular/core';
import {Router} from "@angular/router";
import {SearchService} from "../services/search.service";
import {MovieDTO} from "../../dto/dtos";
import {NgbModal} from "@ng-bootstrap/ng-bootstrap";
import {RateComponent} from "../rate/rate.component";

@Component({
  selector: 'app-main-page',
  templateUrl: './main-page.component.html',
  styleUrls: ['./main-page.component.css']
})
export class MainPageComponent implements OnInit {
  movies: MovieDTO[] = [];
  movieSlides: MovieDTO[][] = [];

  constructor(private route: Router, private service: SearchService, private modal: NgbModal) {
  }

  ngOnInit(): void {
    this.service.findAllMovies().subscribe((response: MovieDTO[]) => {
      this.movies = response;
      this.movieSlides = this.chunkArray(this.movies, 4);
    });
  }

  private chunkArray(myArray, chunk_size) {
    let results = [];

    for (let i = 0; i < myArray.length; i += chunk_size)
      results.push(myArray.slice(i, i + chunk_size));

    return results;
  }

  clickItem(movie: MovieDTO) {
    this.service.findMovieDetails(movie.id).subscribe((response: MovieDTO) => {
      console.log("DTO: ", response);
      this.route.navigate(['/details', {movie: JSON.stringify(response)}]);
    });
  }

  clickAddOnWatchList(movie: MovieDTO) {
    this.service.addOnWatchList(movie.id).subscribe((response: string) => {
      console.log(response);
    })
  }

  clickRateMovie(movie: MovieDTO) {
    const modalRef = this.modal.open(RateComponent, {
      size: 'lg',
      backdrop: 'static',
    });
    modalRef.componentInstance.showModal = true;
  }
}
