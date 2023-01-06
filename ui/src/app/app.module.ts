import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { LoginComponent } from './pages/login/login.component';
import { RegisterComponent } from './pages/register/register.component';
import { AuthLayoutComponent } from './shared/layouts/auth-layout/auth-layout.component';
import { SiteLayoutComponent } from './shared/layouts/site-layout/site-layout.component';
import { DashboardComponent } from './pages/dashboard/dashboard.component';
import { AuthGuard } from './shared/classes/auth-guard';
import { UsersComponent } from './pages/users/users.component';
import { ProxiesComponent } from './pages/proxies/proxies.component';
import { JwtInterceptor } from './shared/classes/jwt-interceptor';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    RegisterComponent,
    AuthLayoutComponent,
    SiteLayoutComponent,
    DashboardComponent,
    UsersComponent,
    ProxiesComponent,
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    AppRoutingModule,
    FormsModule,
    ReactiveFormsModule,

    //FlexLayoutModule,
  ],
  providers: [
    AuthGuard,
    {
      provide: HTTP_INTERCEPTORS,
      multi: true,
      useClass: JwtInterceptor
    }
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
