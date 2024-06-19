package ws

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Player),
		Unregister: make(chan *Player),
		Broadcast:  make(chan *Message, 10),
	}
}

func (h *Hub) Run() {

	for {
		select {
		// If join rooms
		case player := <-h.Register:
			// log.Print("Connected")
			h.Broadcast <- createMsg(player.RoomID, REGISTERED, player.Name)

		// Check if everyone is connected

		case player := <-h.Unregister:
			player.Ready = false
			// player.Conn = nil
			// player.Message <- createMsg(player.RoomID, DISCONNECTED, "Someone Disconnected")
			// log.Print("Player stuff disconnected!")
			// log.Print(player)
			// h.Rooms[player.RoomID].CheckConn()

		case m := <-h.Broadcast:

			if _, ok := h.Rooms[m.RoomID]; ok {
				for _, player := range h.Rooms[m.RoomID].Players {
					player.Message <- m
				}
			}
		}
	}
}
