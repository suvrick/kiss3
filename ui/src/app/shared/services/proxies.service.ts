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
    return this.http.get<Response<Proxy[]>>('/api/proxies/')
  }

  addProxies(proxy: Proxy): Observable<Response<Proxy>> {
    return this.http.post<Response<Proxy>>('/api/proxies/', JSON.stringify(proxy))
  }

  deleteProxy(proxy: Proxy): Observable<Response<{}>> {
    return this.http.delete<Response<Proxy>>('/api/proxies/' + proxy.id)
  }
}
