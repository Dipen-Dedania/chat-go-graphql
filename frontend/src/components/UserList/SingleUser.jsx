import React, { Component } from "react";
import UserImg from "../../images/userimg.png";
export default class SingleUser extends Component {
  render() {
    return (
      <div>
        <div class="item">
          <img class="ui avatar image" src={UserImg} />

          <a class="header">Akhil</a>
        </div>
      </div>
    );
  }
}
