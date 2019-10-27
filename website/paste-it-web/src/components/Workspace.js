import React, { Component } from 'react'

export class Workspace extends Component {
  render() {
    return (
      <div>
        <h1>{this.props.selected}</h1>
      </div>
    )
  }
}

export default Workspace
