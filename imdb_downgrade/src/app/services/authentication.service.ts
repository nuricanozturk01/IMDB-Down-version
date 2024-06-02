import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {catchError, map, Observable} from "rxjs";
import {LoginDTO} from "../../dto/LoginDTO";
import {RegisterDTO} from "../../dto/dtos";
import {Router} from "@angular/router";

@Injectable({
  providedIn: 'root'
})
export class AuthenticationService {

  constructor(private http: HttpClient, private router: Router) {
  }


  loginWithGoogle() {
    window.location.href = 'http://localhost:5050/api/auth/google/login';
  }


  login(loginModel: LoginDTO): Observable<any> {
    return this.http.post(`http://localhost:5050/api/auth/login`, loginModel, {withCredentials: true}).pipe(
      map((response: any) => {
        console.log('Response:', response);
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
    return this.http.post('http://localhost:5050/api/auth/register', dto).pipe(
      map((response: any) => {
        return response.status_code === 201;
      }),
      catchError((error: any) => {
          return [false];
        }
      ));
  }

  logout(): Observable<boolean> {
    return this.http.post('http://localhost:5050/api/auth/logout', {}, {withCredentials: true}).pipe(
      map((response: any) => {
        if (response.status_code === 200) {
          localStorage.clear();
          return true;
        }
        return false;
      }),
      catchError((error: any) => {
        return [false];
      })
    );
  }
}
