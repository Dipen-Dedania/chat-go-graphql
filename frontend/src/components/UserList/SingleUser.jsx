import React, { Component } from "react";
import UserImg from "../../images/userimg.png";
export default class SingleUser extends Component {
  render() {
    return (
      <div class="item">
        <div class="right floated content">
          <div class="ui button">Add</div>
        </div>
        <img class="ui avatar image" src={UserImg} />>
        <div class="content">Lena</div>
      </div>
    );
  }
}
