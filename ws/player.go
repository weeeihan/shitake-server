package ws

import (
	"encoding/json"
	"log"
	"time"
)

// Write Message
func (p *Player) writeMessages(hub *Hub) {
	room := hub.Rooms[p.RoomID]
	defer func() {
		// log.Println("CLOSE WRITEMSG")
		delete(room.Online, p.ID)
		p.Ready = false

		if len(room.Online) == 0 {
			// Start counting.
			room.Idle = time.NewTicker(1 * time.Second)
			go func() {

				defer func() {
					room.Idle = nil
				}()
				i := 300
				for {
					select {
					case <-room.Idle.C:
						if i == 0 {
							// delete room!
							delete(hub.Rooms, room.ID)
							return
						}
						i--

					case <-room.Stopper:
						return
					}
				}
			}()

		}
		log.Println("Disconnected")
		// Check if there are no online players in the room.
	}()
	log.Println("Connected")
	if room.Idle != nil {
		room.Stopper <- true
	}
	for {
		msg, ok := <-p.Message

		if !ok {
			return
		}
		// log.Println("WRITEMSG")
		p.Conn.WriteJSON(msg)
	}
}

// Read Message
func (p *Player) readMessages(hub *Hub) {
	defer func() {
		log.Print("closing messages")
		log.Print(p.Message)
		// p.Message = nil
		close(p.Message)
		p.Conn.Close()
		// log.Println("CLOSE READMESG")
	}()

	for {
		// log.Printf("READMSG")
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
