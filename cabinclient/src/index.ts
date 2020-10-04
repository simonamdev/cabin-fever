const message: string = "Hello, World!";

console.log(message);

setInterval(async () => {
    const response = await fetch('http://localhost:8080', {
        headers: {
            'Content-Type': 'text/plain'
        },
    });
    const text = await response.text();
    document.getElementById('root').innerText = text;
}, 1000)