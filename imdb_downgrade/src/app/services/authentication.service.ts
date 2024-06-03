import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {catchError, map, Observable} from "rxjs";
import {LoginDTO} from "../../dto/LoginDTO";
import {RegisterDTO} from "../../dto/dtos";
import {Router} from "@angular/router";
import {REQUEST_GOOGLE_LOGIN, REQUEST_LOGIN, REQUEST_LOGOUT, REQUEST_REGISTER} from "./connection";

@Injectable({
  providedIn: 'root'
})
export class AuthenticationService {

  constructor(private http: HttpClient, private router: Router) {
  }


  loginWithGoogle() {
    window.location.href = REQUEST_GOOGLE_LOGIN;
  }


  login(loginModel: LoginDTO): Observable<any> {
    return this.http.post(REQUEST_LOGIN, loginModel, {withCredentials: true}).pipe(
      map((response: any) => {
        localStorage.setItem("email", response.data.email)
        localStorage.setItem("first_name", response.data.first_name)
        localStorage.setItem("last_name", response.data.last_name)
        localStorage.setItem("id", response.data.id)
        return response;
      }),
      catchError((error: any) => {
        return [false]
      })
    );
  }

  register(dto: RegisterDTO) {
    return this.http.post(REQUEST_REGISTER, dto).pipe(
      map((response: any) => {
        return response.status_code === 201;
      }),
      catchError((error: any) => {
          return [false];
        }
      ));
  }

  logout(): Observable<boolean> {
    return this.http.post(REQUEST_LOGOUT, {}, {withCredentials: true}).pipe(
      map((response: any) => {
        localStorage.clear();
        return true
      }),
      catchError((error: any) => {
        return [false];
      })
    );
  }
}
