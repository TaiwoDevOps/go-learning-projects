package main

import (
	"log"

	"github.com/TaiwoDevOps/chat-app/db"
	"github.com/TaiwoDevOps/chat-app/internal/user"
	"github.com/TaiwoDevOps/chat-app/internal/ws"
	"github.com/TaiwoDevOps/chat-app/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not connect to DB %v", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())

	userSrv := user.NewService(userRep)
	userHandler := user.NewHandler(userSrv)

	// websocket handler
	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	// create a separate go routine to run the hub
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)

	router.Start(":8080")
}
