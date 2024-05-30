import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {catchError, map, Observable} from "rxjs";
import {REQUEST_SEARCH} from "./connection";
import {CookieService} from "ngx-cookie-service";

@Injectable({
  providedIn: 'root'
})
export class SearchService {

  constructor(private http: HttpClient, private cookieService: CookieService) {
  }

  search(keyword: string): Observable<any> {
    console.log(this.cookieService.getAll())
    return this.http.get(REQUEST_SEARCH(keyword), {withCredentials: true}).pipe(
      map((response: any) => {
        return response;
      }),
      catchError((error: any) => {
          return [false];
        }
      ));
  }
}
