import React, {
  useState,
  useEffect,
  useContext,
  FunctionComponent as FC
} from 'react';
import { useHistory, useParams } from 'react-router-dom';

import * as gen from './generated/graphql';

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

const useNewGame = () => {
  const hist = useHistory();
  const [ newGame, { data } ] = gen.useNewGameMutation();

  useEffect(() => {
    if (!data) return;
    hist.push(data.newGame.id);
  }, [hist, data])

  return newGame;
};

const useGame = (id: string) => {
  const vars = { variables: { id: id } }
  const query = gen.useCurrGameQuery(vars);
  gen.useWatchGameSubscription(vars);

  return query;
}

const useAddGuess = (id: string) => {
  const [ addGuess, { data } ] = gen.useAddGuessMutation();
  const [ event, setEvent ] = useState<Event | null>(null);

  useEffect(() => {
    if (!data) return;

    setEvent({
      correct: data.addGuess.correct,
      word: data.addGuess.word
    });

    const timer = setTimeout(() => setEvent(null), 3000)

    return () => clearTimeout(timer);
  }, [data]);

  const addGuessID = (guess: string) => addGuess({variables: {id, guess}});

  return [addGuessID, event] as const;
};

const Game: FC = () => {
  const user = useContext(UserContext);

  const { id } = useParams<{id: string}>();
  const { loading, error, data } = useGame(id);
  const [ addGuess, event ] = useAddGuess(id);
  const newGame = useNewGame();

  const [ guess, setGuess ] = useState("");

  if (loading) return <div>Loading</div>;
  if (error || !data) return <div>Error</div>;

  const game = data.game;

  return (
    <div>
      <div>{user.id} {user.name}</div>
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
            addGuess(guess);
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

        <button
          style={{margin: "5px", height: "40px", width: "130px"}}
          onClick={() => newGame()}
        >
          New Game
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

const Home: FC = () => {
  const newGame = useNewGame();
  return <button onClick={() => newGame()}>New Game!</button>
}

const UserContext = React.createContext<gen.User>(null as unknown as gen.User);

const EnsureUser: FC = ({ children }) => {
  const [ setUser, { data } ] = gen.useSetUserMutation();
  useEffect(() => { setUser(); }, [setUser]);

  // TODO: Better error handling
  if (!data) return <div>Loading</div>;

  return (
    <UserContext.Provider value={data.setUser}>
      {children}
    </UserContext.Provider>
  );
}

export { Game, Home, EnsureUser };
