package socketserver

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/redis/go-redis/v9"
	"jubeliotesting/pkg/config"
	"jubeliotesting/pkg/middleware"
	"log"
	"time"
)

type WebsocketServer struct {
	Config config.GetEnvConfig
	Client *redis.Client
}

func NewWebsocketServer(config config.GetEnvConfig, client *redis.Client) *WebsocketServer {
	return &WebsocketServer{Config: config, Client: client}
}

func (ws *WebsocketServer) Start(app *fiber.App) {
	app.Use("/ws", middleware.LoggingMiddleware(), ws.Upgrade)
	app.Get("/ws/:id", middleware.LoggingMiddleware(), websocket.New(func(conn *websocket.Conn) {
		clientID := conn.Params("id")
		log.Printf("client %s connected", clientID)
		defer conn.Close()

		for {
			ctx := context.Background()
			key := clientID + ":process"
			exists, err := ws.Client.Exists(ctx, key).Result()
			if err != nil {
				log.Println(err)
			}

			if exists > 0 {
				processPercent, errget := ws.Client.Get(ctx, key).Result()
				if errget != nil {
					log.Println(errget)
				}

				if err = conn.WriteMessage(websocket.TextMessage, []byte(processPercent)); err != nil {
					log.Println("write:", err)
				}
			} else {
				if err = conn.WriteMessage(websocket.TextMessage, []byte("")); err != nil {
					log.Println("write:", err)
				}
			}

			time.Sleep(5 * time.Second)
		}
	}))
}

func (ws *WebsocketServer) Upgrade(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}
