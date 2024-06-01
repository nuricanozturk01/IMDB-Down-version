import {Component, OnDestroy, OnInit} from '@angular/core';
import {FormControl} from "@angular/forms";
import {MovieDTO, TvShowDTO} from "../../dto/dtos";
import {SearchService} from "../services/search.service";
import {MessageService} from "../services/message.service";
import {NgbActiveModal} from "@ng-bootstrap/ng-bootstrap";

@Component({
  selector: 'app-rate',
  templateUrl: './rate.component.html',
  styleUrls: ['./rate.component.css']
})
export class RateComponent implements OnInit, OnDestroy {
  ctrl = new FormControl(null);
  movie: MovieDTO;
  tvShow: TvShowDTO;
  mediaType: string;

  constructor(public activeModal: NgbActiveModal,
              private service: SearchService, private messageService: MessageService) {

  }

  ngOnDestroy() {
    this.ctrl.reset();
    this.movie = null;
    this.tvShow = null;
    this.mediaType = null;
  }

  ngOnInit() {

    this.ctrl.valueChanges.subscribe(value => {
      if (value !== null) {
        this.rateMedia(value);
      }
    });
  }

  rateMedia(value: number) {
    if (this.mediaType === 'movie') {
      this.service.rateMovie(this.movie.id, value).subscribe(response => {
        if (response === "Already rated!") {
          this.messageService.showWarning("Rate", response);
        } else this.messageService.showSuccess("Rate", response);
      });
    }
    if (this.mediaType === 'tv') {
      this.service.rateTvShow(this.tvShow.id, value).subscribe(response => {
        if (response === "Already rated!") {
          this.messageService.showWarning("Rate", response);
        } else this.messageService.showSuccess("Rate", response);
      });
    }
  }

  closeModal() {
    this.activeModal.close();
  }
}
