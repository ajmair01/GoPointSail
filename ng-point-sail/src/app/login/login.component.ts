import { Component } from '@angular/core';
import {FormsModule} from "@angular/forms";
import {User} from "../model/user";
import {NgIf} from "@angular/common";
import {UserService} from "../service/user.service";
import {Router} from "@angular/router";
import {log} from "@angular-devkit/build-angular/src/builders/ssr-dev-server";

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [
    FormsModule,
    NgIf
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

  submitSignup() {
    this.userService.saveUser(this.user).subscribe({
      next: value => {
        console.log(value);
        this.router.navigate(['users']).then(()=>{});
      },
      error: error => console.error(error),
      complete: () => console.info('complete')
    });
  }
}
