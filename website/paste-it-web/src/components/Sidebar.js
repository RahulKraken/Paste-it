import React, { Component } from "react";

import SidebarItem from "./SidebarItem";
import axios from "axios";

export class Sidebar extends Component {
  // state
  state = {
    items: this.props.items
  }

  // change display mode
  setSideBarVisible = val => {
    this.setState({ display: val });
  };

  // change items
  setItems = itemList => {
    this.setState({ items: itemList });
  };

  // render multiple sidebar items
  renderSidebarItems = () => {
    console.log("sidebar items", this.props.items)
    return this.props.items.map(item => {
      return (
        <div>
          <SidebarItem key={item.id} title={item.title} />
        </div>
      );
    });
  };

  render() {
    console.log("rendering sidebar items")
    return <div>{this.renderSidebarItems()}</div>;
  }
}

export default Sidebar;
