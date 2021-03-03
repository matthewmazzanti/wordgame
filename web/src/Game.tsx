import React, { useState, useEffect } from 'react';
import {
  gql,
  useQuery,
  useMutation,
  useSubscription,
} from '@apollo/client';

const CURR_GAME = gql`
  query {
    game(id: "asdf") {
      id,
      letters,
      correct,
      incorrect,
      total
    }
  }
`;

const ADD_GUESS = gql`
  mutation addGuess($guess: String!) {
    addGuess(guess: $guess) {
      correct,
      word,
      game {
        id,
        letters,
        correct,
        incorrect,
      }
    }
  }
`;

const WATCH_GAME = gql`
  subscription {
    watchGame {
      correct,
      word,
      game {
        id,
        letters,
        correct,
        incorrect,
      }
    }
  }
`;

type Event = {
  correct: boolean,
  word: string,
};

type ShowEventProps = {
  event: Event | null,
};

const ShowEvent = ({ event }: ShowEventProps) => {
  if (!event) return <div>&nbsp;</div>;

  return event.correct
    ? <div>Correct!</div>
    : <div>{event.word} not in word list</div>;
}

const useAddGuess = () => {
  const [ addGuess, { data } ] = useMutation(ADD_GUESS);
  const [ event, setEvent ] = useState<Event | null>(null);

  useEffect(() => {
    if (data === undefined) return;

    setEvent({
      // @ts-ignore
      correct: data.addGuess.correct,
      word: data.addGuess.word
    });

    const timer = setTimeout(() => setEvent(null), 3000)

    return () => clearTimeout(timer);
  }, [data]);

  return [addGuess, event] as const;
}

const Game = () => {
  const { loading, error, data } = useQuery(CURR_GAME);
  const [ addGuess, event ] = useAddGuess();
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

      <ShowEvent event={event}/>

      <div>
        <div>Correct: ({game.correct.length}/{game.total})</div>
        <ul>
          {game.correct.map((word: string) => 
            <li key={word}>{word}</li>
          )}
        </ul>
      </div>

      <div>
        <div>Incorrect:</div>
        <ul>
          {game.incorrect.map((word: string) => 
            <li key={word}>{word}</li>
          )}
        </ul>
      </div>
    </div>
  );
};

export { Game };
