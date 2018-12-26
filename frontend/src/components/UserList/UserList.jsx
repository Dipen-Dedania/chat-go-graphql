import React, { Component } from "react";

//GraphQL
import { Query } from "react-apollo";
import gql from "graphql-tag";

//Components Import
import ChatBox from "../ChatBox/ChatBox";
import SingleUser from "./SingleUser";

// Query
const USER_LIST = gql`
  {
    users {
      name
    }
  }
`;

class UserList extends Component {
  constructor(props) {
    super(props);
    this.state = {
      userListDisplay: "none",
      userSelected: "",
      refresh: props.refresh
    };

    this.toggleUserList = this.toggleUserList.bind(this);
    this.callBack = this.callBack.bind(this);
    console.log("Hello");
  }

  toggleUserList() {
    var userListDisplay = this.state.userListDisplay === "none" ? "" : "none";
    this.setState({
      userListDisplay: userListDisplay
    });
    console.log(this.state.refresh);
  }

  callBack(userToChatWith) {
    this.setState({
      userSelected: userToChatWith
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
            <Query query={USER_LIST}>
              {({ loading, error, data }) => {
                if (loading) return <div>Loading...</div>;
                if (error) return <div>Error</div>;

                const userList = data.users;

                return userList.map(user => {
                  return (
                    <SingleUser
                      key={user.name}
                      username={user.name}
                      callBackProp={this.callBack}
                    />
                  );
                });
              }}
            </Query>
          </div>
        </div>
        <ChatBox name={this.state.userSelected} />
      </div>
    );
  }
}

export default UserList;
