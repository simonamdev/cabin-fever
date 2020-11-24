// import * as PIXI from 'pixi.js'
import * as PIXI from 'pixi.js-legacy'

const app = new PIXI.Application();

document.body.appendChild(app.view);

const graphics = new PIXI.Graphics();

graphics.beginFill(0xFFFF00);
graphics.lineStyle(5, 0xFF0000);
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

let text = new PIXI.Text(`X: ${playerPosition.x} Y: ${playerPosition.y}`, { fontFamily: 'Arial', fontSize: 24, fill: 0xff1010, align: 'center' });
text.anchor.x = 0.5;
app.stage.addChild(text);

// Listen for frame updates
app.ticker.add(() => {
    graphics.rotation += 0.01;
    // Update position
    graphics.x = playerPosition.x;
    graphics.y = playerPosition.y;
    text.x = playerPosition.x;
    text.y = playerPosition.y
    text.text = `X: ${playerPosition.x} Y: ${playerPosition.y}`
    text.updateText(true);
});

enum Direction {
    Up = 'UP',
    Down = 'DOWN',
    Left = 'LEFT',
    Right = 'RIGHT'
}

const ws = new WebSocket("ws://localhost:8080/game");

const getDir = (key: string): Direction => {
    if (key === 'w') {
        return Direction.Up;
    } else if (key === 'a') {
        return Direction.Left;
    } else if (key === 's') {
        return Direction.Down;
    } else if (key === 'd') {
        return Direction.Right;
    }
}

ws.onopen = (e) => {
    console.log('Connection Setup');

    // Setup keyboard events
    const body = document.getElementById('body');
    if (body) {
        body.onkeydown = (e) => {
            console.log(e.key);
            if (['w', 'a', 's', 'd'].includes(e.key)) {
                ws.send(getDir(e.key));
            }
        }
    }
}

ws.onmessage = (e) => {
    // When we receive the updated position, update the client side state
    const { data } = e;
    console.log(data);

    const position: PlayerPosition = JSON.parse(data); // Ideally we do validation here
    // gameRoot.innerText = JSON.stringify(position);
    playerPosition.x = position.x * 10;
    playerPosition.y = position.y * 10;
}



