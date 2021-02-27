import React, { useEffect, useState } from 'react';
import './App.css';
import {
  BrowserRouter as Router,
  Route,
  useHistory,
  match
} from "react-router-dom";

const randomFrom = (chars: string) =>
  chars.charAt(Math.floor(Math.random() * chars.length))

const alpha = () => randomFrom("abcdefghijklmnopqrstuvwxyz");
const vowel = () => randomFrom("aeiou");

const shuffle = (arr: any[]) => {
  arr = arr.slice(0);
  for (let i = arr.length - 1; i > 0; i--) {
    let rand = Math.floor(Math.random() * (i + 1));
    [arr[i], arr[rand]] = [arr[rand], arr[i]];
  }

  return arr;
}

const generate = (count: number) => {
  if (count < 1) {
    return "";
  }

  let chars = [vowel()];
  while (chars.length < count) {
    let char = alpha();
    if (!chars.includes(char)) {
      chars.push(char);
    }
  }

  return shuffle(chars).join("");
}

const Home = () => {
  const hist = useHistory();

  return <button onClick={() => hist.push(generate(7))}>
    New Game
  </button>;
}

type PlayProps = { match: match<{chars: string}>; }

type State = {
  chars: string;
  charArr: string[];
  guess: string;
  words: string[];
  correct: string[];
};

const addChar = (c: string, state: State) => ({
  ...state,
  guess: state.guess + c
})

const enter = (state: State) => {
  const words = state.words.filter(word => word !== state.guess);
  const correct =
    words.length !== state.words.length
      ? [...state.correct, state.guess]
      : state.correct

  return ({
    ...state,
    words,
    correct,
    guess: "",
  })
}

const backspace = (state: State) => ({
  ...state,
  guess: state.guess.slice(0, state.guess.length - 1)
})

const shuffleState = (state: State) => ({
  ...state,
  charArr: [state.charArr[0], ...shuffle(state.charArr.slice(1))]
})

const Play = ({ match }: PlayProps) => {
  const chars = match.params.chars;
  const [state, setState] = useState<State>();

  console.log(state);

  useEffect(() => {
    if (!state || state.chars !== chars) {
      fetch("http://localhost:8000/", {
        method: "POST",
        body: chars,
      })
      .then(res => res.text())
      .then(body => setState({
        chars,
        charArr: chars.split(""),
        guess: "",
        words: body.split("\n"),
        correct: []
      }))
    }
  }, [chars, state]);

  if (state === undefined) {
    return <div>Loading</div>
  }

  return <div>
    <div>&nbsp;{state.guess}</div>

    {state.charArr.map(c =>
      <button key={c} onClick={() => setState(addChar(c, state))}>{c}</button>
    )}

    <button
      onClick={() => setState(enter(state))}
      disabled={state.guess.length <= 3}
    >
      Enter
    </button>

    <button onClick={() => setState(backspace(state))}>Delete</button>

    <button onClick={() => setState(shuffleState(state))}>Shuffle</button>

    <ul>
      {state.correct.map(word => 
        <li key={word}>{word}</li>
      )}
    </ul>
  </div>;
}

const App = () => {
  return <Router>
    <Route exact path="/">
      <Home/>
    </Route>

    <Route path="/:chars" component={Play}/>
  </Router>
}

export default App;
