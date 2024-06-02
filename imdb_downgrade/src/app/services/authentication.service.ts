import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {catchError, map, Observable} from "rxjs";
import {LoginDTO} from "../../dto/LoginDTO";

@Injectable({
  providedIn: 'root'
})
export class AuthenticationService {

  constructor(private http: HttpClient) {
  }


  loginWithGoogle() {
    window.location.href = 'http://localhost:5050/api/auth/google/login';
  }


  login(loginModel: LoginDTO): Observable<any> {
    return this.http.post('http://localhost:5050/api/v1/auth/login', loginModel).pipe(
      map((response: any) => {
        console.log(response);
        return response;
      }),
      catchError((error: any) => {
          return [false];
        }
      ));
  }
}
