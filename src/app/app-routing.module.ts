import { HomeComponent }        from './home/home.component';
import { RouterModule, Routes } from '@angular/router';
import { NgModule }             from '@angular/core';
import { AuthGuardService }     from './auth-guard.service';
import { CallbackComponent }    from './callback/callback.component';
import { PinnedReposComponent } from './pinned-repos/pinned-repos.component';

const routes: Routes = [
  { path: '', redirectTo: 'home', pathMatch: 'full' },
  { path: 'home', component: HomeComponent },
  { path: 'pinnedRepos/:organization', component: PinnedReposComponent,  canActivate: [AuthGuardService] },
  { path: 'login/callback', component: CallbackComponent }
];

@NgModule({
  imports: [ RouterModule.forRoot(routes) ],
  exports: [ RouterModule ]
})
export class AppRoutingModule { }
