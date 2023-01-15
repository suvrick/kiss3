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

  displayItems: Proxy[]
  page: number
  selectCount: number
  canPrev: boolean
  canNext: boolean

  constructor(private proxyService: ProxiesService) {
    this.page = 0
    this.selectCount = 5
    this.displayItems = []
   }

  ngOnInit(): void {
    this.proxies = [];
    this.proxyService.getProxies().subscribe({
      next: (resp) => {
        console.log(resp)
        if (resp.code == 200) {
          this.proxies.push(...resp.data)
          this.updateDisplay()
        }
      },
      error: (error) => {
        console.warn(error)
      }
    })
  }

  onFileSelected(fileInput: any) {
    const reader = new FileReader();
    reader.onload = (e: any) => {
      let text = e.target.result as string
      if(text){
        let lines = text.split("\n")
        lines.forEach((proxy)=>{
          let result = this.parseProxy(proxy)
          if (result) {
            this.createProxy(result)
          }
        }) 
      }
    }

    reader.readAsText(fileInput.target.files[0])
  }

  parseProxy(line: string): Proxy | undefined {
    let array = line.split(":")
    if(array.length == 4){
      let proxy = new Proxy()
      proxy.host = array[0]
      proxy.port = Number.parseInt(array[1])
      proxy.username = array[2]
      proxy.password = array[3]

      return proxy
    } else {
      return undefined
    }
  }

  createProxy(proxy: Proxy) {
    this.proxyService.addProxies(proxy).subscribe({
      next: (resp) => {
        console.log(resp)
        if (resp.code == 200) {
          this.proxies.push(resp.data)
          this.updateDisplay()
        }
      },
      error: (err) => {
        console.error(err)
      }
    })
  }

  deleteProxy() {
    this.proxies.forEach(p => {
      if (p.selected) {
        console.log(p)
        this.proxyService.deleteProxy(p).subscribe({
          next: (resp) => {
            console.log(resp)
            if (resp.code == 200) {
              this.proxies = this.proxies.filter(proxy => proxy.id != p.id)
              this.updateDisplay()
            }
          },
          error: (err) => {
            console.error(err)
          }
        })
      }
    });
  }

  toggleSelect(e: any) {
    this.proxies.forEach(p => {
      p.selected = e.target.checked
    })
  }

  updateDisplay(){
    this.displayItems = []
    let start = this.page * this.selectCount;
    let end = this.page * this.selectCount + this.selectCount;

    this.canPrev = this.page - 1 < 0;
    this.canNext = (this.page + 1)* this.selectCount >= this.proxies.length;

    if (start >= 0) {
      for (let index = start; index < end; index++) {
        if (index < this.proxies.length){
          const element = this.proxies[index];
          this.displayItems.push(element)
        }
      }
    }
  }

  pagePrev() {
    this.page = this.page - 1;
    this.updateDisplay();
  }
  
  pageNext() {
    this.page = this.page + 1;
    this.updateDisplay()
  }
}
