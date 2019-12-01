import { Injectable }  from '@angular/core';
import { HttpInterceptor, HttpRequest, HttpHandler, HttpEvent } from '@angular/common/http';
import { Observable }  from 'rxjs/internal/Observable';
import { AuthService } from '@app/_services/auth.service';

@Injectable()
export class TokenInterceptor implements HttpInterceptor {
    constructor(public auth: AuthService) {}
    intercept(request: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
          request = request.clone({
                  setHeaders: {
                            Authorization: this.auth.createAuthHeaderValue()
                          }
                });
          return next.handle(request);
        }
}