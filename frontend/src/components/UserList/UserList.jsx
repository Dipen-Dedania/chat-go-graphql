import React, { Component } from "react";

import SingleUser from "./SingleUser";

export default class UserList extends Component {
  render() {
    return (
      <div class="ui middle aligned divided list">
        <SingleUser />
        <SingleUser />
        <SingleUser />
        <SingleUser />
      </div>
    );
  }
}
