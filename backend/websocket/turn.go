package websocket

import (
	"math/rand/v2"
)

func NextPlayer(room *Room, currentPlayer string) {
	playerIDs := make([]string, 0, len(room.Players))
	for id := range room.Players {
		playerIDs = append(playerIDs, id)
	}

	if len(playerIDs) == 0 {
		return
	}

	if currentPlayer == "" {
		room.CurrentTurnID = playerIDs[rand.IntN(len(playerIDs))]
		return
	}

	for i, id := range playerIDs {
		if id == currentPlayer {
			next := playerIDs[(i+1)%len(playerIDs)]
			room.CurrentTurnID = next
			return
		}
	}
}
