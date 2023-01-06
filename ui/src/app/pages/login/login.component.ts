import { Component, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';
import { Route, Router } from '@angular/router';
import { concatWith } from 'rxjs';
import { AuthService } from '../../shared/services/auth.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

  errorMsg: string

  constructor(private authService: AuthService, private route: Router) { }

  ngOnInit(): void {
  }

  onSubmit(form: NgForm) {
    this.authService.login(form.value.email, form.value.password)
      .subscribe({
        next: (res) => {  
          //
        },
        error: err => {
          this.errorMsg = "bad request"
        },
        complete: () => {
          this.route.navigate(["dashboard"])
        }
    })
  }
}