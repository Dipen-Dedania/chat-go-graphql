import React, { Component } from "react";

export default class SelfName extends Component {
  render() {
    return (
      <div className="enter-name">
        <input className="cbox" type="checkbox" />
        <label className="add" htmlFor="cbox">
          Enter Your Name
        </label>
        <input className="message" type="text" />
      </div>
    );
  }
}
