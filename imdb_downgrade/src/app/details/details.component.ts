import {Component, OnInit} from '@angular/core';
import {MovieDTO} from "../../dto/dtos";
import {ActivatedRoute} from "@angular/router";
import {SearchService} from "../services/search.service";
import {NgbModal} from "@ng-bootstrap/ng-bootstrap";
import {DomSanitizer} from "@angular/platform-browser";

@Component({
  selector: 'app-details',
  templateUrl: './details.component.html',
  styleUrls: ['./details.component.css']
})
export class DetailsComponent implements OnInit {
  movie: MovieDTO;

  constructor(private route: ActivatedRoute, private service: SearchService, private modal: NgbModal, private sanitizer: DomSanitizer) {

  }

  ngOnInit(): void {
    this.route.params.subscribe(params => {
      if (params['movie']) {
        this.movie = JSON.parse(params['movie']);
        console.log("Movie: ", this.movie);
      }
    });
    this.movie.trailers.forEach(trailer => {
      trailer.safeUrl = this.sanitizer.bypassSecurityTrustResourceUrl(trailer.url);
    })
  }
}
