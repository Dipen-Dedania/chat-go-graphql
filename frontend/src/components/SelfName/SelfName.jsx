import React, { Component } from "react";
import { Query } from "react-apollo";
import gql from "graphql-tag";

export default class SelfName extends Component {
  constructor() {
    super();
    this.state = {
      enterNameDisplay: "",

      welcomeName: "",
      welcomeDisplay: "none"
    };

    this.nameEntered = this.nameEntered.bind(this);
  }

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
      }
    }
  }

  render() {
    return (
      <div className="selfname">
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
