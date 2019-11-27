import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { BrowserModule } from '@angular/platform-browser';
import { ClarityModule } from '@clr/angular';
import { FormsModule }   from '@angular/forms';
import { NgModule }      from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';

import { AppComponent }         from './app.component';
import { CallbackComponent }    from './callback/callback.component';
import { HomeComponent }        from './home/home.component';
import { PinnedReposComponent } from './pinned-repos/pinned-repos.component';

import { AuthGuardService }   from './auth-guard.service';
import { AuthService }        from './auth.service';
import { PinnedReposService } from './pinned-repos.service';
import { TokenInterceptor }   from './token.interceptor';

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
  providers: [{
    provide: HTTP_INTERCEPTORS,
    useClass: TokenInterceptor,
    multi: true
  }],
  bootstrap: [AppComponent]
})
export class AppModule { }
