import React, { Component } from 'react'

export class SidebarItem extends Component {
  
  // state
  state = {
    key: -1,
    selected: false
  }

  // modify selected state
  // @param "val" : bool
  setSelected = (val) => {
    this.setState({selected: val})
  }

  render() {
    return (
      <div>
        <h3>SidebarItem</h3>
      </div>
    )
  }
}

export default SidebarItem
