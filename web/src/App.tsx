import React from 'react';
import {
  ApolloClient,
  InMemoryCache,
  ApolloProvider,
  HttpLink,
  split
} from '@apollo/client';

import { WebSocketLink } from '@apollo/client/link/ws';
import { getMainDefinition } from '@apollo/client/utilities';
import { Game } from './Game';
import './App.css';

const wsLink = new WebSocketLink({
  uri: "ws://lambda.olympus:8080/query",
  options: {
    reconnect: true
  }
});

const httpLink = new HttpLink({
  uri: "http://lambda.olympus:8080/query",
});

const splitLink = split(
  ({ query }) => {
    const definition = getMainDefinition(query);
    return (
      definition.kind === "OperationDefinition"
      && definition.operation === "subscription"
    )
  },
  wsLink,
  httpLink
);

const client = new ApolloClient({
  link: splitLink,
  cache: new InMemoryCache()
});

const App = () => {
  return (
    <ApolloProvider client={client}>
      <Game/>
    </ApolloProvider>
  );
}

export default App;
