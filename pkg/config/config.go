package config

import (
	"github.com/gofiber/websocket/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"jubeliotesting/internal/dto"
	"sync"
)

type GetEnvConfig struct {
	Host          string `env:"HOST" envDefault:"host=192.168.100.73 user=ridwan password=M1r34cl3 dbname=jubeliotest port=5432 sslmode=disable TimeZone=Asia/Jakarta"`
	RedisAddr     string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	APIKey        string `env:"API_KEY" envDefault:"SECRET_KEY"`
	TempoEndpoint string `env:"COLLECTOR_ENDPOINT" envDefault:"localhost:4318"`
}

var (
	clients map[string]*websocket.Conn
	mu      sync.RWMutex
)

func SetConn(clientID string, c *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()
	clients[clientID] = c
}

func GetConn(clientID string) (*websocket.Conn, bool) {
	mu.RLock()
	defer mu.RUnlock()
	conn, ok := clients[clientID]
	return conn, ok
}

func RemoveConn(clientID string) {
	mu.Lock()
	defer mu.Unlock()
	delete(clients, clientID)
}

func Broadcast(message dto.MessageSocket) {
	mu.RLock()
	defer mu.RUnlock()
	for id, conn := range clients {
		if err := conn.WriteJSON(message); err != nil {
			conn.Close()
			delete(clients, id)
		}
	}
}

func BroadcastToClient(clientID string, message dto.MessageSocket) {
	mu.RLock()
	defer mu.RUnlock()
	for id, conn := range clients {
		if id == clientID {
			if err := conn.WriteJSON(message); err != nil {
				conn.Close()
				delete(clients, id)
			}
		}
	}
}

func CreateConnection(config GetEnvConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.Host), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewRedisClient(config GetEnvConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: "",
		DB:       1,
	})
}
