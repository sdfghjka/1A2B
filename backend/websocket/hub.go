package websocket

type Hub struct {
	Rooms map[string]*Room
	// 其他未來要加：配對佇列、廣播通道等
}

var GameHub = &Hub{
	Rooms: make(map[string]*Room),
}

func (h *Hub) JoinRoom(roomID string, player *Player) {
	room, ok := h.Rooms[roomID]
	if !ok {
		room = &Room{ID: roomID, Players: make(map[string]*Player)}
		h.Rooms[roomID] = room
	}
	room.Players[player.ID] = player
}
