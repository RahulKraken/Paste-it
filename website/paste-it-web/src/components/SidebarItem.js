import React, { Component } from 'react'

export class SidebarItem extends Component {
  
  // state
  state = {
    key: this.props.id,
    selected: false
  }

  // modify selected state
  // @param "val" : bool
  setSelected = () => {
    console.log(this.state.key)
    this.setState({selected: true})
  }

  render() {
    return (
      <div onClick={this.setSelected}>
        <h3>{this.props.title}</h3>
      </div>
    )
  }
}

export default SidebarItem
