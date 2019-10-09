import React, { Component } from "react";

import Workspace from "./Workspace";
import Sidebar from "./Sidebar";
import axios from "axios";

export class Dashboard extends Component {
  // state
  state = {
    items: []
  };

  /**
   * just test function to fetch list of 
   * items for the given user
   */
  fetchItemList = () => {
    axios
      .get(`http://localhost:5000/api/paste/1`, {
        headers: {
          Token: window.localStorage.getItem("Token")
        }
      })
      .then(res => {
        console.log(res.data);
      })
      .catch(err => {
        console.log(err)
      });
  };

  componentDidMount() {
    // fetch fetchItemList
    this.fetchItemList()
  }

  render() {
    return (
      <div className="row">
        <div className="col-4">
          <Sidebar />
        </div>
        <div className="col-8">
          <Workspace />
        </div>
      </div>
    );
  }
}

export default Dashboard;