import {Component, OnInit} from '@angular/core';
import {MovieDTO} from "../../dto/dtos";
import {ActivatedRoute} from "@angular/router";
import {DomSanitizer, SafeResourceUrl} from "@angular/platform-browser";

@Component({
  selector: 'app-details',
  templateUrl: './details.component.html',
  styleUrls: ['./details.component.css']
})
export class DetailsComponent implements OnInit {
  movie: MovieDTO;
  safeUrl: SafeResourceUrl

  constructor(private route: ActivatedRoute, private sanitizer: DomSanitizer) {

  }

  ngOnInit(): void {
    this.movie = null;
    this.safeUrl = null;

    this.route.params.subscribe(params => {
      if (params['movie']) {
        this.movie = JSON.parse(atob(params['movie']));
        this.safeUrl = this.sanitizer.bypassSecurityTrustResourceUrl(this.movie.trailers[0].url);
      }
    });
  }
}
