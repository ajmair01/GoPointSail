import {Injectable} from "@angular/core";
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {User} from "../model/user";

@Injectable()
export class UserService {
  constructor(private http: HttpClient) {
  }

  userUrl: string = "http://localhost:8080/users"

  getUsers(): Observable<any> {
    return this.http.get<User[]>(this.userUrl);
  }

  saveUser(user: User): Observable<User> {
    return this.http.post<User>(this.userUrl, user);
  }
}
