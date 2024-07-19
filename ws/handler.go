package ws

import (
	"log"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

func NewHandler(h *Hub) *Handler {

	// FOR TESTING PURPOSE

	return &Handler{
		hub: h,
	}
}

func (h *Handler) GetConstants(c *gin.Context) {
	c.JSON(http.StatusOK, GameConstant{
		States: gamestates,
	})
}

func (h *Handler) GetStates(c *gin.Context) {
	c.JSON(http.StatusOK, gamestates)
}

func (h *Handler) CreateRoom(c *gin.Context) {

	// Generate room code
	id := getRandID(h.hub.Rooms)
	// id := "1111"

	playerID := getPlayerID(id)

	player := newPlayer(playerID, c.Query("name"), id)
	// Register player into the room
	// player := &Player{
	// 	RoomID:  id,
	// 	ID:      playerID,
	// 	Play:    -1,
	// 	Name:    c.Query("name"),
	// 	Hand:    []int{},
	// 	HP:      100,
	// 	Ready:   false,
	// 	Message: make(chan *Message, 10),
	// 	End:     0,
	// 	DamageReport: &DamageReport{
	// 		Mushrooms:      0,
	// 		Damage:         0,
	// 		RoundMushrooms: 0,
	// 		RoundDamage:    0,
	// 		MushroomTypes:  []int{},
	// 	},
	// }

	players := make(map[string]*Player)
	players[player.ID] = player

	// Create room
	h.hub.Rooms[id] = &Room{
		ID:        id,
		Deck:      [][]int{},
		Players:   players,
		State:     INIT,
		Played:    make(map[int]string),
		Hands:     make(map[string][]int),
		HPs:       make(map[string]int),
		Select:    make(map[string]int),
		Pause:     false,
		Ready:     0,
		Chooser:   "",
		Moves:     [][]string{},
		Mushrooms: getMushrooms(),
		Online:    make(map[string]bool),
		Stopper:   make(chan bool),
	}

	c.JSON(http.StatusOK, newPlayerRes(player))

}

func (h *Handler) JoinRoom(c *gin.Context) {
	roomID := c.Param("roomID")
	name := c.Query("name")

	// check if room exists
	if _, ok := h.hub.Rooms[roomID]; !ok {
		log.Printf("Room does not exist!")
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Room does not exist!"})
		return
	}

	room := h.hub.Rooms[roomID]

	// if room is not in lobby state
	if room.State != INIT {
		log.Printf("Room is not INIT")
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Game has already started"})
		return
	}

	// check if name is already used
	for _, p := range h.hub.Rooms[roomID].Players {
		if name == p.Name {
			log.Printf("Name used")
			c.JSON(http.StatusBadRequest, gin.H{"Message": "Name is used already, please choose another name!"})
			return
		}
	}

	// generate random ID
	playerID := getPlayerID(roomID)

	// c.JSON(http.StatusOK, gin.H{"PlayerID": playerID})
	r := h.hub.Rooms[roomID]

	// Create player
	player := newPlayer(playerID, name, roomID)
	// player := &Player{
	// 	RoomID:  roomID,
	// 	ID:      playerID,
	// 	Name:    name,
	// 	Hand:    []int{},
	// 	Play:    -1,
	// 	HP:      100,
	// 	Ready:   false,
	// 	End:     0,
	// 	Message: make(chan *Message, 10),
	// 	DamageReport: &DamageReport{
	// 		Mushrooms:      0,
	// 		Damage:         0,
	// 		RoundMushrooms: 0,
	// 		RoundDamage:    0,
	// 		MushroomTypes:  []int{},
	// 	},
	// }

	// Register player into the room
	r.Players[player.ID] = player

	// playerRes := &PlayerRes{
	// 	ID:           player.ID,
	// 	Name:         player.Name,
	// 	Hand:         player.Hand,
	// 	Play:         player.Play,
	// 	HP:           player.HP,
	// 	Ready:        player.Ready,
	// 	End:          player.End,
	// 	DamageReport: *player.DamageReport,
	// }

	c.JSON(http.StatusOK, newPlayerRes(player))

}

// func (h *Handler) CheckPlayer(c *gin.Context) {
// 	playerID := c.Param("playerID")
// 	roomID := playerID[len(playerID)-4:]

// 	if _, ok := h.hub.Rooms[roomID]; !ok {
// 		log.Println("ROOMS dont exist")
// 		c.JSON(http.StatusBadRequest, false)
// 		return
// 	}

// 	r := h.hub.Rooms[roomID]
// 	if _, ok := r.Players[playerID]; !ok {
// 		log.Println("PLAYES DONT EXIST")
// 		c.JSON(http.StatusBadRequest, false)
// 		return
// 	}

// 	c.JSON(http.StatusOK, true)
// 	// go player.writeMessages()
// 	// player.readMessages(h.hub)

// }

func (h *Handler) ConnectToGame(c *gin.Context) {

	// log.Printf("CONNECTING TO GAME!")
	playerID := c.Param("playerID")
	roomID := playerID[len(playerID)-4:]
	if _, ok := h.hub.Rooms[roomID]; !ok {
		log.Printf("Room does not exist!")
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Room does not exist!"})
		return
	}
	room := h.hub.Rooms[roomID]
	room.Online[playerID] = true
	player := h.hub.Rooms[roomID].Players[playerID]

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, p := range room.Players {
		// Notify others that you just joined.
		if p.ID == playerID || p.IsBot {
			continue
		}
		p.Message <- &Message{RoomID: roomID, State: NEW_PLAYER_JOINED, Remark: "Player joined!"}

	}

	player.Message = make(chan *Message, 10)
	player.Conn = conn
	room.Online[playerID] = true
	go player.writeMessages(h.hub)
	player.readMessages(h.hub)

}

func (h *Handler) GetRooms(c *gin.Context) {
	rooms := make([]RoomRes, 0)
	for _, r := range h.hub.Rooms {
		rooms = append(rooms, RoomRes{
			ID:      r.ID,
			State:   r.State,
			Players: playersArr(r.Players),
			Online:  r.Online,
		})
	}
	c.JSON(http.StatusOK, rooms)
}

func (h *Handler) GetRoom(c *gin.Context) {
	roomId := c.Param("roomID")

	if _, ok := h.hub.Rooms[roomId]; !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "Room does not exist"})
		return
	}

	r := h.hub.Rooms[roomId]
	c.JSON(http.StatusOK, RoomRes{
		ID:      r.ID,
		State:   r.State,
		Deck:    r.Deck,
		Players: playersArr(r.Players),
		Played:  getPlayed(r),
		Chooser: getChooser(r),
		Moves:   r.Moves,
	})
}

