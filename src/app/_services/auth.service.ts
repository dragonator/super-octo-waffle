import { Injectable }  from '@angular/core';
import { Router }      from '@angular/router';
import { BehaviorSubject, Observable } from 'rxjs';

import { environment } from '@environments/environment';
import { User, Token } from '@app/_models';

import * as auth0 from 'auth0-js';

(window as any).global = window;

@Injectable()
export class AuthService {

  private currentUserSubject: BehaviorSubject<User>;
  public  currentUser:        Observable<User>;

  constructor(public router: Router)  {
    this.currentUserSubject = new BehaviorSubject<User>(JSON.parse(localStorage.getItem('currentUser')));
    this.currentUser = this.currentUserSubject.asObservable();
  }

  auth0 = new auth0.WebAuth({
    clientID:     environment.clientId,
    domain:       environment.domain,
    audience:     environment.audience,
    responseType: 'token id_token',
    redirectUri:  environment.callback,
    scope:        'openid profile'
  });

  public get currentUserValue(): User {
    return this.currentUserSubject.value;
  }

  public login(): void {
    this.auth0.authorize();
  }

  public handleAuthentication(): void {
    this.auth0.parseHash((err, authResult) => {
      if (err) {
        return console.log(err);
      }

      this.auth0.client.userInfo(authResult.accessToken, (err, auth0User) => {
        if (err) {
          return console.log(err);
        }

        var expiresAt = JSON.stringify((authResult.expiresIn * 1000) + new Date().getTime());
        var token     = new Token(authResult.accessToken, authResult.idToken, authResult.expiresAt);
        var user      = new User(auth0User.nickname, token);
        localStorage.setItem('currentUser', JSON.stringify(user));
        this.currentUserSubject.next(user);
      });

      this.router.navigate(['/home']);
    });
  }

  public logout(): void {
    localStorage.removeItem('currentUser');
    this.currentUserSubject.next(null);
    this.router.navigate(['/']);
  }
}
