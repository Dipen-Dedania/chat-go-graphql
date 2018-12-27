import React, { Component } from "react";
import CloseButton from "../../images/close-button.png";
import SendButton from "../../images/send.png";
import SentChat from "./Message/SentChat";
import ReceivedChat from "./Message/ReceivedChat";

export default class ChatBox extends Component {
  constructor() {
    super();
    this.state = {
      chatDisplay: false,
      chatContent: ""
    };

    this.showChat = this.showChat.bind(this);
    this.closeChat = this.closeChat.bind(this);
    this.messageSent = this.messageSent.bind(this);
  }

  showChat() {
    this.setState({
      chatDisplay: true
    });
  }
  closeChat() {
    this.setState({
      chatDisplay: false
    });
  }

  messageSent(e) {
    if (e.key === "Enter" && !(e.target.value === "")) {
      this.setState({
        chatContent: (
          <div>
            <ReceivedChat receiverName={this.props.name} />
            <SentChat welcomeName={this.props.welcomeName} />
            <ReceivedChat receiverName={this.props.name} />
            <SentChat welcomeName={this.props.welcomeName} />
            <ReceivedChat receiverName={this.props.name} />
            <SentChat welcomeName={this.props.welcomeName} />
            <ReceivedChat receiverName={this.props.name} />
            <SentChat welcomeName={this.props.welcomeName} />
            <ReceivedChat receiverName={this.props.name} />
            <SentChat welcomeName={this.props.welcomeName} />
          </div>
        )
      });
      e.target.value = "";
    }
  }

  render() {
    return (
      <div className="chatbox">
        {/* Open Button */}
        <button className="open-button" onClick={this.showChat}>
          Chat
        </button>

        {/* REAL CHATBOX */}
        <div
          className="chat-popup"
          id="myChat"
          style={{ display: this.state.chatDisplay === false ? "none" : "" }}
        >
          <div className="chatbox-header">
            <span className="chat-name" style={{ overflow: "hidden" }}>
              {this.props.name}
            </span>
            <span className="chat-close" onClick={this.closeChat}>
              <img src={CloseButton} alt="" />
            </span>
          </div>
          <div className="chatbox-message-area">{this.state.chatContent}</div>
          <div className="chatbox-footer">
            <input
              type="text"
              className="message-input"
              placeholder="Enter to Send"
              onKeyPress={this.messageSent}
            />
            <button className="send-msg">
              <img src={SendButton} alt="Send" />
            </button>
          </div>
        </div>
      </div>
    );
  }
}
