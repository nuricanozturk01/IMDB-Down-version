import {Component, EventEmitter, Input, Output} from '@angular/core';
import {TvShowDTO} from "../../dto/dtos";
import {SearchService} from "../services/search.service";
import {Router} from "@angular/router";

@Component({
  selector: 'app-tv-card',
  templateUrl: './tv-card.component.html',
  styleUrls: ['./tv-card.component.css']
})
export class TvCardComponent {
  @Input() tvShow: TvShowDTO;
  @Output() removeFromWatchList = new EventEmitter<TvShowDTO>();

  constructor(private searchService: SearchService, private route: Router) {
  }

  clickRemoveFromWatchList(tvShow: TvShowDTO) {
    this.searchService.removeOnWatchList(tvShow.id, "tv_show").subscribe((response: string) => {
      this.removeFromWatchList.emit(tvShow);
    });
  }

  clickShowDetailsTvShow(tvShow: TvShowDTO) {
    this.searchService.findTvShowDetails(tvShow.id).subscribe((response: TvShowDTO) => {
      this.route.navigate(['/tv-details', {tvShow: btoa(JSON.stringify(response))}]);
    });
  }
}
