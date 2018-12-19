import React, { Component } from "react";

import SingleUser from "./SingleUser";

export default class UserList extends Component {
  render() {
    return (
      <div className="userlist">
        <div class="ui list">
          <SingleUser />
          <SingleUser />
          <SingleUser />
        </div>
      </div>
    );
  }
}
