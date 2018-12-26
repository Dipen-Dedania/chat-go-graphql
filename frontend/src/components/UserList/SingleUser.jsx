// React Import
import React, { Component } from "react";

// Profile Image Import
import UserImg from "../../images/userimg.png";

export default class SingleUser extends Component {
  render() {
    return (
      <div className="item">
        <div className="right floated content">
          <div
            className="ui button"
            onClick={() => this.props.callBackProp(this.props.username)}
          >
            Chat
          </div>
        </div>
        <img className="ui avatar image" src={UserImg} alt="Profile Pic" />
        <div className="content">{this.props.username}</div>
      </div>
    );
  }
}
