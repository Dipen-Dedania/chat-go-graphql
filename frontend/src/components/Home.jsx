//React Import
import React, { Component } from "react";

//BOOTSTRAP IMPORT
import { Grid, Row, Col } from "react-bootstrap";

//Imported SelfName and UserList
import SelfName from "./SelfName/SelfName";
import UserList from "./UserList/UserList";

export default class Home extends Component {
  constructor() {
    super();
    this.state = {
      success: false
    };
    this.success = this.success.bind(this);
  }

  success(flag) {
    console.log(flag);
    this.setState({
      success: true
    });
  }
  render() {
    console.log(this.state.success);
    return (
      <div className="homepage">
        <Grid>
          <Row>
            <Col xs={12} sm={6} md={6}>
              {/* Enter Your Name Component */}
              <SelfName success={this.success} />
            </Col>
            <Col xs={12} sm={6} md={6}>
              {/* Users Online Components */}
              <UserList />
            </Col>
          </Row>
        </Grid>
      </div>
    );
  }
}
