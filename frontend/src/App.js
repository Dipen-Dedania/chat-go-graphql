import React, { Component } from "react";
import Home from "./components/Home";
import ApolloClient from "apollo-boost";
import { ApolloProvider } from "react-apollo";

const client = new ApolloClient({
  uri: "http://192.168.1.145:9000/query"
});

class App extends Component {
  render() {
    return (
      <ApolloProvider client={client}>
        <div className="App">
          <Home />
        </div>
      </ApolloProvider>
    );
  }
}

export default App;
