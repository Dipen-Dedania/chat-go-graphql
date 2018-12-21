import React, { Component } from "react";
import { Query } from "react-apollo";
import gql from "graphql-tag";

import SingleUser from "./SingleUser";

export default class UserList extends Component {
  constructor() {
    super();
    this.state = {
      userListDisplay: "none"
    };

    this.toggleUserList = this.toggleUserList.bind(this);
  }

  toggleUserList() {
    var userListDisplay = this.state.userListDisplay === "none" ? "" : "none";
    this.setState({
      userListDisplay: userListDisplay
    });
  }

  render() {
    return (
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
          {/* QUERY */}
          <Query
            query={gql`
              {
                users {
                  id
                  name
                }
              }
            `}
          >
            {({ loading, error, data }) => {
              if (loading) return <p>Loading...</p>;
              if (error) return <p>Error :(</p>;

              return data.users.map(({ id, name }) => (
                <SingleUser key={id} username={name} />
              ));
            }}
          </Query>
        </div>
      </div>
    );
  }
}
