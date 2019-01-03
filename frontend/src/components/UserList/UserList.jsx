import React, { Component } from "react";

// GraphQL
import gql from "graphql-tag";
import { graphql } from "react-apollo";

//Components Import
import ChatBox from "../ChatBox/ChatBox";
import SingleUser from "./SingleUser";

// Query
const query = gql`
  {
    users {
      name
    }
  }
`;

class UserList extends Component {
  constructor() {
    super();
    this.state = {
      userListDisplay: "none",
      userSelected: "",
      openChatBox: false
    };

    this.toggleUserList = this.toggleUserList.bind(this);
    this.callBack = this.callBack.bind(this);
    this.closeChatBox = this.closeChatBox.bind(this);
  }

  toggleUserList() {
    var userListDisplay = this.state.userListDisplay === "none" ? "" : "none";
    this.setState({
      userListDisplay: userListDisplay
    });
  }

  callBack(userToChatWith) {
    this.setState({
      userSelected: userToChatWith,
      openChatBox: false
    });

    setTimeout(
      function() {
        this.setState({ openChatBox: true });
      }.bind(this),
      200
    );
  }

  closeChatBox() {
    this.setState({
      openChatBox: false
    });
  }

  renderUserList() {
    if (this.props.data.loading) {
      return <div className="ui active centered inline loader" />;
    }

    if (this.props.data.error) {
      return <h3>Error : Server is Down :( </h3>;
    }

    return this.props.data.users.map(user => {
      if (this.props.welcomeName === user.name) {
        return null;
      }
      return (
        <SingleUser
          key={user.name}
          username={user.name}
          callBackProp={this.callBack}
        />
      );
    });
  }

  render() {
    return (
      <div>
        <div className="userList">
          <div className="viewUsers">
            <div className="ui button" onClick={this.toggleUserList}>
              View Online Users
            </div>
          </div>

          <div
            className="ui middle aligned divided list"
            style={{ display: this.state.userListDisplay }}
          >
            {this.renderUserList()}
          </div>
        </div>
        <ChatBox
          friendName={this.state.userSelected}
          welcomeName={this.props.welcomeName}
          openChatBox={this.state.openChatBox}
          closeChatBox={this.closeChatBox}
        />
      </div>
    );
  }
}

export default graphql(query)(UserList);
