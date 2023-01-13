package main

// Welcome to
// __________         __    __  .__                               __
// \______   \_____ _/  |__/  |_|  |   ____   ______ ____ _____  |  | __ ____
//  |    |  _/\__  \\   __\   __\  | _/ __ \ /  ___//    \\__  \ |  |/ // __ \
//  |    |   \ / __ \|  |  |  | |  |_\  ___/ \___ \|   |  \/ __ \|    <\  ___/
//  |________/(______/__|  |__| |____/\_____>______>___|__(______/__|__\\_____>
//
// This file can be a nice home for your Battlesnake logic and helper functions.
//
// To get you started we've included code to prevent your Battlesnake from moving backwards.
// For more info see docs.battlesnake.com

import (
	"log"
	"math/rand"
	"strconv"
)

// info is called when you create your Battlesnake on play.battlesnake.com
// and controls your Battlesnake's appearance
// TIP: If you open your Battlesnake URL in a browser you should see this data
func info() BattlesnakeInfoResponse {
	log.Println("INFO")

	return BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "RandlSnek", // TODO: Your Battlesnake username
		Color:      "##ff6600",  // TODO: Choose color
		Head:       "evil",      // TODO: Choose head
		Tail:       "hook",      // TODO: Choose tail
	}
}

// start is called when your Battlesnake begins a game
func start(state GameState) {
	log.Println("GAME START")
}

// end is called when your Battlesnake finishes a game
func end(state GameState) {
	log.Printf("GAME OVER\n\n")
}

// move is called on every turn and returns your next move
// Valid moves are "up", "down", "left", or "right"
// See https://docs.battlesnake.com/api/example-move for available data
func move(state GameState) BattlesnakeMoveResponse {
	log.Print("state.Turn: ", state.Turn)

	isMoveSafe := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}

	// We've included code to prevent your Battlesnake from moving backwards
	myHead := state.You.Body[0] // Coordinates of your head
	myNeck := state.You.Body[1] // Coordinates of your "neck"

	if myNeck.X < myHead.X { // Neck is left of head, don't move left
		isMoveSafe["left"] = false
		log.Print("cant move left, would hit neck")

	} else if myNeck.X > myHead.X { // Neck is right of head, don't move right
		isMoveSafe["right"] = false
		log.Print("cant move right, would hit neck")

	} else if myNeck.Y < myHead.Y { // Neck is below head, don't move down
		isMoveSafe["down"] = false
		log.Print("cant move down, would hit neck")

	} else if myNeck.Y > myHead.Y { // Neck is above head, don't move up
		isMoveSafe["up"] = false
		log.Print("cant move up, would hit neck")

	}

	// Prevent Battlesnake from moving out of bounds
	boardWidth := state.Board.Width
	boardHeight := state.Board.Height
	//log.Print("Board width: ", boardWidth)
	//log.Print("Board height: ", boardHeight)

	log.Print("My head x: ", myHead.X)
	log.Print("My head y: ", myHead.Y)
	if myHead.Y+1 >= boardHeight {
		isMoveSafe["up"] = false
		log.Print("cant move up, oob")
	}
	if myHead.Y-1 < 0 {
		isMoveSafe["down"] = false
		log.Print("cant move down, oob")
	}
	if myHead.X+1 >= boardWidth {
		isMoveSafe["right"] = false
		log.Print("cant move right, oob")
	}
	if myHead.X-1 < 0 {
		isMoveSafe["left"] = false
		log.Print("cant move left, oob")
	}

	// Prevent Battlesnake from colliding with itself
	myBody := state.You.Body
	bodyLenth := len(myBody)
	log.Print("BodyLength: " + strconv.Itoa(bodyLenth))
	for i := 1; i < bodyLenth; i++ {
		if myHead.X+1 == myBody[i].X {
			if myHead.Y == myBody[i].Y {
				isMoveSafe["right"] = false
				log.Print("cant move right, would hit body")
			}
		}
		if myHead.X-1 == myBody[i].X {
			if myHead.Y == myBody[i].Y {
				isMoveSafe["left"] = false
				log.Print("cant move left, would hit body")
			}
		}
		if myHead.Y+1 == myBody[i].Y {
			if myHead.X == myBody[i].X {
				isMoveSafe["up"] = false
				log.Print("cant move up, would hit body")
			}
		}
		if myHead.Y-1 == myBody[i].Y {
			if myHead.X == myBody[i].X {
				isMoveSafe["down"] = false
				log.Print("cant move down, would hit body")
			}
		}
	}

	// TODO: Step 3 - Prevent your Battlesnake from colliding with other Battlesnakes
	// opponents := state.Board.Snakes

	// Are there any safe moves left?
	safeMoves := []string{}
	for move, isSafe := range isMoveSafe {
		if isSafe {
			safeMoves = append(safeMoves, move)
		}
	}

	if len(safeMoves) == 0 {
		log.Printf("MOVE %d: No safe moves detected! Moving down\n", state.Turn)
		return BattlesnakeMoveResponse{Move: "down"}
	}

	// Choose a random move from the safe ones
	nextMove := safeMoves[rand.Intn(len(safeMoves))]

	// TODO: Step 4 - Move towards food instead of random, to regain health and survive longer
	// food := state.Board.Food

	log.Printf("Possbile moves: %s", safeMoves)
	log.Printf("Moving %d: %s\n", state.Turn, nextMove)
	return BattlesnakeMoveResponse{Move: nextMove}
}

func main() {
	RunServer()
}
