package ws

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHandler(h *Hub) *Handler {

	// FOR TESTING PURPOSE

	return &Handler{
		hub: h,
	}
}

func (h *Handler) CreateRoom(c *gin.Context) {

	// Generate room code
	id := getRandID(h.hub.Rooms)

	// Register player into the room
	player := &Player{
		ID:     getPlayerID(id),
		Name:   c.Query("name"),
		RoomID: id,
	}
	players := make(map[string]*Player)
	players[player.ID] = player

	// Create room
	h.hub.Rooms[id] = &Room{
		ID:      id,
		Players: players,
		State:   LOBBY,
	}

	// Initialize game
	h.hub.Games[id] = &Game{
		ID:      id,
		Deck:    [][]int{},
		Players: make(map[string]*Player),
		State:   INIT,
		Played:  make(map[int]string),
		Hands:   make(map[string][]int),
		Scores:  make(map[string]int),
		Select:  make(map[string]int),
		Pause:   false,
		Ready:   0,
	}

	c.JSON(http.StatusOK, gin.H{"playerID": player.ID, "roomID": id})
}

func (h *Handler) JoinRoom(c *gin.Context) {

	roomID := c.Param("roomID")
	name := c.Query("name")

	// check if room exists
	if _, ok := h.hub.Rooms[roomID]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Room does not exist!"})
		return
	}

	room := h.hub.Rooms[roomID]

	// if room is not in lobby state
	if room.State != LOBBY {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Game has already started"})
		return
	}

	// check if name is already used
	for _, p := range h.hub.Rooms[roomID].Players {
		if name == p.Name {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "Name is used already, please choose another name!"})
			return
		}
	}

	// generate random ID
	playerID := getPlayerID(roomID)

	r := h.hub.Rooms[roomID]

	// Create player
	player := &Player{
		RoomID: roomID,
		ID:     playerID,
		Name:   name,
		Hand:   []int{},
		Score:  0,
	}

	// Register player into the room
	r.Players[player.ID] = player

	log.Print(player)

	c.JSON(http.StatusOK, gin.H{"PlayerID": player.ID})
}

func (h *Handler) ConnectToGame(c *gin.Context) {
	gameID := c.Param("gameID")
	playerID := c.Query("playerID")

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p := h.hub.Rooms[gameID].Players[playerID]
	room := h.hub.Rooms[gameID]
	game := h.hub.Games[gameID]

	_, exists := room.Players[playerID]

	var player *Player

	if !exists {
		log.Print("NOT IN GAME")
		player = &Player{
			ID:      p.ID,
			RoomID:  gameID,
			Name:    p.Name,
			Conn:    conn,
			Message: make(chan *Message, 10),
			Hand:    []int{},
			Score:   0,
			Play:    0,
			Ready:   false,
		}
	} else {
		room.Players[playerID].Ready = false
		player = room.Players[playerID]
		player.Conn = conn
		player.Message = make(chan *Message, 10)
	}
	game.Players[player.ID] = player

	h.hub.Register <- player
	h.hub.Broadcast <- &Message{
		GameID: gameID,
		State:  NEW_PLAYER_JOINED,
		Remark: "Player joined!",
	}

	go player.writeMessages()
	player.readMessages(h.hub)
}

func (h *Handler) GetRooms(c *gin.Context) {
	rooms := make([]RoomRes, 0)
	for _, r := range h.hub.Rooms {
		rooms = append(rooms, RoomRes{
			ID:      r.ID,
			State:   r.State,
			Players: playersArr(r.Players),
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
		Players: playersArr(r.Players),
	})
}

func (h *Handler) GetGames(c *gin.Context) {
	games := make([]GameRes, 0)
	for _, g := range h.hub.Games {
		games = append(games, GameRes{
			ID:      g.ID,
			Deck:    g.Deck,
			State:   g.State,
			Players: playersArr(g.Players),
		})
	}
	c.JSON(http.StatusOK, games)
}

func (h *Handler) GetGame(c *gin.Context) {
	gameID := c.Param("gameID")

	if _, ok := h.hub.Games[gameID]; !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "Game does not exist"})
		return
	}

	g := h.hub.Games[gameID]
	c.JSON(http.StatusOK, GameRes{
		ID:      g.ID,
		Deck:    g.Deck,
		State:   g.State,
		Players: playersArr(g.Players),
	})
}

// LeaveRoom removes the player from the room
func (h *Handler) LeaveRoom(c *gin.Context) {
	playerID := c.Param("playerID")
	roomID := playerID[len(playerID)-4:]

	if _, ok := h.hub.Rooms[roomID]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Room does not exist"})
		return
	}

	r := h.hub.Rooms[roomID]
	g := h.hub.Games[roomID]
	delete(r.Players, playerID)
	delete(g.Players, playerID)

	if len(r.Players) == 0 {
		delete(h.hub.Rooms, roomID)
		c.JSON(http.StatusOK, gin.H{"Message": "Clear room"})
	} else {
		h.hub.Broadcast <- &Message{
			GameID: roomID,
			State:  PLAYER_LEFT,
			Remark: "Player left!",
		}
		c.JSON(http.StatusOK, gin.H{"Message": "Player left room"})
	}

}

func (h *Handler) GetPlayer(c *gin.Context) {
	playerID := c.Param("playerID")
	roomID := c.Param("roomID")

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
		ID:    playerID,
		Name:  r.Players[playerID].Name,
		Hand:  r.Players[playerID].Hand,
		Score: r.Players[playerID].Score,
		Ready: r.Players[playerID].Ready,
	}

	c.JSON(http.StatusOK, player)
}
