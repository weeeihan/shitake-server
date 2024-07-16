package ws

import (
	"encoding/json"
	"log"
)

// Write Message
func (p *Player) writeMessages() {
	defer func() {
		log.Println("CLOSE WRITEMSG")

	}()

	for {
		msg, ok := <-p.Message

		if !ok {
			return
		}
		log.Println("WRITEMSG")
		p.Conn.WriteJSON(msg)
	}
}

// Read Message
func (p *Player) readMessages(hub *Hub) {
	defer func() {
		p.Conn.Close()
		close(p.Message)
		log.Println("CLOSE READMESG")
	}()

	for {
		log.Printf("READMSG")
		_, m, err := p.Conn.ReadMessage()
		if err != nil {
			return
		}

		// Parse message

		var msgReq *MessageReq
		if err := json.Unmarshal(m, &msgReq); err != nil {
			log.Println("ERROR2")
			log.Printf("error: %v", err)
		}
		room := hub.Rooms[p.RoomID]
		room.gameState(msgReq, p, hub)

	}
}
