import Clubs9 from "../assets/cardImages/9 of Clubs.png";
import Clubs10 from "../assets/cardImages/10 of Clubs.png";
import ClubsJ from "../assets/cardImages/Jack of Clubs.png";
import ClubsQ from "../assets/cardImages/Queen of Clubs.png";
import ClubsK from "../assets/cardImages/King of Clubs.png";
import ClubsA from "../assets/cardImages/Ace of Clubs.png";

import Diamonds9 from "../assets/cardImages/9 of Diamonds.png";
import Diamonds10 from "../assets/cardImages/10 of Diamonds.png";
import DiamondsJ from "../assets/cardImages/Jack of Diamonds.png";
import DiamondsQ from "../assets/cardImages/Queen of Diamonds.png";
import DiamondsK from "../assets/cardImages/King of Diamonds.png";
import DiamondsA from "../assets/cardImages/Ace of Diamonds.png";

import Spades9 from "../assets/cardImages/9 of Spades.png";
import Spades10 from "../assets/cardImages/10 of Spades.png";
import SpadesJ from "../assets/cardImages/Jack of Spades.png";
import SpadesQ from "../assets/cardImages/Queen of Spades.png";
import SpadesK from "../assets/cardImages/King of Spades.png";
import SpadesA from "../assets/cardImages/Ace of Spades.png";

import Hearts9 from "../assets/cardImages/9 of Hearts.png";
import Hearts10 from "../assets/cardImages/10 of Hearts.png";
import HeartsJ from "../assets/cardImages/Jack of Hearts.png";
import HeartsQ from "../assets/cardImages/Queen of Hearts.png";
import HeartsK from "../assets/cardImages/King of Hearts.png";
import HeartsA from "../assets/cardImages/Ace of Hearts.png";

function PlayerHand({ index1, index2, index3, index4, index5 }) {
  // clubs
  const clubs9 = Clubs9;
  const clubs10 = Clubs10;
  const clubsJ = ClubsJ;
  const clubsQ = ClubsQ;
  const clubsK = ClubsK;
  const clubsA = ClubsA;

  // diamonds
  const diamonds9 = Diamonds9;
  const diamonds10 = Diamonds10;
  const diamondsJ = DiamondsJ;
  const diamondsQ = DiamondsQ;
  const diamondsK = DiamondsK;
  const diamondsA = DiamondsA;

  // spades
  const spades9 = Spades9;
  const spades10 = Spades10;
  const spadesJ = SpadesJ;
  const spadesQ = SpadesQ;
  const spadesK = SpadesK;
  const spadesA = SpadesA;

  // hearts
  const hearts9 = Hearts9;
  const hearts10 = Hearts10;
  const heartsJ = HeartsJ;
  const heartsQ = HeartsQ;
  const heartsK = HeartsK;
  const heartsA = HeartsA;

  const deck = [
    clubs9,
    clubs10,
    clubsJ,
    clubsQ,
    clubsK,
    clubsA,
    diamonds9,
    diamonds10,
    diamondsJ,
    diamondsQ,
    diamondsK,
    diamondsA,
    spades9,
    spades10,
    spadesJ,
    spadesQ,
    spadesK,
    spadesA,
    hearts9,
    hearts10,
    heartsJ,
    heartsQ,
    heartsK,
    heartsA,
  ];

  let card1, card2, card3, card4, card5;
  function dealCards(a, b, c, d, e) {
    card1 = deck[a];
    card2 = deck[b];
    card3 = deck[c];
    card4 = deck[d];
    card5 = deck[e];
  }

  dealCards(index1, index2, index3, index4, index5);

  return (
    <div>
      <img src={card1} className="base" width="75" alt=""></img>
      <img src={card2} className="base" width="75" alt=""></img>
      <img src={card3} className="base" width="75" alt=""></img>
      <img src={card4} className="base" width="75" alt=""></img>
      <img src={card5} className="base" width="75" alt=""></img>
    </div>
  );
}

export default PlayerHand;
