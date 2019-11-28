import { HomeComponent }        from './home/home.component';
import { RouterModule, Routes } from '@angular/router';
import { NgModule }             from '@angular/core';
import { AuthGuardService }     from './auth-guard.service';
import { CallbackComponent }    from './callback/callback.component';
import { PinnedReposComponent } from './pinned-repos/pinned-repos.component';

const routes: Routes = [
  { path: '',      redirectTo: 'home',                         pathMatch: 'full' },
  { path: 'login', redirectTo: 'dev-a638a1un.auth0.com/login', pathMatch: 'full'},

  { path: 'login/callback',            component: CallbackComponent },
  { path: 'home',                      component: HomeComponent },
  { path: 'pinnedRepos/:organization', component: PinnedReposComponent,  canActivate: [AuthGuardService] },
];

@NgModule({
  imports: [ RouterModule.forRoot(routes) ],
  exports: [ RouterModule ]
})
export class AppRoutingModule { }
