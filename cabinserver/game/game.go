package game

import (
	"cabinserver/physics"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

// State represents the state of the entire game
type State struct {
	Players []Player `json:"players"`
}

// Player represents a single player in the game
type Player struct {
	physics.Position
	physics.Direction
	Conn *websocket.Conn `json:"-"`
}

// Config represents the parameters around which the game will operate
type Config struct {
	TicksPerSecond uint
}

func SendUpdateToClient(c *websocket.Conn, gameState State) {
	serialisedGameState, err := json.Marshal(gameState.Players[0])
	if err != nil {
		log.Println("write:", err)
		panic(err)
	}
	err = c.WriteMessage(websocket.TextMessage, serialisedGameState)
	if err != nil {
		log.Println("write:", err)
		panic(err)
	}
}

// RunGameLoop runs the core game loop
func RunGameLoop(config Config, addPlayerC chan Player) (chan bool, chan physics.Direction) {
	// Initialise Game Loop
	timePerTick := time.Duration(1000/config.TicksPerSecond) * time.Millisecond
	ticker := time.NewTicker(timePerTick)
	done := make(chan bool)
	updates := make(chan physics.Direction)

	// Wait for first (and currently only) player to be added
	go func() {
		fmt.Println("Waiting for player...")
		player := <-addPlayerC
		fmt.Println("Player joined!")

		// Initialise game state
		player.X = 0
		player.Y = 0
		gameState := State{
			Players: []Player{player},
		}

		// Start game loop in separate routine
		go func() {
			for {
				select {
				case <-done:
					{
						// When the game ends
						ticker.Stop()
					}
				case update := <-updates:
					{
						fmt.Println("Handling update in direction: ", update)
						// Only one player for now, so only update that player
						gameState.Players[0].Direction = update
					}
				case t := <-ticker.C:
					{
						fmt.Println("Tick at: ", t)
						// Recalc game state
						gameState.Players[0].Position = physics.CalculateNextPosition(gameState.Players[0].Position, &gameState.Players[0].Direction)
						// TODO: Push out data to all users
						// ... for now one user
						go SendUpdateToClient(gameState.Players[0].Conn, gameState)
					}
				}
			}
		}()
	}()
	return done, updates
}
