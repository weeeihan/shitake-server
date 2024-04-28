package ws

import (
	"encoding/json"
	"log"
)

// Write Message
func (p *Player) writeMessages() {
	defer func() {
		p.Conn.Close()
	}()

	for {
		msg, ok := <-p.Message
		if msg.State == DISCONNECTED {
			log.Println("DISCONNECTED!")
			return
		}
		if !ok {
			log.Println("CLOSED GOROUTINE!")
			return
		}
		p.Conn.WriteJSON(msg)
	}
}

// Read Message
func (p *Player) readMessages(hub *Hub) {
	defer func() {
		hub.Unregister <- p
		p.Conn.Close()
	}()

	for {
		_, m, err := p.Conn.ReadMessage()
		if err != nil {
			log.Printf("error: %v", err)

			break
		}
		var msgReq *MessageReq
		if err := json.Unmarshal(m, &msgReq); err != nil {
			log.Printf("error: %v", err)
		}
		room := hub.Rooms[p.RoomID]
		log.Print(p.RoomID)
		log.Print(room)

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
