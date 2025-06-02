package controllers

import (
	"backend/helpers"
	"backend/service"
	ws "backend/websocket"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func JoinRoomHandler(s service.GameService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// roomID := c.Query("roomId")
		token := c.Query("token")
		claim, msg := helpers.ValidateToken(token)
		if msg != "" {
			log.Println("Token validation failed:", msg)
			c.Writer.WriteHeader(http.StatusUnauthorized)
			c.Writer.Write([]byte("Unauthorized"))
			return
		}
		playerID := claim.Uid
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("WebSocket upgrade failed:", err)
			return
		}

		player := &ws.Player{
			ID:          playerID,
			Conn:        conn,
			Send:        make(chan []byte, 256),
			GameService: s,
		}
		log.Println("WebSocket connection established for player:", playerID)
		go player.ReadMessages()
		go player.WriteMessages()
		ws.GameHub.MatchPlayer(player)
	}
}
