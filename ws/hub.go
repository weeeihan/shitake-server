package ws

import (
	"log"
	"time"
)

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
			if _, ok := h.Rooms[player.RoomID]; ok {
				room := h.Rooms[player.RoomID]
				room.Online = room.Online + 1
			}

			// h.Broadcast <- createMsg(player.RoomID, REGISTERED, player.Name)

		// Check if everyone is connected

		case player := <-h.Unregister:

			if _, ok := h.Rooms[player.RoomID]; ok {
				room := h.Rooms[player.RoomID]
				room.Online = room.Online - 1

				if room.Online == 0 {
					// Start timer.
					room.Ticker = time.NewTicker(1 * time.Second)
					go room.timer(room.ID, 20, IDLE, player, h)
				}
				if room.State != CHOOSE_CARD {
					log.Printf("Falsify ready")
					player.Ready = false
				}
			}
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
