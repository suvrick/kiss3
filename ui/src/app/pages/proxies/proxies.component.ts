import { Component, OnInit } from '@angular/core';
import { Proxy } from 'src/app/shared/classes/proxy';
import { ProxiesService } from 'src/app/shared/services/proxies.service';

@Component({
  selector: 'app-proxies',
  templateUrl: './proxies.component.html',
  styleUrls: ['./proxies.component.scss']
})
export class ProxiesComponent implements OnInit {

  proxies: Proxy[]

  constructor(private proxyService: ProxiesService) { }

  ngOnInit(): void {
    this.proxies = [];
    this.proxyService.getProxies().subscribe({
      next: (resp) => {
        if (resp.code == 200) {
          this.proxies.push(...resp.data)
        }
      },
      error: (error) => {
        console.warn(error)
      }
    })
  }

}
