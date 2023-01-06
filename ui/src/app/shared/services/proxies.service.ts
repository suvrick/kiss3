import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Proxy } from '../classes/proxy';
import { Response } from '../classes/response';

@Injectable({
  providedIn: 'root'
})
export class ProxiesService {

  constructor(private http: HttpClient) { }

  getProxies(): Observable<Response<Proxy[]>> {
    return this.http.get<Response<Proxy[]>>('/api/proxies')
  }
}
