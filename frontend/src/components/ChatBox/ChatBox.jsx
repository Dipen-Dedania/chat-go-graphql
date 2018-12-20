import React, { Component } from "react";

export default class ChatBox extends Component {
  constructor() {
    super();
    this.state = {
      chatDisplay: "none"
    };

    this.toggleChat = this.toggleChat.bind(this);
  }
  toggleChat() {
    this.setState({
      chatDisplay: ""
    });
  }

  render() {
    return (
      <div className="chatbox">
        <button className="open-button" onClick={this.toggleChat}>
          Chat
        </button>

        <div
          className="chat-popup"
          id="myChat"
          style={{ display: this.state.chatDisplay }}
        >
          <form action="/action_page.php" className="form-container">
            <h1>Chat</h1>

            <label htmlFor="msg">
              <b>Message</b>
            </label>
            <textarea placeholder="Type message.." name="msg" required />

            <button type="submit" className="btn">
              Send
            </button>
          </form>
        </div>
      </div>
    );
  }
}
