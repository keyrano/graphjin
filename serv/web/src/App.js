import React from "react";

import { createGraphiQLFetcher } from "@graphiql/toolkit";
import { GraphiQL } from "graphiql";

import "graphiql/graphiql.css";

const fetcher = createGraphiQLFetcher({
  url: `${window.location.protocol}//${window.location.host}/api/v1/graphql`,
  subscriptionUrl: `ws://${window.location.host}/api/v1/graphql`,
});

// import GitHubButton from "react-github-btn";

// const url = `${window.location.protocol}//${window.location.host}/api/v1/graphql`;
// const subscriptionUrl = ;

const defaultQuery = `
# Welcome to GraphJin Web

# Use this editor to build and test your GraphQL queries
# Set a query name if you want the query saved to the
# allow list to use in production

query {
  users(id: "3") {
    id
    full_name
    email
  }
}
`;

const App = () => <GraphiQL fetcher={fetcher} query={defaultQuery} />;

export default App;
