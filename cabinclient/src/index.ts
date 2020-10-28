// HTTP Method
/*
setInterval(async () => {
    const response = await fetch('http://localhost:8080', {
        headers: {
            'Content-Type': 'text/plain'
        },
    });
    const text = await response.text();
    document.getElementById('root').innerText = text;
}, 5000);
*/

// WS Method
const testws = new WebSocket("ws://localhost:8080/ws");

// On open, send the date every 1 second
testws.onopen = (e) => {
    setInterval(() => {
        const date = new Date();
        console.log(`Sending date: ${date} `);
        testws.send(date.toISOString())
    }, 1000)
}

// On receive, update the DOM
testws.onmessage = (ev) => {
    document.getElementById('rootws').innerText = `WS Data received: ${ev.data}`;
}

enum Direction {
    Up = 'UP',
    Down = 'DOWN',
    Left = 'LEFT',
    Right = 'RIGHT'
}

interface PlayerPosition {
    x: number;
    y: number;
}

const position: PlayerPosition = { x: 0, y: 0 };
const ws = new WebSocket("ws://localhost:8080/game");
const body = document.getElementById('body');
const gameRoot = document.getElementById('game');

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
    body.onkeydown = (e) => {
        console.log(e.key);
        if (['w', 'a', 's', 'd'].includes(e.key)) {
            ws.send(getDir(e.key));
        }
    }
}

ws.onmessage = (e) => {
    // When we receive the updated position, update the client side state
    const { data } = e;
    const position: PlayerPosition = JSON.parse(data); // Ideally we do validation here
    gameRoot.innerText = JSON.stringify(position);
}



