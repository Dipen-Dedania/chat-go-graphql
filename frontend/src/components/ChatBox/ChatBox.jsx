import React, { Component } from "react";
import CloseButton from "../../images/close-button.png";
import SendButton from "../../images/send.png";

export default class ChatBox extends Component {
  constructor(props) {
    super(props);
    this.state = {
      chatDisplay: false
    };

    this.showChat = this.showChat.bind(this);
    this.closeChat = this.closeChat.bind(this);
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

  // static getDerivedStateFromProps(props, state) {
  //   if (state.chatDisplay !== props.chatDisplay) {
  //     return {
  //       derivedValue: props.chatDisplay,
  //       mirroredProp: props.chatDisplay
  //     };
  //   }
  //   // when null is returned no update is made to the state
  //   return null;
  // }

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
            <span className="chat-name">{this.props.name}</span>
            <span className="chat-close" onClick={this.closeChat}>
              <img src={CloseButton} alt="" />
            </span>
          </div>
          <div className="chatbox-message-area">Heloooooooooooooooooo</div>
          <div className="chatbox-footer">
            <input
              type="text"
              className="message-input"
              placeholder="Enter to Send"
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
