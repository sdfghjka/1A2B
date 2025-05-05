package websocket

import (
	"backend/helpers"
	"log"
	"sync"

	jsoniter "github.com/json-iterator/go"
)

type Hub struct {
	Rooms      map[string]*Room
	MatchQueue []*Player
	Mutex      sync.Mutex
}

var GameHub = &Hub{
	Rooms:      make(map[string]*Room),
	MatchQueue: make([]*Player, 0),
}

func (h *Hub) MatchPlayer(player *Player) {
	h.Mutex.Lock()
	log.Printf("%s Waiting....", player.ID)
	defer h.Mutex.Unlock()

	if len(h.MatchQueue) > 0 {
		opponent := h.MatchQueue[0]
		h.MatchQueue = h.MatchQueue[1:]

		roomID := helpers.GenerateRandomPassword(6)
		log.Println(roomID)
		room := &Room{
			ID:      roomID,
			Players: make(map[string]*Player),
			Answer:  helpers.GenerateAnswer(),
		}
		room.Players[player.ID] = player
		log.Println(player.ID)
		room.Players[opponent.ID] = opponent
		log.Println(opponent.ID)
		GameHub.Rooms[roomID] = room
		player.RoomID = roomID
		opponent.RoomID = roomID
		room.CurrentTurnID = NextPlayer(room.Players, "")
		roomMsg := Message{
			Type: "roomJoined",
			Payload: map[string]string{
				"roomId": roomID,
			},
			From: player.ID,
		}
		roomJSON, err := jsoniter.Marshal(roomMsg)
		if err != nil {
			log.Println("Failed to marshal roomJoined:", err)
			return
		}

		log.Println(string(roomJSON))
		player.Send <- []byte(roomJSON)
		opponent.Send <- []byte(roomJSON)
	} else {
		h.MatchQueue = append(h.MatchQueue, player)
	}
}
