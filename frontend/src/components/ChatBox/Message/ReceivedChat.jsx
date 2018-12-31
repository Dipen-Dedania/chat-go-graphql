import React, { Component } from "react";
import SenderImage from "../../../images/userimg.png";

export default class ReceivedChat extends Component {
  render() {
    return (
      <div className="received-chat ">
        <div className="ui comments">
          <div className="comment">
            <a href="/" className="avatar">
              <img src={SenderImage} alt="Sender" />
            </a>
            <div className="content">
              <a href="/" className="author" style={{ color: "#009427" }}>
                {this.props.receiverName}
              </a>
              <div className="metadata">
                <span className="date">Today at 5:42PM</span>
              </div>
              <div className="text">{this.props.receivedMessage}</div>
            </div>
          </div>
        </div>
      </div>
    );
  }
}
