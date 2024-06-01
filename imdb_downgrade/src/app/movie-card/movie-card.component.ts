import {Component, EventEmitter, Input, Output} from '@angular/core';
import {MovieDTO} from "../../dto/dtos";
import {SearchService} from "../services/search.service";
import {Router} from "@angular/router";

@Component({
  selector: 'app-movie-card',
  templateUrl: './movie-card.component.html',
  styleUrls: ['./movie-card.component.css']
})
export class MovieCardComponent {
  @Input() movie: MovieDTO;
  @Output() removeFromWatchList = new EventEmitter<MovieDTO>();

  constructor(private searchService: SearchService, private route: Router) {
  }

  clickRemoveFromWatchList(movie: MovieDTO) {
    this.searchService.removeOnWatchList(movie.id, "tv_show").subscribe((response: string) => {
      console.log(response);
      this.removeFromWatchList.emit(movie);
    });
  }

  clickDetails(movie: MovieDTO) {
    this.searchService.findMovieDetails(movie.id).subscribe((response: MovieDTO) => {
      console.log("DTO: ", response);
      this.route.navigate(['/details', {movie: JSON.stringify(response)}]);
    });

  }
}