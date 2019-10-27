import React, { Component } from "react";

import SidebarItem from "./SidebarItem";

export class Sidebar extends Component {
  // state
  state = {
    items: this.props.items
  }

  // render multiple sidebar items
  renderSidebarItems = () => {
    return this.props.items.map(item => {
      return (
        <div>
          <SidebarItem key={item.id} id={item.id} title={item.title} />
        </div>
      );
    });
  };

  render() {
    return <div>{this.renderSidebarItems()}</div>;
  }
}

export default Sidebar;
