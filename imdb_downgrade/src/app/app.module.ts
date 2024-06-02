import {NgModule} from '@angular/core';
import {BrowserModule} from '@angular/platform-browser';

import {AppComponent} from './app.component';
import {NgbModule, NgbRatingModule} from '@ng-bootstrap/ng-bootstrap';
import {LoginComponent} from './login/login.component';
import {AppRoutingModule} from './app-routing.module';
import {NavbarComponent} from './navbar/navbar.component';
import {MainPageComponent} from './main-page/main-page.component';
import {RegisterComponent} from './register/register.component';
import {DetailsComponent} from './details/details.component';
import {HttpClientModule} from "@angular/common/http";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {IntroComponent} from './intro/intro.component';
import {RateComponent} from './rate/rate.component';
import {MovieCardComponent} from './movie-card/movie-card.component';
import {WatchListComponent} from './watch-list/watch-list.component';
import {TvCardComponent} from './tv-card/tv-card.component';
import {ToastrModule} from "ngx-toastr";
import { TvDetailsComponent } from './tv-details/tv-details.component';
import { CelebDetailsComponent } from './celeb-details/celeb-details.component';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    NavbarComponent,
    MainPageComponent,
    RegisterComponent,
    DetailsComponent,
    IntroComponent,
    RateComponent,
    MovieCardComponent,
    WatchListComponent,
    TvCardComponent,
    TvDetailsComponent,
    CelebDetailsComponent
  ],
    imports: [
        HttpClientModule,
        BrowserModule,
        NgbModule,
        AppRoutingModule,
        ReactiveFormsModule,
        NgbRatingModule,
        ToastrModule.forRoot({
            timeOut: 3000,
            closeButton: true,
            progressBar: true,
            progressAnimation: 'increasing',
            preventDuplicates: true,
            positionClass: 'toast-top-right',
        }),
        FormsModule,
    ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {
}
