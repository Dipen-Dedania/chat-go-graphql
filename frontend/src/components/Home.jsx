import React, { Component } from "react";
import { Grid, Row, Col } from "react-bootstrap";
import SelfName from "./SelfName/SelfName";
import ChatBot from "./ChatBox/ChatBot";
import UserList from "./UserList/UserList";

export default class Home extends Component {
  render() {
    return (
      <div className="homepage">
        <Grid>
          <Row>
            <Col xs={12} sm={5} md={5}>
              <SelfName />
            </Col>
            <Col xs={12} sm={7} md={7}>
              <UserList />
            </Col>
          </Row>

          <ChatBot />
        </Grid>
      </div>
    );
  }
}
