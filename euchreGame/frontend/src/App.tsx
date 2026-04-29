// Wails setup based on tutorial https://www.youtube.com/watch?v=HZRvQVtlOHo
import "./App.css";

import { useEffect } from "react";

import gameMusic from "./assets/songs/High Stakes, Low Voices.wav";

import GameTable from "./assets/images/vecteezy_green-casino-poker-table-texture-game-background_24232274.jpg";

import PlayerHand from "./components/PlayerHand";

// import { RandomIndices } from "../wailsjs/go/main/App";

// import { OnFileDrop } from "../wailsjs/runtime";

function App() {
  // let card1,
  //   card2,
  //   card3,
  //   card4,
  //   card5 = RandomIndices();

  function play() {
    new Audio(gameMusic).play();
  }

  const gameTableStyle = {
    backgroundImage: `url(${GameTable})`,
    backgroundSize: "cover",
    backgroundPosition: "center",
    height: "400px",
    width: "100%",
  };

  useEffect(() => {
    console.log("this ran");
  });

  return (
    <>
      <section id="center">
        <div className="hero">
          {/* <img src={} className="base" width="170" height="179" alt="" /> */}
        </div>
        <div>
          <h1>Welcome to the Euchre Game</h1>
          <button onClick={play}>Play game music</button>
        </div>
      </section>

      <section id="center" style={gameTableStyle}>
        <div className="game-table"></div>
      </section>

      <section id="next-steps">
        <PlayerHand index1={0} index2={18} index3={8} index4={23} index5={13} />
      </section>

      <div className="ticks"></div>

      <div className="ticks"></div>
      <section id="spacer"></section>
    </>
  );
}

export default App;
