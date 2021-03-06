import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { BrowserModule }           from '@angular/platform-browser';
import { ClarityModule }           from '@clr/angular';
import { FormsModule }             from '@angular/forms';
import { NgModule }                from '@angular/core';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';

import { AppRoutingModule }          from '@app/app.routing';
import { AppComponent }              from '@app/app.component';
import { CallbackComponent }         from '@app/callback';
import { HomeComponent }             from '@app/home';
import { PinnedReposComponent }      from '@app/pinned-repos';
import { RepoComponent }             from '@app/repo';
import { AuthGuard, JwtInterceptor } from '@app/_helpers';
import { AuthService, API }          from '@app/_services';

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    PinnedReposComponent,
    CallbackComponent,
    RepoComponent
  ],
  imports: [
    AppRoutingModule,
    BrowserModule,
    ClarityModule,
    FormsModule,
    HttpClientModule
  ],
  providers: [AuthGuard, AuthService, API, {
    provide: HTTP_INTERCEPTORS,
    useClass: JwtInterceptor,
    multi: true
  }],
  bootstrap: [AppComponent]
})
export class AppModule { }
