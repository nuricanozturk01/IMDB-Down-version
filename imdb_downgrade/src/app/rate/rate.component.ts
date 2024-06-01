import {Component, OnInit} from '@angular/core';
import {FormControl} from "@angular/forms";

@Component({
  selector: 'app-rate',
  templateUrl: './rate.component.html',
  styleUrls: ['./rate.component.css']
})
export class RateComponent implements OnInit {
  ctrl = new FormControl(null);

  ngOnInit() {
    this.ctrl.valueChanges.subscribe(value => {
      if (value !== null) {
        this.rateMedia(value);
      }
    });
  }

  rateMedia(value: number) {
    console.log(value);
  }
}
