// import * as PIXI from 'pixi.js'
import * as PIXI from "pixi.js-legacy";

const app = new PIXI.Application();

document.body.appendChild(app.view);

const graphics = new PIXI.Graphics();

graphics.beginFill(0xffff00);
graphics.lineStyle(5, 0xff0000);
graphics.drawRect(0, 0, 300, 300);
graphics.pivot.x = 150;
graphics.pivot.y = 150;
graphics.position.x = 300;
graphics.position.y = 300;
app.stage.addChild(graphics);

interface PlayerPosition {
  x: number;
  y: number;
}

const playerPosition: PlayerPosition = { x: 0, y: 0 };

let lastKeyText = new PIXI.Text(`Last Key: N/A`, {
  fontFamily: "Arial",
  fontSize: 24,
  fill: 0xff1010,
  align: "center",
});
let positionText = new PIXI.Text(
  `X: ${playerPosition.x} Y: ${playerPosition.y}`,
  { fontFamily: "Arial", fontSize: 24, fill: 0xff1010, align: "center" }
);

lastKeyText.anchor.x = 0.5;
lastKeyText.x = 500;
lastKeyText.y = 500;
positionText.anchor.x = 0.5;
app.stage.addChild(lastKeyText);
app.stage.addChild(positionText);

// Listen for frame updates
app.ticker.add(() => {
  graphics.rotation += 0.01;
  // Update position
  graphics.x = playerPosition.x;
  graphics.y = playerPosition.y;
  positionText.x = playerPosition.x;
  positionText.y = playerPosition.y;
  positionText.text = `X: ${playerPosition.x} Y: ${playerPosition.y}`;
  positionText.updateText(true);
});

enum Direction {
  Up = "UP",
  Down = "DOWN",
  Left = "LEFT",
  Right = "RIGHT",
}

const url: string = `ws://${
  window.location.host.includes(":")
    ? window.location.host.split(":")[0]
    : window.location.host
}:8080/game`;
console.log(url);

const ws = new WebSocket(url);

const getDir = (key: string): Direction => {
  if (key === "w") {
    return Direction.Up;
  } else if (key === "a") {
    return Direction.Left;
  } else if (key === "s") {
    return Direction.Down;
  } else if (key === "d") {
    return Direction.Right;
  }
};

ws.onopen = (e) => {
  console.log("Connection Setup");

  // Setup keyboard events
  const body = document.getElementById("body");
  if (body) {
    body.onkeydown = (e) => {
      console.log(e.key);
      lastKeyText.text = `Last Key: ${e.key}`;
      lastKeyText.updateText(true);
      if (["w", "a", "s", "d"].includes(e.key)) {
        ws.send(getDir(e.key));
      }
    };
  }
};

ws.onmessage = (e) => {
  // When we receive the updated position, update the client side state
  const { data } = e;
  console.log(data);

  const position: PlayerPosition = JSON.parse(data); // Ideally we do validation here
  // gameRoot.innerText = JSON.stringify(position);
  playerPosition.x = position.x * 10;
  playerPosition.y = position.y * 10;
};
