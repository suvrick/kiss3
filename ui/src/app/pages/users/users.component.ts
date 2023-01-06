import { Component, OnInit } from '@angular/core';
import { User } from 'src/app/shared/classes/user';
import { UsersService } from 'src/app/shared/services/users.service';

@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.scss']
})
export class UsersComponent implements OnInit {

  users: User[]

  constructor(private usersService: UsersService) { }

  ngOnInit(): void {
    this.users = []
    this.usersService.getUsers().subscribe({
      next: (resp) => {
        console.log(resp)
        if (resp.code == 200) {
          this.users.push(...resp.data)
        }
      },
      error: (err) => {
        console.warn(err)
      }
    })
  }

}
