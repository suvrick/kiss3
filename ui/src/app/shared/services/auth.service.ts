import { HttpClient } from '@angular/common/http';
import { Binary } from '@angular/compiler';
import { Injectable } from '@angular/core';
import { Observable, tap } from 'rxjs';
import { Response } from "../classes/response"
import { User } from '../classes/user';

@Injectable({
    providedIn: 'root'
})
export class AuthService {

    token: string | null = null;

    constructor(private http: HttpClient) {}

    login(email: string, password: string): Observable<Response<{token: string}>> {
        return this.http.post<Response<{token: string}>>('/api/auth/login', {
            email: email,
            password: password
        }).pipe(
            tap(
                resp => this.setToken(resp.data.token)
            )
        )
    }

    register(email: string, password: string): Observable<{}> {
        return this.http.post<Response<{}>>('/api/auth/register', {
            email: email,
            password: password
        })
    }

    logout() {
        this.token = ""
        localStorage.removeItem("token")
    }

    isAuth(): boolean {
        return !!this.getToken()
    }
    
    setToken(token: string) {
        this.token = token
        this.setSession()
    }

    getToken(): string | null {
        return this.token
    }

    setSession() {
        localStorage.setItem("token", this.token as string)
    }

    restoreSession() {
        this.token = localStorage.getItem("token")
    }

    getUser(): User | undefined {
        if (!!this.token && this.token.split(".").length == 3) {
            let substr = this.token.split(".")[1];
            return JSON.parse(atob(substr)) as User
        }
        return undefined
    }
}
