package websocket

import (
	"math/rand/v2"

	jsoniter "github.com/json-iterator/go"
)

func NextPlayer(room *Room, currentPlayer string) {
	playerIDs := make([]string, 0, len(room.Players))
	for id := range room.Players {
		playerIDs = append(playerIDs, id)
	}

	if len(playerIDs) == 0 {
		return
	}
	turnMessage := Message{
		Type:    "system",
		Payload: "現在輪到你猜",
	}
	turnData, _ := jsoniter.Marshal(turnMessage)
	if currentPlayer == "" {
		room.CurrentTurnID = playerIDs[rand.IntN(len(playerIDs))]

		for _, p := range room.Players {
			if p.ID == room.CurrentTurnID {
				p.Send <- turnData
			}
		}
		return
	}

	for i, id := range playerIDs {
		if id == currentPlayer {
			next := playerIDs[(i+1)%len(playerIDs)]
			room.CurrentTurnID = next
			for _, p := range room.Players {
				if p.ID == room.CurrentTurnID {
					p.Send <- turnData
				}
			}
			return
		}
	}
}
