import {CanActivateFn, Router} from '@angular/router';
import {inject} from "@angular/core";


export const authGuard: CanActivateFn = (route, state) => {
  const email = localStorage.getItem('email');
  const id = localStorage.getItem('id');
  const router = inject(Router);

  if (email) {
    return true;
  } else {
    router.navigate(['/sign-in']); // Sign-in sayfasına yönlendirme
    return false;
  }

};
