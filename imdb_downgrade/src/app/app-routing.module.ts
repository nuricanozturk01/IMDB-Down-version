import {NgModule} from '@angular/core';
import {RouterModule, Routes} from "@angular/router";
import {LoginComponent} from "./login/login.component";
import {RegisterComponent} from "./register/register.component";
import {DetailsComponent} from "./details/details.component";
import {WatchListComponent} from "./watch-list/watch-list.component";
import {TvDetailsComponent} from "./tv-details/tv-details.component";
import {CelebDetailsComponent} from "./celeb-details/celeb-details.component";
import {authGuard} from "./guards/auth.guard";
import {MainPageComponent} from "./main-page/main-page.component";


const routes: Routes = [
  {path: 'sign-in', component: LoginComponent},
  {path: '', component: LoginComponent},
  {path: 'register', component: RegisterComponent},
  {path: 'details', component: DetailsComponent, canActivate: [authGuard]},
  {path: 'tv-details', component: TvDetailsComponent, canActivate: [authGuard]},
  {path: 'celebrity', component: CelebDetailsComponent, canActivate: [authGuard]},
  {path: 'watch-list', component: WatchListComponent, canActivate: [authGuard]},
  {path: 'main-page', component: MainPageComponent, canActivate: [authGuard]},
  {path: '**', redirectTo: '/sign-in'}

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
