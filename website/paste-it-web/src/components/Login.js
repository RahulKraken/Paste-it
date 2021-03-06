import React, { Component } from "react";
import axios from "axios";
import "../css/login.css";

export class Login extends Component {
  // login data state
  state = {
    user_name: "",
    pasword: "",
    email: ""
  };

  // signup btn style
  btnSignupStyle = {
    marginLeft: "16px"
  };

  // handle changes to text
  handleTfChange = event => {
    this.setState({ user_name: event.target.value })
  };

  // handle password changes
  handlePsdChange = event => {
    this.setState({ pasword: event.target.value })
  };

  // login handler
  handleLogin = event => {
    event.preventDefault();
    console.log("Login now clicked!!!")
    console.log(this.state)
    axios
      .post(`http://localhost:5000/login`, {
        user_name: this.state.user_name,
        pasword: this.state.pasword
      })
      .then(res => {
        console.log(res.data)
        // put token in local storage
        window.localStorage.setItem("token", res.data.token)
        window.localStorage.setItem("userid", res.data.id)
        this.sendToDashboard();
      })
      .catch(err => {
        console.log(err)
        alert("make sure username and password are correct!!!")
      });
  };

  handleCreateAccount = (event) => {
    event.preventDefault()
    console.log("redirecting to signup page")
    this.props.signupstatusNav(true)
  }

  sendToDashboard() {
    this.props.loginstatusNav(true)
  }

  render() {
    return (
      <div style={this.containerStyle}>
        <h1 className="title text-center">Paste it</h1>
        <div className="row">
          <p className="col-4"></p>
          <h3 className="text-left col">Welcome back :)</h3>
        </div>
        <div className="row">
          <p className="col-4"></p>
          <p className="col-4">
            <small>
              To keep connected with us, please login with your username and
              password.
            </small>
          </p>
        </div>
        <div className="row">
          <p className="col-4"></p>
          <form className="col-4">
            <div className="form-group">
              <label htmlFor="inputUsername">Username</label>
              <input
                type="text"
                className="form-control"
                id="inputUsername"
                aria-describedby="emailHelp"
                placeholder="Username"
                onChange={this.handleTfChange}
              />
            </div>
            <div className="form-group">
              <label htmlFor="inputpasword">Password</label>
              <input
                type="password"
                className="form-control"
                id="inputpasword"
                placeholder="Password"
                onChange={this.handlePsdChange}
              />
            </div>
            <button
              type="submit"
              className="btn btn-primary"
              onClick={this.handleLogin}
            >
              Login now
            </button>
            <button
              type="submit"
              className="btn btn-outline-secondary"
              style={this.btnSignupStyle}
              onClick={this.handleCreateAccount}
            >
              Create Account
            </button>
          </form>
        </div>
      </div>
    );
  }
}

export default Login;
