package main

import (
	"github.com/shitake/router"
	"github.com/shitake/ws"
)

func main() {
	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(wsHandler)
	router.Start("0.0.0.0:8080")

}
