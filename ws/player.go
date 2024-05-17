package ws

import (
	"encoding/json"
	"log"
	"runtime"
)

// Write Message
func (p *Player) writeMessages() {
	defer func() {
		log.Printf("Goroutines: %v ", runtime.NumGoroutine())

		log.Println("CLOSE WRITEMSG")
		p.Conn.Close()
	}()

	for {
		msg, ok := <-p.Message
		if msg.State == DISCONNECTED {
			log.Println("DISCONNECTED!")
			return
		}
		if !ok {
			log.Println("NOT OK")
			log.Println("CLOSED GOROUTINE!")
			return
		}
		// log.Println("WRITING STUFF")
		// log.Println(msg.Remark)
		p.Conn.WriteJSON(msg)
	}
}

// Read Message
func (p *Player) readMessages(hub *Hub) {
	defer func() {
		// hub.Unregister <- p
		p.Conn.Close()
		log.Println("CLOSE READMESG")
	}()

	for {
		_, m, err := p.Conn.ReadMessage()
		if err != nil {
			p.Message <- createMsg(p.RoomID, DISCONNECTED, "Someone Disconnected")
			log.Println("ERROR1")

			break
		}
		var msgReq *MessageReq
		if err := json.Unmarshal(m, &msgReq); err != nil {
			log.Println("ERROR2")
			log.Printf("error: %v", err)
		}
		room := hub.Rooms[p.RoomID]
		log.Print(p.RoomID)
		log.Print(room)

		log.Printf("READING MESSAGE: %v", msgReq.Action)

		room.gameState(msgReq, p, hub)

	}
}

// // Read Remark
// func (c *Client) readRemark(hub *Hub) {
// 	defer func() {
// 		hub.Unregister <- c
// 		c.Conn.Close()
// 	}()

// 	for {
// 		_, m, err := c.Conn.ReadRemark()
// 		if err != nil {
// 			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
// 				log.Printf("error: %v", err)
// 			}
// 			break
// 		}

// 		remark := &Remark{
// 			Content:  string(m),
// 			RoomID:   c.RoomID,
// 			Username: c.Username,
// 		}

// 		hub.Broadcast <- remark
// 	}
// }
