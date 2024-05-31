import {Component} from '@angular/core';
import {Router} from "@angular/router";
import {SearchService} from "../services/search.service";
import {SearchDTO} from "../../dto/dtos";

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent {
  results = [
    {title: 'Cannes Film Festival', image: 'assets/cannes.jpg'},
    {title: 'Harry Wild', image: 'assets/harry-wild.jpg'},
    {title: 'Harry Potter and the Sorcerer\'s Stone', image: 'assets/harry-potter.jpg'},
    {title: 'Harry Melling', image: 'assets/harry-melling.jpg'},
    {title: 'Harrison Ford', image: 'assets/harrison-ford.jpg'}
  ];

  filteredResults = [];

  constructor(private router: Router, private service: SearchService) {
  }

  handleSignInClick() {
    this.router.navigate(['/sign-in']);
  }

  onSearch(query: string) {
    if (query.length >= 3) {
     /* this.filteredResults = this.results.filter(result =>
        result.title.toLowerCase().includes(query.toLowerCase())
      );
*/
      this.service.search(query).subscribe((response: SearchDTO) => {
        console.log(response);
      });

    } else {
      this.filteredResults = [];
    }
  }
}
