import React from 'react';
import './App.css';

const test = async () => {
  console.log("test started");
  const res = await fetch("http://localhost:8000/");
  console.log(res);
  console.log(await res.text());
}

function App() {
  return <button onClick={test}>Test</button>;
}

export default App;
