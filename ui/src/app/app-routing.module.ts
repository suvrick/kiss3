import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { DashboardComponent } from './pages/dashboard/dashboard.component';
import { LoginComponent } from './pages/login/login.component';
import { ProxiesComponent } from './pages/proxies/proxies.component';
import { RegisterComponent } from './pages/register/register.component';
import { UsersComponent } from './pages/users/users.component';
import { AuthGuard } from './shared/classes/auth-guard';
import { AuthLayoutComponent } from './shared/layouts/auth-layout/auth-layout.component';
import { SiteLayoutComponent } from './shared/layouts/site-layout/site-layout.component';
import { AuthService } from './shared/services/auth.service';
import { UsersService } from './shared/services/users.service';

const routes: Routes = [
  {
    path: '', component: SiteLayoutComponent, canActivate: [AuthGuard], children: [
      { path: '', redirectTo: 'dashboard', pathMatch: 'full' },
      { path: "dashboard", component: DashboardComponent },
      { path: "proxies", component: ProxiesComponent },
      { path: "users", component: UsersComponent },
    ]
  },
  {
    path: '', component: AuthLayoutComponent, children: [
      { path: '', redirectTo: 'dashboard', pathMatch: 'full' },
      { path: 'login', component: LoginComponent },
      { path: 'register', component: RegisterComponent },
    ]
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
  constructor(private authService: AuthService){}
}
