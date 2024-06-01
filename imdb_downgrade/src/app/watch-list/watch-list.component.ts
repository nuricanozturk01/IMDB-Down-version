import {Component, Input, OnInit} from '@angular/core';
import {SearchService} from "../services/search.service";
import {MovieDTO, TvShowDTO, WatchListDTO} from "../../dto/dtos";
import {MessageService} from "../services/message.service";

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

  constructor(private searchService: SearchService, private messageService: MessageService) {
  }

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
    this.messageService.showSuccess("Success", "Movie removed from watch list");
  }

  onRemoveFromWatchListTvShow(tvShow: TvShowDTO) {
    this.tvShows = this.tvShows.filter(t => t.id !== tvShow.id);
    this.messageService.showSuccess("Success", "TV show removed from watch list");
  }
}
