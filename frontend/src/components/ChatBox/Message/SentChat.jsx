import React, { Component } from "react";
import SenderImage from "../../../images/userimg.png";

export default class SentChat extends Component {
  render() {
    return (
      <div className="sent-chat">
        <div className="ui comments">
          <div className="comment">
            <div className="content">
              <div className="sent-text">
                <div className="message-info">
                  <div className="metadata">
                    <span className="date">Today at 5:42PM</span>
                  </div>
                  <a href="/" className="author" style={{ color: "#006df0" }}>
                    {this.props.welcomeName}
                  </a>
                </div>
                <div className="message">
                  <div className="text">How dsadasd!</div>
                </div>
              </div>
              <div className="sent-image">
                <a href="/" className="avatar">
                  <img src={SenderImage} alt="Sender" />
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>
    );
  }
}
