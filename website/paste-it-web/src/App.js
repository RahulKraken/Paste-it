import React from 'react';
import './App.css';

// components
import Login from './components/Login'
import Dashboard from './components/Dashboard'
import Signup from './components/Signup'

class App extends React.Component {

  // login state
  state = {
    loginStatus: true,
    signupStatus: false
  }

  // method to modify state
  isLoggedIn = (val) => {
    this.setState({ loginStatus: val })
  }

  // method to modify signup state
  wantsSignup = (val) => {
    this.setState({ signupStatus: val })
  }

  render() {
    if (this.state.signupStatus) {
      return (
        <div className="App">
          <Signup loginstatusNav={ this.isLoggedIn } signupstatusNav={ this.wantsSignup }/>
        </div>
      )
    } else if (!this.state.loginStatus) {
      return (
        <div className="App">
          <Login loginstatusNav={ this.isLoggedIn } signupstatusNav={ this.wantsSignup }/>
        </div>
      );
    } else {
      return (
        <div className="App">
          <Dashboard loginstatusNav={ this.isLoggedIn } signupstatusNav={ this.wantsSignup }/>
        </div>
      )
    }
  }
}

export default App;
