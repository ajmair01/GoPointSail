import {Component, EventEmitter, Output} from '@angular/core';
import {FormsModule} from "@angular/forms";
import {User} from "../model/user";
import {NgIf} from "@angular/common";
import {UserService} from "../service/user.service";
import {Router} from "@angular/router";
import {NgxMaskDirective} from "ngx-mask";

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [
    FormsModule,
    NgIf,
    NgxMaskDirective
  ],
  templateUrl: './login.component.html',
  styleUrl: './login.component.scss'
})
export class LoginComponent {
  user: User = {};
  signup: boolean = false;

  constructor(private userService: UserService,
              private router: Router) {
  }

  login() {
    this.userService.login(this.user).subscribe({
      next: value => {
        this.router.navigate(['users']).then(()=>{});
      },
      error: error => console.error(error),
      complete: () => console.info('complete login')
    })
  }

  submitSignup() {
    this.userService.saveUser(this.user).subscribe({
      next: value => {
        console.log(value);
        this.router.navigate(['users']).then(()=>{});
      },
      error: error => console.error(error),
      complete: () => console.info('complete signup')
    });
  }
}
