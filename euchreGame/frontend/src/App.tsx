// Wails setup based on tutorial https://www.youtube.com/watch?v=HZRvQVtlOHo
import "./App.css";

import { useEffect } from "react";

import Clubs10 from "./assets/cardImages/10 of Clubs.png";
import DiamondsJ from "./assets/cardImages/Jack of Diamonds.png";
import CardBack from "./assets/cardImages/back.png";
import DiamondsA from "./assets/cardImages/Ace of Diamonds.png";
import HeartsK from "./assets/cardImages/King of Hearts.png";

import { OnFileDrop } from "../wailsjs/runtime";

function App() {
  const card1 = Clubs10;
  const card2 = DiamondsJ;
  const card3 = DiamondsA;
  const card4 = HeartsK;
  const card5 = CardBack;

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
        </div>
      </section>

      <section id="center">
        <div className="game-table">Game table</div>
      </section>

      <section id="next-steps">
        {/* Player hand */}
        <div>
          {/* Five cards */}
          <div>
            <img src={card1} className="base" width="75" alt=""></img>
            <img src={card2} className="base" width="75" alt=""></img>
            <img src={card3} className="base" width="75" alt=""></img>
            <img src={card4} className="base" width="75" alt=""></img>
            <img src={card5} className="base" width="75" alt=""></img>
          </div>
        </div>
      </section>

      <div className="ticks"></div>

      <div className="ticks"></div>
      <section id="spacer"></section>
    </>
  );
}

export default App;
