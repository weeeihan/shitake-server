package ws

import (
	"log"
)

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Games:      make(map[string]*Game),
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
			log.Print(player)

			// Check if everyone is connected

		case player := <-h.Unregister:
			log.Printf("Player %v disconnected!", player.Name)

		case m := <-h.Broadcast:

			if _, ok := h.Games[m.GameID]; ok {
				for _, player := range h.Games[m.GameID].Players {
					player.Message <- m
				}
			}
		}
	}
}
