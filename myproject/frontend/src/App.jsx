import { useState } from "react";
import logo from "./assets/images/euchre_logo.jpg";
import "./App.css";
import { Greet } from "../wailsjs/go/main/App";
import GameRoom from "./components/gameRoom/gameRoom";

function App() {
  const [resultText, setResultText] = useState(
    "Please enter your name below 👇",
  );
  const [name, setName] = useState("");
  const [gameActive, setGameActive] = useState(false);
  const [exit, setExit] = useState(false);
  const updateName = (e) => setName(e.target.value);
  const updateResultText = (result) => setResultText(result);

  function play() {
    Greet(name).then(updateResultText);
    setGameActive(true);
    setExit(false);
  }

  function exitGame() {
    setExit(true);
  }

  if (gameActive && !exit) {
    // Game Screen
    return (
      <div>
        <div>Let's play some Euchre!</div>
        <button className="btn" onClick={exitGame}>
          Exit
        </button>
        <GameRoom />
      </div>
    );
  } else {
    return (
      // Home Screen
      <div id="App">
        <img src={logo} id="logo" alt="logo" />
        <div id="result" className="result">
          {resultText}
        </div>
        <div id="input" className="input-box">
          <input
            id="name"
            className="input"
            onChange={updateName}
            autoComplete="off"
            name="input"
            type="text"
          />
          <button className="btn" onClick={play}>
            Play
          </button>
        </div>
      </div>
    );
  }
}

export default App;
