package ws

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

// Write Message
func (p *Player) writeMessages() {
	defer func() {
		p.Conn.Close()
	}()

	for {
		msg, ok := <-p.Message
		if !ok {
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
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		var msgReq *MessageReq
		if err := json.Unmarshal(m, &msgReq); err != nil {
			log.Printf("error: %v", err)
		}
		game := hub.Games[p.RoomID]
		log.Print(p.RoomID)
		log.Print(game)

		msg := gameState(msgReq, hub, game, p)

		hub.Broadcast <- msg
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
