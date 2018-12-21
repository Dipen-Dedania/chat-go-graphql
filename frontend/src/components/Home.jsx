import React, { Component } from "react";
import { Grid, Row, Col } from "react-bootstrap";
import SelfName from "./SelfName/SelfName";
import ChatBox from "./ChatBox/ChatBox";
import UserList from "./UserList/UserList";

export default class Home extends Component {
  constructor() {
    super();

    this.state = {
      chatBoxName: ""
    };
  }

  render() {
    return (
      <div className="homepage">
        <Grid>
          <Row>
            <Col xs={12} sm={6} md={6}>
              <SelfName />
            </Col>
            <Col xs={12} sm={6} md={6}>
              <UserList />
            </Col>
          </Row>

          <ChatBox name={this.state.chatBoxName} />
        </Grid>
      </div>
    );
  }
}
