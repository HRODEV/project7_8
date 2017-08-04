import * as React from 'react';
import './App.css';
import { User } from './API_models';
import UserService from './UserAPI.service'

const logo = require('./logo.svg');

type AppState = { users: User[] };
class App extends React.Component<{}, AppState> {
  constructor(props: {}, context: any) {
    super(props, context);
    this.state = {users: []};    
  }

  componentWillMount(): void {
    UserService.GetAllUsers()
        .then(j => this.setState({...this.state, users: j as User[] }))
        .catch(e => console.error("there is an network error"));
  }
  
  render() {
    
    return (
      <div className="App">
        <div className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h2>Declaration App</h2>
        </div>
        <ol>
          {this.state.users.map(user => <li>ID:{user.ID}   Email:{user.Email} </li>)}
        </ol>
      </div>
    );
  }
}

export default App;
