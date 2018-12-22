import React, { Component } from "react";
import gql from "graphql-tag";
import { graphql } from "react-apollo";

class SelfName extends Component {
  constructor() {
    super();
    this.state = {
      // The Entered name
      enterNameDisplay: "",
      // The Welcome MEssage Name
      welcomeName: "",
      // Show or Hide Enter Name Box
      welcomeDisplay: "none"
    };

    this.nameEntered = this.nameEntered.bind(this);
  }

  //This function is called when the name is entered and we hit enter
  nameEntered(e) {
    if (e.key === "Enter") {
      if (e.target.value === "") {
        alert("Please Enter a Name");
      } else {
        this.setState({
          welcomeName: e.target.value,
          welcomeDisplay: "",
          enterNameDisplay: "none"
        });
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

const mutation = gql`
  mutation AddName($name: String!) {
    userjoin(name: $name) {
      name
    }
  }
`;

export default graphql(mutation)(SelfName);
