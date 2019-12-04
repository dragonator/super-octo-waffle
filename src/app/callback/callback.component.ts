import { Component, OnInit } from '@angular/core';
import { Router }            from '@angular/router';

import { AuthService } from '@app/_services/auth.service';

@Component({
  selector:    'app-callback',
  templateUrl: './callback.component.html',
  styleUrls:  ['./callback.component.scss']
})
export class CallbackComponent implements OnInit {

  constructor(private router: Router, private auth: AuthService) {}

  ngOnInit() {
    this.auth.handleAuthentication();
    this.router.navigate(['/home']);
  }
}
