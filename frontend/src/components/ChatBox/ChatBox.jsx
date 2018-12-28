import React, { Component } from "react";
import CloseButton from "../../images/close-button.png";
import SendButton from "../../images/send.png";

import gql from "graphql-tag"; // GraphQL
import { graphql, compose } from "react-apollo";

import SentChat from "./Message/SentChat"; //Component
import ReceivedChat from "./Message/ReceivedChat";

const query = gql`
  query chat {
    chats(sender_name: "Aneri", receiver_name: "AAyush") {
      sender_name
      receiver_name
      message
      createdAt
    }
  }
`;

const mutation = gql`
  mutation addMessage {
    postMessage(sender_name: "", receiver_name: "", message: "") {
      sender_name
      receiver_name
      message
    }
  }
`;

class ChatBox extends Component {
  constructor(props) {
    super(props);
    this.state = {
      chatDisplay: this.props.openChatBox,
      welcomeName: "",
      friendName: ""
    };

    this.closeChat = this.closeChat.bind(this);
    this.messageSent = this.messageSent.bind(this);
  }

  static getDerivedStateFromProps(nextProps, prevState) {
    return {
      chatDisplay: nextProps.openChatBox,
      welcomeName: nextProps.welcomeName,
      friendName: nextProps.friendName
    };
  }

  closeChat() {
    this.setState({
      chatDisplay: !this.state.openChatBox
    });
  }

  chatHistory() {
    return <div>Hello</div>;
  }

  messageSent(e) {
    //   if (e.key === "Enter" && !(e.target.value === "")) {
    //     const messageSent = (
    //       <div>
    //         <SentChat
    //           welcomeName={this.props.welcomeName}
    //           inputMessage={e.target.value}
    //         />
    //       </div>
    //     );
    //     this.setState({
    //       sentMessage: [...this.state.sentMessage, messageSent]
    //     });
    //     e.target.value = "";
    //   }
  }

  render() {
    return (
      <div className="chatbox">
        {/* REAL CHATBOX */}
        <div
          className="chat-popup"
          id="myChat"
          style={{ display: this.state.chatDisplay === false ? "none" : "" }}
        >
          <div className="chatbox-header">
            <span className="chat-name" style={{ overflow: "hidden" }}>
              {this.state.friendName}
            </span>
            <span className="chat-close" onClick={this.props.closeChatBox}>
              <img src={CloseButton} alt="" />
            </span>
          </div>
          <div className="chatbox-message-area">{this.chatHistory()} </div>
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

export default compose(
  graphql(mutation),
  graphql(query)
)(ChatBox);
