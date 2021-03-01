import React, { useState } from 'react';
import {
  ApolloClient,
  InMemoryCache,
  gql,
  ApolloProvider,
  useQuery,
  useMutation,
  useSubscription,
  HttpLink,
  split
} from '@apollo/client';

import { WebSocketLink } from '@apollo/client/link/ws';
import { getMainDefinition } from '@apollo/client/utilities';
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

const CURR_GAME = gql`
  query {
    game(id: "asdf") {
      id,
      letters,
      guessed
    }
  }
`;

const ADD_GUESS = gql`
  mutation addGuess($guess: String!) {
    addGuess(guess: $guess) {
      id,
      letters,
      guessed
    }
  }
`

const WATCH_GAME = gql`
  subscription {
    watchGame {
      id,
      letters,
      guessed
    }
  }
`

const Test = () => {
  const { loading, error, data } = useQuery(CURR_GAME);
  const [addGuess] = useMutation(ADD_GUESS);
  useSubscription(WATCH_GAME);
  const [guess, setGuess] = useState("");

  if (loading) return <div>Loading</div>;
  if (error) return <div>Error</div>;

  const game = data.game;

  return (
    <div>
      <div>&nbsp;{guess}</div>

      {game.letters.split("").map((c: string) =>
        <button
          key={c}
          onClick={() => setGuess(guess + c)}
          style={{
            margin: "5px",
            height: "40px",
            width: "50px",
            fontSize: "18px",
          }}
        >
          {c}
        </button>
      )}

      <div>
        <button
          disabled={guess.length <= 3}
          style={{margin: "5px", height: "40px", width: "130px"}}
          onClick={() => {
            addGuess({ variables: { guess }});
            setGuess("");
          }}
        >
          Guess!
        </button>

        <button
          style={{margin: "5px", height: "40px", width: "130px"}}
          onClick={() => setGuess(guess.slice(0, guess.length - 1))}
        >
          Delete
        </button>
      </div>

      <ul>
        {game.guessed.map((word: string) => 
          <li key={word}>{word}</li>
        )}
      </ul>

    </div>
  );
};

const App = () => {
  return (
    <ApolloProvider client={client}>
      <Test/>
    </ApolloProvider>
  );
}

export default App;
