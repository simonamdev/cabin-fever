package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// Websocket equivalent of CORS
func checkOrigin(r *http.Request) bool {
	return true
}

var upgrader = websocket.Upgrader{CheckOrigin: checkOrigin}

func handler(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("The time now is %s", time.Now().String())
	log.Println(message)
	// Allow cors
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, message)
}

// https://github.com/gorilla/websocket
func websocketHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Err During Upgrade", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

type GameState struct {
	X uint `json:"x"`
	Y uint `json:"y"`
}

func gameWebsocketHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Err During Upgrade", err)
		return
	}
	defer c.Close()
	gs := GameState{X: 0, Y: 0}
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		// So far we just expect UP, DOWN, LEFT or RIGHT
		dir := string(message)
		if dir == "UP" {
			gs.Y--
		} else if dir == "DOWN" {
			gs.Y++
		} else if dir == "LEFT" {
			gs.X--
		} else if dir == "RIGHT" {
			gs.X++
		}
		serialisedGameState, err := json.Marshal(gs)
		if err != nil {
			log.Println("write:", err)
			break
		}
		err = c.WriteMessage(mt, serialisedGameState)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

// Simplest web server example, taken from https://golang.org/doc/articles/wiki/
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/ws", websocketHandler)
	http.HandleFunc("/game", gameWebsocketHandler)
	log.Println("Starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
