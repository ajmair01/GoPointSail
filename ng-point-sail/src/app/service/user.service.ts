import {Injectable} from "@angular/core";
import {HttpClient} from "@angular/common/http";
import {map, Observable, scan, Subject} from "rxjs";
import {User} from "../model/user";

@Injectable()
export class UserService {
  constructor(private http: HttpClient) {
  }

  private loggedInUserSource = new Subject<User>();
  loggedInUser$ = this.loggedInUserSource.asObservable();

  userUrl: string = "http://localhost:8080/users"
  loginUrl: string = "http://localhost:8080/login"

  getUsers(): Observable<User[]> {
    return this.http.get<User[]>(this.userUrl);
  }

  saveUser(user: User): Observable<User> {
    return this.http.post<User>(this.userUrl, user)
      .pipe(map(user => {
        this.loggedInUserSource.next(user);
        return user;
      }));
  }

  login(user: User): Observable<User> {
    return this.http.post<User>(this.loginUrl, user)
      .pipe(map(user => {
        this.loggedInUserSource.next(user);
        return user;
      }));
  }

  logout(): void {
    this.loggedInUserSource.next({});
  }
}
