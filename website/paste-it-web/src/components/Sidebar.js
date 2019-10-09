import React, { Component } from "react";

import SidebarItem from "./SidebarItem";

export class Sidebar extends Component {
  // state
  state = {
    items: [1, 2, 3, 4, 5, 6, 7, 8],
    display: true
  };

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
    return this.state.items.map(item => {
      return (
        <div>
          <SidebarItem />
        </div>
      );
    });
  };

  render() {
    return <div>{this.renderSidebarItems()}</div>;
  }
}

export default Sidebar;
