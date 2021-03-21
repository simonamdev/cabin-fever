package main

import (
	"cabinserver/game"
	"cabinserver/physics"
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Embed all static files

//go:embed static/index.html
var indexHTML string

//go:embed static
var staticFiles embed.FS

// Websocket equivalent of CORS
func checkOrigin(r *http.Request) bool {
	return true
}

var upgrader = websocket.Upgrader{CheckOrigin: checkOrigin}

func serveFrontendHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(indexHTML))
	}
}

func gameWebsocketHandler(addPlayerC chan game.Player, updatesC chan physics.Direction) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal("Err During Upgrade", err)
			return
		}
		defer conn.Close()

		// On open of the connection, pass in the new player
		addPlayerC <- game.Player{Conn: conn}

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", message)

			// Pass received direction to the updates channel
			// This part is simply because we need to work around the type
			dir := string(message)
			if dir == "UP" {
				updatesC <- physics.Up
			} else if dir == "DOWN" {
				updatesC <- physics.Down
			} else if dir == "LEFT" {
				updatesC <- physics.Left
			} else if dir == "RIGHT" {
				updatesC <- physics.Right
			}
		}
	}
}

// Simplest web server example, taken from https://golang.org/doc/articles/wiki/
func main() {
	// Setup embedded file system
	fsys, err := fs.Sub(staticFiles, "static")
	if err != nil {
		panic(err)
	}
	// Setup game
	config := game.Config{TicksPerSecond: 10}
	addPlayerC := make(chan game.Player)
	_, updatesC := game.RunGameLoop(config, addPlayerC)
	http.HandleFunc("/", serveFrontendHandler())
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.FS(fsys))))
	http.HandleFunc("/game", gameWebsocketHandler(addPlayerC, updatesC))
	log.Println("Starting server")
	go log.Fatal(http.ListenAndServe(":8080", nil))
}
