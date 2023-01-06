import { Component, OnInit } from '@angular/core';
import { User } from 'src/app/shared/classes/user';
import { AuthService } from 'src/app/shared/services/auth.service';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {

  user: User | undefined

  constructor(private authService: AuthService) { }

  ngOnInit(): void {
    this.user = this.authService.getUser();
    console.log(this.user)
  }

}
