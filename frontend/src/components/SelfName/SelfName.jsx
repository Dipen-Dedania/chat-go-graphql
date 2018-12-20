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

        // CODE TO ADD ENTERED NAME TO USER LIST
      }
    }
  }

  render() {
    return (
      <div className="self-name">
        {/* <Query
    query={gql`
      {
        allCourses {
          id
          title
          author
          description
          topic
          url
        }
      }
    `}
  >
    {({ loading, error, data }) => {
      if (loading) return <p>Loading...</p>;
      if (error) return <p>Error :(</p>;
      return data.allCourses.map(({ id, title, author, description, topic, url }) => (
        <div key={id}>
          <p>{`${title} by ${author}`}</p>
        </div>
      ));
    }}
  </Query> */}

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
