package Repositories

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/ariyn/Lcd/Models"
	"github.com/go-redis/redis"
)

type UserRepository Repository

var User *UserRepository

// InitUserRepositorySingleton initializes repository.
// this function should be called before main logic.
// because it's not thread safe.
func InitUserRepositorySingleton(redis *redis.Client) bool {
	if User == nil {
		User = &UserRepository{
			Client: redis,
			Prefix: "User",
		}
		return true
	}
	return false
}

func (r *UserRepository) READ(key string) (*Models.User, error) {
	var User *Models.User

	redisKey := r.Prefix + "::" + key
	val, err := r.Client.Get(redisKey).Result()
	if err == redis.Nil {
		log.Printf("[Warning]redis %s does not exists", redisKey)
		r.Client.Set(redisKey, 1, 0)
		User = nil
	} else if err != nil {
		log.Println("[Error]redis unknown error", err)
		User = nil
	}

	err = json.Unmarshal([]byte(val), &User)
	if err != nil {
		log.Println("[Warning]json can't parse redis result", val, err)
		User = nil
	}

	return User, err
}

func (r *UserRepository) CREATE(v *Models.User) (int, error) {
	counter := r.Client.Incr(r.Prefix + ":counter").Val()

	v.ID = int(counter)
	jsonString, err := json.Marshal(v)
	if err != nil {
		log.Println("[Error]json can't marshal User", err)
		return -1, err
	}

	redisKey := r.Prefix + "::" + strconv.Itoa(v.ID)
	err = r.Client.Set(redisKey, jsonString, 0).Err()

	fmt.Println(redisKey)
	if err != nil {
		log.Println("[Error]redis failed to save User", jsonString, err)
		return -1, err
	}

	return v.ID, nil
}
