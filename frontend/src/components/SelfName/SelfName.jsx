import React, { Component } from "react";

export default class SelfName extends Component {
  render() {
    return (
      <div class="enter-name">
        <input class="cbox" type="checkbox" />
        <label class="add" for="cbox">
          Enter Name
        </label>
        <input class="message" type="text" />
      </div>
    );
  }
}
