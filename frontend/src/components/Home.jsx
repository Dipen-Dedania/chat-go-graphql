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
      userList: "",
      welcomeName: ""
    };
    this.displayList = this.displayList.bind(this);
  }

  displayList(welcomeName) {
    if (welcomeName) {
      this.setState(
        {
          welcomeName: welcomeName
        },
        () => {
          this.setState({
            userList: <UserList welcomeName={this.state.welcomeName} />
          });
        }
      );
    }
  }

  render() {
    return (
      <div className="homepage">
        <Grid>
          <Row>
            <Col xs={12} sm={6} md={6}>
              {/* Enter Your Name Component */}
              <SelfName displayList={this.displayList} />
            </Col>
            <Col xs={12} sm={6} md={6}>
              {/* Users Online Components */}
              {this.state.userList}
            </Col>
          </Row>
        </Grid>
      </div>
    );
  }
}
