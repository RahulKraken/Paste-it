import React, { Component } from "react";

import Workspace from "./Workspace";
import Sidebar from "./Sidebar";
import axios from "axios";

export class Dashboard extends Component {
  // state
  state = {
    items: [],
    selected: -1
  };

  setSelected = (idx) => {
    this.setState({ selected: idx })
  }

  /**
   * just test function to fetch list of 
   * items for the given user
   */
  fetchItemList = () => {
    console.log("fetching item list")
    axios
      .get(`http://localhost:5000/api/pastes/` + window.localStorage.getItem("userid"), {
        headers: {
          Token: window.localStorage.getItem("Token")
        }
      })
      .then(res => {
        // console.log("successfully logged in")
        console.log(res.data)
        this.setState({ items: res.data })
        console.log("state:", this.state.items)
      })
      .catch(err => {
        console.log(err)
        // this.logout()
      });
  };

  /**
   * performs logout operations
   * make jwt token = ""
   * make userid = ""
   * navigate to login screen
   */
  logout = () => {
    console.log("Logging out")
    window.localStorage.setItem("token", "")
    window.localStorage.setItem("userid", "")
    this.props.loginstatusNav(false)
    console.log("logged out")
  }

  componentDidMount() {
    // fetch fetchItemList
    if (window.localStorage.getItem("token") === "") {
      this.logout()
    }
    this.fetchItemList()
  }

  render() {
    return (
      <div className="row">
        <div className="col-4">
          <Sidebar items={this.state.items} setSelected={this.setSelected}/>
        </div>
        <div className="col-8">
          <Workspace selected={this.state.selected}/>
        </div>
      </div>
    );
  }
}

export default Dashboard;
