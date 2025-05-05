package websocket

import (
	"math/rand/v2"
)

func NextPlayer(players map[string]*Player, currentPlayer string) string {
	var playerIDs []string
	for id := range players {
		playerIDs = append(playerIDs, id)
	}

	if len(playerIDs) == 0 {
		return ""
	}

	if currentPlayer == "" {
		return playerIDs[rand.IntN(len(playerIDs))]
	}

	for i, id := range playerIDs {
		if id == currentPlayer {
			return playerIDs[(i+1)%len(playerIDs)]
		}
	}

	return playerIDs[rand.IntN(len(playerIDs))]
}
