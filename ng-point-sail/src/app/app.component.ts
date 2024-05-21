import { Component } from '@angular/core';
import {ActivatedRoute, Router, RouterLink, RouterLinkActive, RouterOutlet} from '@angular/router';
import {NgbNav, NgbNavItem, NgbNavLinkBase} from "@ng-bootstrap/ng-bootstrap";
import {User} from "./model/user";
import {UserService} from "./service/user.service";
import {NgClass, NgIf} from "@angular/common";

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, RouterLink, RouterLinkActive, NgbNav, NgbNavItem, NgbNavLinkBase, NgClass, NgIf],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent {
  loggedInUser: User = {}

  constructor(public router: Router,
              private userService: UserService) {
    userService.loggedInUser$.subscribe(
      user => {
        this.loggedInUser = user;
      }
    )
  }

  logout() {
    this.userService.logout();
    this.router.navigate(['/login']).then(()=>{});
  }
}
