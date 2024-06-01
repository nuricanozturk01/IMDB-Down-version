import {Component, Input, OnInit} from '@angular/core';
import {SearchService} from "../services/search.service";
import {MovieDTO, TvShowDTO, WatchListDTO} from "../../dto/dtos";

@Component({
  selector: 'app-watch-list',
  templateUrl: './watch-list.component.html',
  styleUrls: ['./watch-list.component.css']
})
export class WatchListComponent implements OnInit {

  movies: MovieDTO[] = [];
  tvShows: TvShowDTO[] = [];
  @Input() movie: MovieDTO;
  @Input() tvShow: TvShowDTO;

  constructor(private searchService: SearchService) { }

  ngOnInit(): void {
    this.fetchData();
  }

  private fetchData() {
    this.searchService.getWatchList().subscribe((response: WatchListDTO) => {
      this.movies = response.movies;
      this.tvShows = response.tv_shows;
    });
  }

  onRemoveFromWatchList(movie: MovieDTO) {
    this.movies = this.movies.filter(m => m.id !== movie.id);
  }

  onRemoveFromWatchListTvShow(tvShow: TvShowDTO) {
    this.tvShows = this.tvShows.filter(t => t.id !== tvShow.id);
  }
}
