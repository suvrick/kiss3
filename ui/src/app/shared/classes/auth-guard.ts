import { Injectable } from "@angular/core";
import { RouterStateSnapshot, ActivatedRouteSnapshot, CanActivate, CanActivateChild, Router} from "@angular/router";
import { Observable, of } from "rxjs";
import { AuthService } from "../services/auth.service";

@Injectable()
export class AuthGuard implements CanActivate, CanActivateChild {

    constructor(private authService: AuthService, private router: Router){}

    canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): Observable<boolean> {
        if (this.authService.isAuth()) {
            return of(true)
        } else {
            this.router.navigate(['login'])
            return of(false) 
        }
    }
    canActivateChild(childRoute: ActivatedRouteSnapshot, state: RouterStateSnapshot): Observable<boolean> {
        return this.canActivate(childRoute, state)
    }

}