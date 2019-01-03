import React, { Component } from "react";
import CloseButton from "../../images/close-button.png";
import SendButton from "../../images/send.png";

import gql from "graphql-tag"; // GraphQL
import { graphql, compose } from "react-apollo";

import SentChat from "./Message/SentChat"; //Component
import ReceivedChat from "./Message/ReceivedChat";

class ChatBox extends Component {
  constructor(props) {
    super(props);
    this.state = {
      chatDisplay: this.props.openChatBox,
      welcomeName: "Akhil",
      friendName: "Aneri",
      message: ""
    };

    this.closeChat = this.closeChat.bind(this);
    this.messageSent = this.messageSent.bind(this);
  }

  componentDidMount() {
    // RISKY CODE
    this.createMessageSubscription = this.props.data.subscribeToMore({
      document: Subscription,
      updateQuery: (previousState, { subscriptionData }) => {
        console.log("updateQuery");
        let newMessage = subscriptionData.data.messagePosted;

        let messages = Object.assign({}, previousState, {
          chats: [...previousState.chats, newMessage]
        });
        console.log(this.props);
        return messages;
      },
      onError: err => console.error(err)
    });
    // RISKY CODE ABOVE
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

  messageSent(e) {
    if (e.key === "Enter" && !(e.target.value === "")) {
      // console.log("this.state.welcomeName", this.state.welcomeName);
      // console.log("this.state.friendName", this.state.friendName);
      this.props.mutate({
        variables: {
          sender_name: this.state.welcomeName,
          receiver_name: this.state.friendName,
          message: e.target.value
        }
        // refetchQueries: [{ query }]
      });

      e.target.value = "";
    }
  }

  renderChat() {
    const data = this.props.data;
    if (data.loading) {
      return <div className="ui active centered inline loader" />;
    }

    if (data.error) {
      return <h3>Error : Server is Down :( </h3>;
    }

    return data.chats.map((chat, i) => {
      // return (<div key={i}>{chat.message}</div>);
      if (chat.sender_name === this.props.welcomeName) {
        return (
          <SentChat
            key={i}
            welcomeName={this.state.welcomeName}
            inputMessage={chat.message}
          />
        );
      } else {
        return (
          <ReceivedChat
            key={i}
            receiverName={this.state.friendName}
            receivedMessage={chat.message}
          />
        );
      }
    });
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
          <div className="chatbox-message-area">{this.renderChat()} </div>
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

const query = gql`
  query chat($sender_name: String!, $receiver_name: String!) {
    chats(sender_name: $sender_name, receiver_name: $receiver_name) {
      sender_name
      receiver_name
      message
      createdAt
    }
  }
`;

const mutation = gql`
  mutation addMessage(
    $sender_name: String!
    $receiver_name: String!
    $message: String!
  ) {
    postMessage(
      sender_name: $sender_name
      receiver_name: $receiver_name
      message: $message
    ) {
      sender_name
      receiver_name
      message
    }
  }
`;

const Subscription = gql`
  subscription realtimechat {
    messagePosted(id: "bgjubgj") {
      sender_name
      receiver_name
      message
      createdAt
    }
  }
`;

export default compose(
  graphql(mutation),
  graphql(query, {
    options: props => {
      return {
        variables: {
          sender_name: props.welcomeName,
          receiver_name: props.friendName
        }
      };
    }
  })
)(ChatBox);
