import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { BrowserModule }           from '@angular/platform-browser';
import { ClarityModule }           from '@clr/angular';
import { FormsModule }             from '@angular/forms';
import { NgModule }                from '@angular/core';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';

import { AppRoutingModule }     from '@app/app.routing';
import { AppComponent }         from '@app/app.component';
import { CallbackComponent }    from '@app/callback';
import { HomeComponent }        from '@app/home';
import { PinnedReposComponent } from '@app/pinned-repos';
import { AuthGuard, JwtInterceptor }       from '@app/_helpers';
import { AuthService, PinnedReposService } from '@app/_services';

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    PinnedReposComponent,
    CallbackComponent
  ],
  imports: [
    AppRoutingModule,
    BrowserModule,
    ClarityModule,
    FormsModule,
    HttpClientModule
  ],
  providers: [AuthGuard, AuthService, PinnedReposService, {
    provide: HTTP_INTERCEPTORS,
    useClass: JwtInterceptor,
    multi: true
  }],
  bootstrap: [AppComponent]
})
export class AppModule { }
