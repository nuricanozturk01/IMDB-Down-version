import {Component, OnInit} from '@angular/core';
import {TvShowDTO} from "../../dto/dtos";
import {ActivatedRoute} from "@angular/router";
import {DomSanitizer, SafeResourceUrl} from "@angular/platform-browser";

@Component({
  selector: 'app-tv-details',
  templateUrl: './tv-details.component.html',
  styleUrls: ['./tv-details.component.css']
})
export class TvDetailsComponent implements OnInit {
  tvShow: TvShowDTO
  safeUrl: SafeResourceUrl

  constructor(private route: ActivatedRoute, private sanitizer: DomSanitizer) {

  }

  ngOnInit(): void {
    this.tvShow = null;
    this.safeUrl = null;
    this.route.params.subscribe(params => {
      if (params['tvShow']) {
        this.tvShow = JSON.parse(atob(params['tvShow']));
        this.safeUrl = this.sanitizer.bypassSecurityTrustResourceUrl(this.tvShow.trailers[0].url);
      }
    });
  }
}
