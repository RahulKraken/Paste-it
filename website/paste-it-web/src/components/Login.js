import React, { Component } from 'react'
import axios from 'axios'
import '../css/login.css'

// images
// import tick from '../img/checked.png'

export class Login extends Component {

  // login data state
  state = {
    username: '',
    pasword: ''
  }

  // signup btn style
  btnSignupStyle = {
    marginLeft: '16px'
  }

  // handle changes to text
  handleTfChange = event => {
    this.setState({ username: event.target.value })
  }

  // handle password changes
  handlePsdChange = event => {
    this.setState({ pasword: event.target.value })
  }

  // login handler
  handleLogin = event => {
    event.preventDefault()
    console.log("Login now clicked!!!")
    console.log(this.state)
    axios.post(`http://localhost:5000/login`, {
      username: this.state.username,
      pasword: this.state.pasword
    })
      .then(res => {
        console.log(res.data)
        // put token in local storage
        window.localStorage.setItem("Token", res.data.token)
      })
      .catch(err => {
        console.log(err)
      })
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
              <label htmlFor="inputUsername">Username</label>
              <input type="text" className="form-control" id="inputUsername" aria-describedby="emailHelp" placeholder="Username"
                onChange={this.handleTfChange}/>
            </div>
            <div className="form-group">
              <label htmlFor="inputpasword">Password</label>
              <input type="password" className="form-control" id="inputpasword" placeholder="Password" onChange={this.handlePsdChange}/>
            </div>
            <button type="submit" className="btn btn-primary" onClick={this.handleLogin}>Login now</button>
            <button type="submit" className="btn btn-outline-secondary" style={this.btnSignupStyle} onClick={this.handleCreateAccount}>Create Account</button>
          </form>
        </div>
      </div>
    )
  }
}

export default Login
