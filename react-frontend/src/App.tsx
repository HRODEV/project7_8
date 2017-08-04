import * as React from 'react';
import './App.css';
import { User } from './API_models';
import UserService from './UserAPI.service'

const logo = require('./logo.svg');


type UserListItemComponentProps = {user:User, onClick?:((e?:React.MouseEvent<HTMLLIElement>) => void)}
class UserListItemComponent extends React.Component<UserListItemComponentProps,{}> {

    render(): JSX.Element {
      return (
      <li onClick={this.props.onClick} >
        <span className="UserId">{this.props.user.ID}</span> {this.props.user.Email}
      </li>);
    }
}

type AppState = { users: User[], selectedUser: User | undefined };
class App extends React.Component<{}, AppState> {
  constructor(props: {}, context: any) {
    super(props, context);
    this.state = {users: [], selectedUser: undefined};    
  }

  componentWillMount(): void {
    UserService.GetAllUsers()
        .then(j => this.setState({...this.state, users: j as User[] }))
        .catch(e => console.error("there is an network error"));
  }
  
  render(): JSX.Element {
    
    return (
      <div className="App">
        {this.state.selectedUser != undefined ? <div>user info</div> : undefined }
        <div className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h2>Declaration App</h2>
        </div>
        <ol>          
          {this.state.users.map(user => <UserListItemComponent onClick={e => this.setState({...this.state, selectedUser:user})} user={user} />)}
        </ol>
      </div>
    );
  }
}

export default App;
