import {Component, OnInit} from '@angular/core';
import {AuthenticationService} from "../services/authentication.service";
import {FormBuilder, FormGroup, Validators} from "@angular/forms";
import {LoginDTO} from "../../dto/LoginDTO";
import {MessageService} from "../services/message.service";
import {Router} from "@angular/router";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  loginForm!: FormGroup

  constructor(private service: AuthenticationService, private formBuilder: FormBuilder,
              private messageService: MessageService, private router: Router) {

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
        if (response.message === "Success!") {
          this.messageService.showSuccess("Login", "Login successful!")
          setTimeout(() => {
            this.router.navigate(['/'])
          }, 500)
        } else {
          this.messageService.showError("Login", "Invalid credentials!")
        }
      });
  }

  loginWithGoogle() {
    this.service.loginWithGoogle()
  }
}
