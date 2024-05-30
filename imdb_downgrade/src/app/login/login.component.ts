import {Component, OnInit} from '@angular/core';
import {AuthenticationService} from "../services/authentication.service";
import {FormBuilder, FormGroup, Validators} from "@angular/forms";
import {LoginDTO} from "../../dto/LoginDTO";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  loginForm!: FormGroup

  constructor(private service: AuthenticationService, private formBuilder: FormBuilder) {

  }

  ngOnInit() {
    this.loginForm = this.formBuilder.group({
      email: ['', Validators.required],
      password: ['', Validators.required]
    })
  }

  login() {
    this.service.login(new LoginDTO(this.loginForm.value.email, this.loginForm.value.password))
      .subscribe((response: any) => {
        console.log(response)
      })
  }

  loginWithGoogle() {
    this.service.loginWithGoogle()
  }
}
