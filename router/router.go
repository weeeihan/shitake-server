package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shitake/ws"
)

var r *gin.Engine

func InitRouter(wsHandler *ws.Handler) {
	r = gin.Default()

	// r.POST("/signup", wsHandler.CreateUser)
	// r.POST("/newRoom", wsHandler.CreateRoom)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:5173"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/ws/createRoom", wsHandler.CreateRoom)
	r.GET("/ws/joinRoom/:roomID", wsHandler.JoinRoom)
	r.GET("/ws/leaveRoom/:playerID", wsHandler.LeaveRoom)
	r.GET("/ws/checkPlayer/:playerID", wsHandler.CheckPlayer)
	r.GET("/ws/connectToGame/:playerID", wsHandler.ConnectToGame)
	r.GET("/ws/getPlayer/:playerID", wsHandler.GetPlayer)

	r.GET("/getStates", wsHandler.GetStates)
	r.GET("/debug", wsHandler.Debug)

	// r.GET("/ws/getGames", wsHandler.GetGames)
	// r.GET("/ws/getGame/:roomID", wsHandler.GetGame)
	r.GET("/ws/getRooms", wsHandler.GetRooms)
	r.GET("/ws/getRoom/:roomID", wsHandler.GetRoom)
	// r.GET("/ws/getRooms", wsHandler.GetRooms)
	// r.GET("/ws/getRoomInt/:roomId", wsHandler.GetRoomInternal)
	// r.GET("/ws/getRoom/:roomId", wsHandler.GetRoom)
	// r.GET("/ws/getClients/:roomId", wsHandler.GetClients)

}

func Start(addr string) error {
	return r.Run(addr)
}
