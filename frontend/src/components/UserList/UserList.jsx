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
      userSelected: ""
    };

    this.toggleUserList = this.toggleUserList.bind(this);
    this.callBack = this.callBack.bind(this);
  }

  toggleUserList() {
    var userListDisplay = this.state.userListDisplay === "none" ? "" : "none";
    this.setState({
      userListDisplay: userListDisplay
    });
  }

  callBack(userToChatWith) {
    this.setState({
      userSelected: userToChatWith
    });
  }

  renderUserList() {
    if (this.props.data.loading) {
      return <div>Loading...</div>;
    }

    return this.props.data.users.map(user => {
      return (
        <SingleUser
          key={user.id}
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
        <ChatBox name={this.state.userSelected} />
      </div>
    );
  }
}

export default graphql(query)(UserList);
