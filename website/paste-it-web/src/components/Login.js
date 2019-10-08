import React, { Component } from 'react'
import '../css/login.css'

// images
// import tick from '../img/checked.png'

export class Login extends Component {

  // signup btn style
  btnSignupStyle = {
    marginLeft: '16px'
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
          <p className="col-4"><small>To keep connected with us, please login with your username and password.</small></p>
        </div>
        <div className="row">
          <p className="col-4"></p>
          <form className="col-4">
            <div className="form-group">
              <label htmlFor="exampleInputEmail1">Username</label>
              <input type="text" className="form-control" id="exampleInputEmail1" aria-describedby="emailHelp" placeholder="Username"/>
            </div>
            <div className="form-group">
              <label htmlFor="exampleInputPassword1">Password</label>
              <input type="password" className="form-control" id="exampleInputPassword1" placeholder="Password"/>
            </div>
            <button type="submit" className="btn btn-primary" onClick={
              // login handler
              (event) => {
                event.preventDefault()
                console.log("Login now clicked!!!")
              }
            }>Login now</button>
            <button type="submit" className="btn btn-outline-secondary" style={this.btnSignupStyle} onClick={
              // create account handler
              (event) => {
                event.preventDefault()
                console.log("Create accont button clicked!!!")
              }
            }>Create Account</button>
          </form>
        </div>
      </div>
    )
  }
}

export default Login
