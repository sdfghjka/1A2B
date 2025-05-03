package websocket

type Room struct {
	ID      string
	Players map[string]*Player
	Answer  string
}
