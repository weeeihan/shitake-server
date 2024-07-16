package ws

func NewHub() *Hub {
	return &Hub{
		Rooms:     make(map[string]*Room),
		Broadcast: make(chan *Message, 10),
	}
}

func (h *Hub) Run() {

	for m := range h.Broadcast {

		if _, ok := h.Rooms[m.RoomID]; ok {
			for _, player := range h.Rooms[m.RoomID].Players {
				player.Message <- m
			}
		}

	}
}
