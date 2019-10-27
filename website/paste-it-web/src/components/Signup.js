import React, { Component } from "react";
import axios from "axios";

export class Signup extends Component {

  // state for signup
  state = {
    email: "",
    user_name: "",
    pasword: ""
  }

  // login btn style
  btnLoginStyle = {
    marginLeft: "16px"
  };

  // handle textfield changes
  handleEmailChange = (event) => {
    this.setState({ email: event.target.value })
  }

  handleTfChange = (event) => {
    this.setState({ user_name: event.target.value })
  }

  handlePsdChange = (event) => {
    this.setState({ pasword: event.target.value })
  }

  // handle create account
  handleCreateAccount = (event) => {
    event.preventDefault()
    axios.post("http://localhost:5000/signup", this.state)
      .then((res) => {
        console.log(res)
        // put token in local storage
        window.localStorage.setItem("token", res.data.token)
        window.localStorage.setItem("userid", res.data.id)

        this.sendToDashboard()
      })
      .catch((err) => {
        console.log(err)
      })
  }

  // handle login instead

  sendToDashboard = () => {
    console.log("sending to dashboard")
    this.props.loginstatusNav(true)
  }

  render() {
    return (
      <div style={this.containerStyle}>
        <h1 className="title text-center">Paste it</h1>
        <div className="row">
          <p className="col-4" />
          <h3 className="text-left col">Welcome back :)</h3>
        </div>
        <div className="row">
          <p className="col-4" />
          <p className="col-4">
            <small>
              To keep connected with us, please login with your username and
              password.
            </small>
          </p>
        </div>
        <div className="row">
          <p className="col-4" />
          <form className="col-4">
            <div className="form-group">
              <label htmlFor="inputEmail">Email</label>
              <input
                type="text"
                className="form-control"
                id="inputEmail"
                aria-describedby="emailHelp"
                placeholder="Email"
                onChange={this.handleEmailChange}
              />
            </div>
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
              onClick={this.handleCreateAccount}
            >
              Create Account
            </button>
            <button
              type="submit"
              className="btn btn-outline-secondary"
              style={this.btnLoginStyle}
              onClick={this.handleLoginInstead}
            >
              Login instead
            </button>
          </form>
        </div>
      </div>
    );
  }
}

export default Signup;
