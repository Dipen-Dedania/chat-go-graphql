import React, { Component } from "react"; //React

import { gql } from "apollo-boost"; //GraphQL
import { graphql } from "react-apollo";

const mutation = gql`
  mutation adduser($name: String!) {
    joinUser(name: $name) {
      name
    }
  }
`;

class SelfName extends Component {
  constructor() {
    super();
    this.state = {
      // The Entered name
      enterNameDisplay: "",
      // Show or Hide Enter Name Box
      welcomeDisplay: "none",
      // The Welcome Message Name
      welcomeName: ""
    };

    this.nameEntered = this.nameEntered.bind(this);
  }

  //This function is called when the name is entered and we hit enter
  nameEntered(e) {
    if (e.key === "Enter") {
      if (e.target.value === "") {
        alert("Please Enter a Name");
      } else if (e.target.value.length > 10) {
        alert("Please Enter a Name between 1-10 Character");
      } else {
        this.setState(
          {
            welcomeName: e.target.value,
            welcomeDisplay: "",
            enterNameDisplay: "none"
          },
          () => {
            this.props.displayList(this.state.welcomeName);
          }
        );
        this.props.mutate({
          variables: {
            name: e.target.value
          }
        });
      }
    }
  }

  render() {
    return (
      <div className="self-name">
        <div
          className="enter-name"
          style={{ display: this.state.enterNameDisplay }}
        >
          <input className="cbox" type="checkbox" />
          <label className="add" htmlFor="cbox">
            Enter Your Name
          </label>
          <input
            className="message"
            type="text"
            onKeyPress={this.nameEntered}
          />
        </div>
        <div
          className="welcome"
          style={{
            display: this.state.welcomeDisplay
          }}
        >
          Welcome, {this.state.welcomeName}
        </div>
      </div>
    );
  }
}

export default graphql(mutation)(SelfName);
