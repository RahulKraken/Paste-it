import React, { Component } from "react";

export class Signup extends Component {
  render() {
    return (
      <div style={this.containerStyle}>
        <h1 className="title text-center">Paste it</h1>
        <div className="row">
          <p className="col-4"/>
          <h3 className="text-left col">Welcome back :)</h3>
        </div>
        <div className="row">
          <p className="col-4"/>
          <p className="col-4">
            <small>
              To keep connected with us, please login with your username and
              password.
            </small>
          </p>
        </div>
        <div className="row">
          <p className="col-4"/>
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

export default Signup;
