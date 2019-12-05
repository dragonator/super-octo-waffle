import { RouterModule, Routes } from '@angular/router';
import { NgModule }             from '@angular/core';

import { AuthGuard }            from '@app/_helpers';
import { HomeComponent }        from '@app/home';
import { CallbackComponent }    from '@app/callback';
import { PinnedReposComponent } from '@app/pinned-repos';
import { RepoComponent }        from '@app/repo';

const routes: Routes = [
  { path: 'login/callback',            component: CallbackComponent },
  { path: 'home',                      component: HomeComponent },
  { path: 'pinnedRepos/:organization', component: PinnedReposComponent,  canActivate: [AuthGuard] },
  { path: 'repo/:organization/:repo',  component: RepoComponent,         canActivate: [AuthGuard] },

  { path: 'login', redirectTo: 'dev-a638a1un.auth0.com/login', pathMatch: 'full'},
  { path: '**',    redirectTo: 'home',                         pathMatch: 'full' },
];

@NgModule({
  imports: [ RouterModule.forRoot(routes) ],
  exports: [ RouterModule ]
})
export class AppRoutingModule { }
