package Repositories

import (
	"encoding/json"
	"log"
	"reflect"
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
		return nil, Errors.NoSuchUser{
			ID:     id,
			Name:   "Repository.FIND_WITH_ID",
			RawErr: err,
		}
	}
	if err != nil {
		log.Println("[Error]redis unknown error", err)
		return nil, Errors.RedisFailure{
			ID:     id,
			Name:   "Repository.FIND_WITH_ID",
			RawErr: err,
		}
	}

	if err := json.Unmarshal([]byte(val), &User); err != nil {
		log.Println("[Warning]json can't parse redis result", val, err)
		return nil, Errors.InvalidJson{
			Type:   reflect.TypeOf(User),
			Name:   "Repository.READ",
			RawErr: err,
		}
	}

	return User, nil
}

func (r *UserRepository) READ(key string) (*Models.User, error) {
	var User *Models.User

	redisKey := r.Prefix + "::" + key
	val, err := r.Client.Get(redisKey).Result()
	if err == redis.Nil {
		log.Printf("[Warning]redis %s does not exists", redisKey)
		return nil, Errors.NoSuchUser{
			ID:     redisKey,
			Name:   "Repository.READ",
			RawErr: err,
		}
	} else if err != nil {
		log.Println("[Error]redis unknown error", err)
		return nil, Errors.RedisFailure{
			Name:   "Repository.READ",
			RawErr: err,
		}
	}

	log.Println(redisKey, val)
	err = json.Unmarshal([]byte(val), &User)
	if err != nil {
		log.Println("[Warning]json can't parse redis result", val, err)
		return nil, Errors.InvalidJson{
			Type:   reflect.TypeOf(User),
			Name:   "Repository.READ",
			RawErr: err,
		}
	}

	return User, nil
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
	user, _ := r.READ(id)

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

		r.Client.Del(CreateKey(r.Prefix, user.Account))
	}

	return err
}

func (r *UserRepository) UPDATE(updateUser *Models.User) error {
	userID := strconv.Itoa(updateUser.ID)

	if r.EXISTS(userID) == false {
		err := Errors.NoSuchUser{
			Name: "Repository.UPDATE",
			ID:   userID,
		}

		log.Printf("[Error]redis failed to update article: %#v", err)
		return err
	}

	jsonData, err := json.Marshal(updateUser)
	if err != nil {
		log.Printf("[Error]json can't marshal user %#v", err)
		return Errors.InvalidJson{
			Name:   "Repository.UPDATE",
			Type:   reflect.TypeOf(updateUser),
			RawErr: err,
		}
	}

	redisKey := CreateKey(r.Prefix, userID)
	_, err = r.Client.Set(redisKey, jsonData, 0).Result()
	if err != nil {
		log.Printf("[Error]redis can't save user %#v", err)
		return Errors.RedisFailure{
			Name:   "Repository.UPDATE",
			RawErr: err,
		}
	}

	redisReverseKey := r.Prefix + "::" + updateUser.Account
	_, err = r.Client.Set(redisReverseKey, jsonData, 0).Result()
	if err != nil {
		log.Printf("[Error]redis can't save user %#v", err)
		return Errors.RedisFailure{
			Name:   "Repository.UPDATE",
			RawErr: err,
		}
	}

	return nil
}

func (r *UserRepository) EXISTS(userID string) bool {
	redisKey := CreateKey(r.Prefix, userID)
	return r.Client.Exists(redisKey).Val() == 1
}
