import { Routes } from '@angular/router';
import {LoginComponent} from "./login/login.component";
import {UserComponent} from "./user/user.component";
import {GroupComponent} from "./group/group.component";

export const routes: Routes = [
  { path: 'login', component: LoginComponent},
  { path: 'users', component: UserComponent},
  { path: 'groups', component: GroupComponent},
  { path: '', redirectTo: '/login', pathMatch: 'full'}
];
