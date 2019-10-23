import React, { Component } from "react";

import SidebarItem from "./SidebarItem";

export class Sidebar extends Component {
  // state
  state = {
    items: []
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
          <SidebarItem key={item.id} />
        </div>
      );
    });
  };

  render() {
    return <div>{this.renderSidebarItems()}</div>;
  }
}

export default Sidebar;
