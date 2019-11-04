package Repositories

import (
	"log"
	"reflect"
	"runtime"

	"github.com/go-redis/redis"
)

type Repository struct {
	Prefix string
	Client *redis.Client
}

var Client *redis.Client = nil

var initializers = []interface{}{
	InitArticleRepositorySingleton,
	InitUserRepositorySingleton,
}

func InitRedis(db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.0.2:6379",
		Password: "foobared",
		DB:       db,
	})

	pong, err := client.Ping().Result()
	if err != nil || pong != "PONG" {
		log.Println("[Error] Can't initialize redis")
		log.Println("[Error] Exiting server with log.Fatal")
		log.Fatal(err)
	}

	return client
}

func Initialize(client *redis.Client) {
	Client = client

	for _, init := range initializers {
		funcValue := reflect.ValueOf(init)
		funcName := runtime.FuncForPC(funcValue.Pointer()).Name()

		log.Printf("Initializeing %s", funcName)
		init.(func(*redis.Client) bool)(client)
	}
}

func CreateKey(prefix, key string) string {
	return prefix + "::" + key
}
