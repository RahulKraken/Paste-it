import React from 'react';
import './App.css';

// components
import Login from './components/Login'
import Dashboard from './components/Dashboard'

class App extends React.Component {

  // login state
  state = {
    loginStatus: true
  }

  // method to modify state
  isLoggedIn = (val) => {
    this.setState({ loginStatus: val })
  }

  render() {
    if (!this.state.loginStatus) {
      return (
        <div className="App">
          <Login navigator={ this.isLoggedIn }/>
        </div>
      );
    } else {
      return (
        <div className="App">
          <Dashboard navigator={ this.isLoggedIn }/>
        </div>
      )
    }
  }
}

export default App;
