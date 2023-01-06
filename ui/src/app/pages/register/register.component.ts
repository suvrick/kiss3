import { Component, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/shared/services/auth.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss']
})
export class RegisterComponent implements OnInit {

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
          this.errorMsg = "Ошибка при создании аккаунта"
        },
        complete: () => {
          this.route.navigate(["login"])
        }
    })
  }
}
