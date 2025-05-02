package websocket

import (
	"backend/helpers"
	"log"
	"sync"
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

	// 如果佇列中有等待玩家，就配對
	if len(h.MatchQueue) > 0 {
		opponent := h.MatchQueue[0]
		h.MatchQueue = h.MatchQueue[1:]

		roomID := helpers.GenerateRandomPassword(6)
		log.Println(roomID)
		room := &Room{
			ID:      roomID,
			Players: make(map[string]*Player),
		}
		room.Players[player.ID] = player
		log.Println(player.ID)
		room.Players[opponent.ID] = opponent
		log.Println(opponent.ID)
		GameHub.Rooms[roomID] = room

		player.RoomID = roomID
		opponent.RoomID = roomID

		roomMessage := `{"type":"roomJoined", "data":{"roomId":"` + roomID + `"}}`
		player.Send <- []byte(roomMessage)
		opponent.Send <- []byte(roomMessage)
	} else {
		// 沒有對手，加入配對佇列
		h.MatchQueue = append(h.MatchQueue, player)
	}
}
