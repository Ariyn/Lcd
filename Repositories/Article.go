package Repositories

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/ariyn/Lcd/Models"
	"github.com/go-redis/redis"
)

type ArticleRepository Repository

var Article *ArticleRepository

// InitArticleRepositorySingleton initializes repository.
// this function should be called before main logic.
// because it's not thread safe.
func InitArticleRepositorySingleton(redis *redis.Client) bool {
	if Article == nil {
		Article = &ArticleRepository{
			Client: redis,
			Prefix: "Article",
		}
		return true
	}
	return false
}

func (r *ArticleRepository) READ(key string) (*Models.Article, error) {
	var article *Models.Article

	redisKey := r.Prefix + "::" + key
	val, err := r.Client.Get(redisKey).Result()
	if err == redis.Nil {
		log.Printf("[Warning]redis %s does not exists", redisKey)
		r.Client.Set(redisKey, 1, 0)
		article = nil
	} else if err != nil {
		log.Println("[Error]redis unknown error", err)
		article = nil
	}

	err = json.Unmarshal([]byte(val), &article)
	if err != nil {
		log.Println("[Warning]json can't parse redis result", val, err)
		article = nil
	}

	return article, err
}

func (r *ArticleRepository) CREATE(v *Models.Article) (int, error) {
	counter := r.Client.Incr(r.Prefix + ":counter").Val()

	v.ID = int(counter)
	jsonString, err := json.Marshal(v)
	if err != nil {
		log.Println("[Error]json can't marshal article", err)
		return -1, err
	}

	redisKey := r.Prefix + "::" + strconv.Itoa(v.ID)
	err = r.Client.Set(redisKey, jsonString, 0).Err()

	fmt.Println(redisKey)
	if err != nil {
		log.Println("[Error]redis failed to save article", jsonString, err)
		return -1, err
	}

	return v.ID, nil
}
