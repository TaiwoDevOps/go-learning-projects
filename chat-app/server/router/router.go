package router

import (
	"time"

	"github.com/TaiwoDevOps/chat-app/internal/user"
	"github.com/TaiwoDevOps/chat-app/internal/ws"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
	r = gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Allow all origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000" || origin == "http://localhost:8080" // Allow only this origin
		},
		MaxAge: 12 * time.Hour, // 12 hours
	}))

	// Configure trusted proxies
	r.SetTrustedProxies(nil)

	r.POST("/sign-up", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)

	r.POST("/ws/create-room", wsHandler.CreateRoom)
	r.GET("/ws/join-room/:roomId", wsHandler.JoinRoom)
	r.GET("/ws/rooms", wsHandler.GetRooms)
	r.GET("/ws/clients/:roomId", wsHandler.GetClients)
}

func Start(address string) error {

	return r.Run(address)
}
