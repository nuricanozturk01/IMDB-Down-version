import {Component, OnInit} from '@angular/core';
import {TvShowDTO} from "../../dto/dtos";
import {ActivatedRoute} from "@angular/router";
import {SearchService} from "../services/search.service";
import {NgbModal} from "@ng-bootstrap/ng-bootstrap";
import {DomSanitizer} from "@angular/platform-browser";

@Component({
  selector: 'app-tv-details',
  templateUrl: './tv-details.component.html',
  styleUrls: ['./tv-details.component.css']
})
export class TvDetailsComponent implements OnInit {
  tvShow: TvShowDTO

  constructor(private route: ActivatedRoute, private service: SearchService, private modal: NgbModal, private sanitizer: DomSanitizer) {

  }

  ngOnInit(): void {
    this.route.params.subscribe(params => {
      if (params['tvShow']) {
        this.tvShow = JSON.parse(params['tvShow']);
        this.tvShow.trailers.forEach(trailer => {
          trailer.safeUrl = this.sanitizer.bypassSecurityTrustResourceUrl(trailer.url);
        })
      }
    });
  }
}
