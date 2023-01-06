import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Response } from '../classes/response';
import { User } from '../classes/user';

@Injectable({
  providedIn: 'root'
})
export class UsersService {

  constructor(private http: HttpClient) { }

  getUsers(): Observable<Response<User[]>> {
    return this.http.get<Response<User[]>>('/api/users')
  }
}
