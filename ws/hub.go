package ws

import (
	"log"
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
			h.Broadcast <- createMsg(player.RoomID, REGISTERED, "Someone Registered")

			// Check if everyone is connected

		case player := <-h.Unregister:
			player.Ready = false
			log.Printf("Player %v disconnected!", player.Name)

		case m := <-h.Broadcast:

			if _, ok := h.Rooms[m.RoomID]; ok {
				for _, player := range h.Rooms[m.RoomID].Players {
					player.Message <- m
				}
			}
		}
	}
}
