import {Component, OnInit} from '@angular/core';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {SearchService} from '../services/search.service';
import {CityDTO, CountryDTO, RegisterDTO} from '../../dto/dtos';
import {Router} from "@angular/router";
import {AuthenticationService} from "../services/authentication.service";
import {MessageService} from "../services/message.service";

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {
  registerForm: FormGroup;
  countries: CountryDTO[] = [];
  cities: CityDTO[] = [];
  submitted = false;

  constructor(private searchService: SearchService,
              private authService: AuthenticationService,
              private messageService: MessageService,
              private fb: FormBuilder, private router: Router) {
  }

  ngOnInit(): void {
    this.registerForm = this.fb.group({
      firstName: ['', Validators.required],
      lastName: ['', Validators.required],
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required, Validators.pattern('^(?=.*\\d)(?=.*[^a-zA-Z\\d])(?=.{8,}).*$')]],
      confirmPassword: ['', [Validators.required, Validators.pattern('^(?=.*\\d)(?=.*[^a-zA-Z\\d])(?=.{8,}).*$')]],
      country: ['', Validators.required],
      city: ['', Validators.required]
    }, {
      validator: this.confirmPasswordValidator('password', 'confirmPassword')
    });

    this.loadCountries();
  }

  loadCountries(): void {
    this.searchService.getCountries().subscribe((response: CountryDTO[]) => {
      this.countries = response;
      if (this.countries.length) {
        this.registerForm.patchValue({country: this.countries[0].CountryName});
        this.loadCities();
      }
    });
  }

  onCountryChange(): void {
    this.loadCities();
  }

  loadCities(): void {
    const selectedCountry = this.registerForm.get('country').value;
    if (selectedCountry) {
      this.searchService.getCitiesByCountryName(selectedCountry).subscribe((response: CityDTO[]) => {
        this.cities = response;
        if (this.cities.length) {
          this.registerForm.patchValue({city: this.cities[0].CityName});
        }
      });
    }
  }

  confirmPasswordValidator(password: string, confirmPassword: string) {
    return (formGroup: FormGroup) => {
      const passwordControl = formGroup.controls[password];
      const confirmPasswordControl = formGroup.controls[confirmPassword];

      if (!passwordControl || !confirmPasswordControl) {
        return null;
      }

      if (confirmPasswordControl.errors && !confirmPasswordControl.errors.mustMatch) {
        return null;
      }

      if (passwordControl.value !== confirmPasswordControl.value) {
        confirmPasswordControl.setErrors({mustMatch: true});
      } else {
        confirmPasswordControl.setErrors(null);
      }
    };
  }

  onSubmit() {
    if (this.registerForm.valid) {
      const registrationData = this.registerForm.value;
      this.submitted = true;

      const dto = new RegisterDTO();
      dto.first_name = registrationData.firstName;
      dto.last_name = registrationData.lastName;
      dto.email = registrationData.email;
      dto.password = registrationData.password;
      dto.country = registrationData.country;
      dto.city = registrationData.city;

      this.authService.register(dto).subscribe(response => {
        if (response) {
          this.messageService.showSuccess('Registration', 'Registration successful! Please login.');
          this.router.navigate(['/success-page'])
        } else {
          this.messageService.showError('Registration', 'Registration failed! Please control your data.');
        }
      });


    } else {
      console.error('Form is not valid:', this.registerForm.errors);
      alert('Please correct the errors in the form.');
    }
  }
}
