import { Component, OnInit } from '@angular/core';
import { User } from "./user";
import { UserService } from "./user.service";

@Component({
  selector: 'app-root',
  styleUrls: ['./app.component.css'],
  providers: [UserService],
  templateUrl: './app.component.html'
})
export class AppComponent implements OnInit {
  title = 'Declaration App';
  selectedUser: User;
  users: User[];

  constructor(private userService: UserService){
    
  }

  ngOnInit(): void {
    this.getUsers();
  }

  getUsers():void{
    this.userService.getUsers().then(users => this.users = users);
  }

  onSelect(user: User): void{
    this.selectedUser = user;
  }
}
