package Repositories

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/ariyn/Lcd/Models"
	"github.com/ariyn/Lcd/Models/Errors"
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

func (r *UserRepository) FIND_WITH_ID(id string) (*Models.User, error) {
	var User *Models.User

	redisKey := r.Prefix + "::" + id
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

	if err != nil {
		log.Println("[Error]redis failed to save User", jsonString, err)
		return -1, err
	}

	redisKey = r.Prefix + "::" + v.Account
	err = r.Client.Set(redisKey, jsonString, 0).Err()
	if err != nil {
		log.Println("[Error]redis failed to save user index", jsonString, err)
		return -1, err
	}

	return v.ID, nil
}

func (r *UserRepository) DELETE(id string) error {
	redisKey := CreateKey(r.Prefix, id)

	val, err := r.Client.Del(redisKey).Result()
	isExists := (val == 1)

	if err != nil {
		log.Println("[Error]redis failed to delete user", id, err)
		err = Errors.RedisFailure{
			Name:   "Repository.DELETE",
			RawErr: err,
		}
	} else if isExists == false {
		err = Errors.NoSuchUser{
			ID:   id,
			Name: "Repository.DELETE",
		}
	} else {
		r.Client.Decr(r.Prefix + ":counter")
	}

	return err
}
