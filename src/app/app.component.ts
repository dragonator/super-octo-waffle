import { Component } from '@angular/core';
import { Router }    from '@angular/router';

import { AuthService } from '@app/_services/auth.service';
import { User }        from '@app/_models';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  public currentUser: User;

  constructor(private router: Router, private auth: AuthService) {
    this.auth.currentUser.subscribe(x => this.currentUser = x);
  }

  public login() {
    this.auth.login()
  }

  public logout() {
    this.auth.logout();
    this.router.navigate(['/home']);
  }
}
