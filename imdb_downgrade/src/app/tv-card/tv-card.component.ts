import {Component, EventEmitter, Input, Output} from '@angular/core';
import {TvShowDTO} from "../../dto/dtos";
import {SearchService} from "../services/search.service";

@Component({
  selector: 'app-tv-card',
  templateUrl: './tv-card.component.html',
  styleUrls: ['./tv-card.component.css']
})
export class TvCardComponent {
  @Input() tvShow: TvShowDTO;
  @Output() removeFromWatchList = new EventEmitter<TvShowDTO>();

  constructor(private searchService: SearchService) {
  }

  clickRemoveFromWatchList(tvShow: TvShowDTO) {
    this.searchService.removeOnWatchList(tvShow.id, "tv_show").subscribe((response: string) => {
      this.removeFromWatchList.emit(tvShow);
    });
  }

  clickShowDetailsTvShow(tvShow: TvShowDTO) {

  }
}
