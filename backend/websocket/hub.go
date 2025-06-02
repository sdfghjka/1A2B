package websocket

import (
	"backend/helpers"
	"fmt"
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
			ID:        roomID,
			Players:   make(map[string]*Player),
			Answer:    helpers.GenerateAnswer(),
			Player1ID: opponent.ID,
			Player2ID: player.ID,
		}
		room.Players[player.ID] = player
		room.Players[opponent.ID] = opponent
		GameHub.Rooms[roomID] = room
		player.RoomID = roomID
		opponent.RoomID = roomID
		NextPlayer(room, "")
		roomMsg := Message{
			Type: "roomJoined",
			Payload: map[string]string{
				"roomId": roomID,
			},
			From: player.ID,
		}
		log.Println(room.Answer)
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

func (h *Hub) RemoveFormRoom(p *Player) {
	h.Mutex.Lock()
	defer h.Mutex.Unlock()

	room, exist := h.Rooms[p.RoomID]
	if !exist {
		log.Printf("Room %s not found for player %s", p.RoomID, p.ID)
		return
	}
	delete(room.Players, p.ID)
	for _, other := range room.Players {
		msg := Message{
			Type:    "playerLeft",
			Payload: fmt.Sprintf("Player %s has left the room", p.ID),
			From:    "system",
		}
		if b, err := jsoniter.Marshal(msg); err == nil {
			other.Send <- b
		}
	}
	if len(room.Players) <= 2 {
		delete(h.Rooms, room.ID)
		log.Printf("Room %s deleted as it's now empty", room.ID)
	}
}
