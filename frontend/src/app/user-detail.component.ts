import { Component, Input } from '@angular/core';
import { User } from "./user";

@Component({
    selector: 'user-detail',
    template: `
    <div *ngIf="user">
      <h2>{{user.Email}} details!</h2>
      <div><label>id: </label>{{user.ID}}</div>
      <div>
        <label>name: </label>
        <input [(ngModel)]="user.Email" placeholder="name"/>
      </div>
    </div>
    `
})
export class UserDetailComponent {
    @Input() user: User;  
}
