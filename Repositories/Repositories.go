package Repositories

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

type Repository struct {
	Prefix string
	Client *redis.Client
}

var initializers = []interface{}{
	InitArticleRepositorySingleton,
	InitUserRepositorySingleton,
}

func initRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.0.2:6379",
		Password: "foobared",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong)
	if err != nil {
		log.Println("[Error] Can't initialize redis")
		log.Println("[Error] Exiting server with log.Fatal")
		log.Fatal(err)
	}

	return client
}

func Initialize() {
	client := initRedis()

	for _, init := range initializers {
		log.Printf("Initializeing %v", init)
		init.(func(*redis.Client) bool)(client)
	}
}
