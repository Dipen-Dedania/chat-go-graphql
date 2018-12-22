import React from "react";
import ReactDOM from "react-dom";
import "./styles/main.css";
import App from "./App";
import * as serviceWorker from "./serviceWorker";

// Apollo Files
import ApolloClient from "apollo-boost";
import { ApolloProvider } from "react-apollo";
import { createHttpLink } from "apollo-link-http";
import { InMemoryCache } from "apollo-cache-inmemory";

const client = new ApolloClient({
  uri: "http://192.168.1.145:9000/query"
});

// const httpLink = createHttpLink({
//   uri: "http://192.168.1.145:9000/query"
// });

// const client = new ApolloClient({
//   link: httpLink,
//   cache: new InMemoryCache()
// });

ReactDOM.render(
  <ApolloProvider client={client}>
    <App />
  </ApolloProvider>,
  document.getElementById("root")
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: http://bit.ly/CRA-PWA
serviceWorker.unregister();