// func (h *Handler) GetGames(c *gin.Context) {
// 	games := make([]GameRes, 0)
// 	for _, g := range h.hub.Games {
// 		games = append(games, GameRes{
// 			ID:      g.ID,
// 			Deck:    g.Deck,
// 			State:   g.State,
// 			Players: playersArr(g.Players),
// 		})
// 	}
// 	c.JSON(http.StatusOK, games)
// }

// For game info specifically
// func (h *Handler) GetGame(c *gin.Context) {
// 	roomID := c.Param("roomID")

// 	if _, ok := h.hub.Games[gameID]; !ok {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "Game does not exist"})
// 		return
// 	}

// 	g := h.hub.Games[gameID]
// 	c.JSON(http.StatusOK, GameRes{
// 		ID:      g.ID,
// 		Deck:    g.Deck,
// 		State:   g.State,
// 		Players: playersArr(g.Players),
// 	})
// }

// LeaveRoom removes the player from the room
func (h *Handler) LeaveRoom(c *gin.Context) {
	playerID := c.Param("playerID")
	roomID := playerID[len(playerID)-4:]

	if _, ok := h.hub.Rooms[roomID]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Room does not exist"})
		return
	}

	r := h.hub.Rooms[roomID]
	player := r.Players[playerID]
	player.Conn.Close()
	delete(r.Players, playerID)

	if len(r.Players) == 0 {
		delete(h.hub.Rooms, roomID)
		c.JSON(http.StatusOK, gin.H{"Message": "Clear room"})
	} else {
		r.broadcast(createMsg(r.ID, PLAYER_LEFT, "Player left room"))
		c.JSON(http.StatusOK, gin.H{"Message": "Player left room"})
	}

}

func (h *Handler) GetData(c *gin.Context) {
	playerID := c.Param("playerID")
	roomID := playerID[len(playerID)-4:]

	if _, ok := h.hub.Rooms[roomID]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Room does not exist"})
		return
	}

	r := h.hub.Rooms[roomID]

	if _, ok := r.Players[playerID]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Player does not exist"})
		return
	}

	player := PlayerRes{
		ID:           playerID,
		Name:         r.Players[playerID].Name,
		Hand:         r.Players[playerID].Hand,
		Play:         r.Players[playerID].Play,
		HP:           r.Players[playerID].HP,
		Ready:        r.Players[playerID].Ready,
		End:          r.Players[playerID].End,
		DamageReport: *r.Players[playerID].DamageReport,
		IsBot:        r.Players[playerID].IsBot,
	}

	room := RoomRes{
		ID:        r.ID,
		State:     r.State,
		Deck:      r.Deck,
		Players:   playersArr(r.Players),
		Played:    getPlayed(r),
		Chooser:   getChooser(r),
		Moves:     r.Moves,
		Mushrooms: r.Mushrooms,
	}

	c.JSON(http.StatusOK, GameData{
		Room:   room,
		Player: player,
	})
}

func (h *Handler) GetPlayer(c *gin.Context) {
	playerID := c.Param("playerID")
	roomID := playerID[len(playerID)-4:]

	if _, ok := h.hub.Rooms[roomID]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Room does not exist"})
		return
	}

	r := h.hub.Rooms[roomID]
	if _, ok := r.Players[playerID]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Player does not exist"})
		return
	}

	player := &PlayerRes{
		ID:           playerID,
		Name:         r.Players[playerID].Name,
		Hand:         r.Players[playerID].Hand,
		Play:         r.Players[playerID].Play,
		HP:           r.Players[playerID].HP,
		Ready:        r.Players[playerID].Ready,
		End:          r.Players[playerID].End,
		DamageReport: *r.Players[playerID].DamageReport,
	}

	c.JSON(http.StatusOK, player)
}

func (h *Handler) Debug(c *gin.Context) {
	log.Print(runtime.NumGoroutine())
	// log.Printf()
	// for _, r := range h.hub.Rooms {
	// 	for _, p := range r.Players {
	// 		log.Print(p)
	// 	}
	// }
	c.JSON(http.StatusOK, gin.H{"Message": "Debugging"})
}
