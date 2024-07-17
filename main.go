package main

import (

	// "github.com/shitake/router"
	// "github.com/shitake/ws"

	"github.com/shitake/router"
	"github.com/shitake/ws"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {
	// envErr := godotenv.Load(".env")
	// if envErr != nil {
	// 	log.Fatal("Error loading .env file")

	// }
	// url := os.Getenv("TURSO_DATABASE_URL")
	// db, err := sql.Open("libsql", url)
	// if err != nil {
	// 	log.Fatal(os.Stderr, "Failed to open db")
	// 	os.Exit(1)
	// }

	// defer db.Close()

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)

	router.InitRouter(wsHandler)

	router.Start("0.0.0.0:8080")

}
