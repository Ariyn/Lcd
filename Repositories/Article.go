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

func (r *ArticleRepository) READ(id string) (*Models.Article, error) {
	var article *Models.Article

	redisKey := CreateKey(r.Prefix, id)
	val, err := r.Client.Get(redisKey).Result()

	if err == redis.Nil {
		log.Printf("[Warning]redis %s does not exists", redisKey)
		article, err = nil, Errors.NoSuchArticle{
			ID:   id,
			Name: "Repository.READ",
		}
	} else if err != nil {
		log.Println("[Error]redis unknown error", err)
		article = nil
	} else {
		err = json.Unmarshal([]byte(val), &article)
		if err != nil {
			log.Printf("[Warning]json can't parse redis result: %s: %s\n", val, err)
			article, err = nil, Errors.InvalidJson{
				Type:   reflect.TypeOf(article),
				Name:   "Repository.READ",
				RawErr: err,
			}
		}
	}

	return article, err
}

func (r *ArticleRepository) CREATE(v *Models.Article) (string, error) {
	counterKey := r.Prefix + ":counter"
	counter := r.Client.Incr(counterKey).Val()
	v.ID = int(counter)

	jsonString, err := json.Marshal(v)
	if err != nil {
		log.Println("[Error]json can't marshal article", err)

		log.Println("[Info]redis decreases article counter", err)
		r.Client.Decr(counterKey)

		return "", Errors.InvalidJson{
			Type:   reflect.TypeOf(v),
			Name:   "Repository.CREATE",
			RawErr: err,
		}
	}

	id := strconv.FormatInt(counter, 10)
	redisKey := r.Prefix + "::" + id
	err = r.Client.Set(redisKey, jsonString, 0).Err()

	if err != nil {
		log.Println("[Error]redis failed to save article", jsonString, err)

		log.Println("[Info]redis decreases article counter", err)
		r.Client.Decr(counterKey)

		return "", Errors.RedisFailure{
			Name:   "Repository.CREATE",
			RawErr: err,
		}
	}

	return id, nil
}

func (r *ArticleRepository) DELETE(id string) error {
	redisKey := CreateKey(r.Prefix, id)

	val, err := r.Client.Del(redisKey).Result()
	isExists := (val == 1)

	if err != nil {
		log.Println("[Error]redis failed to delete article", id, err)
		err = Errors.RedisFailure{
			Name:   "Repository.DELETE",
			RawErr: err,
		}
	} else if isExists == false {
		err = Errors.NoSuchArticle{
			ID:   id,
			Name: "Repository.DELETE",
		}
	} else {
		r.Client.Decr(r.Prefix + ":counter")
	}

	return err
}

func (r *ArticleRepository) UPDATE(v *Models.Article) (string, error) {
	articleId := strconv.Itoa(v.ID)

	if r.EXISTS(strconv.Itoa(v.ID)) == false {
		err := Errors.NoSuchArticle{
			ID:   CreateKey(r.Prefix, articleId),
			Name: "Repository.UPDATE",
		}
		log.Printf("[Error]redis failed to update article: %#v", err)

		return "", err
	}

	jsonString, err := json.Marshal(v)
	if err != nil {
		log.Printf("[Error]json can't marshal article %#v", err)

		return "", Errors.InvalidJson{
			Type:   reflect.TypeOf(v),
			Name:   "Repository.CREATE",
			RawErr: err,
		}
	}

	redisKey := r.Prefix + "::" + articleId
	err = r.Client.Set(redisKey, jsonString, 0).Err()

	if err != nil {
		log.Println("[Error]redis failed to update article", jsonString, err)

		return "", Errors.RedisFailure{
			Name:   "Repository.CREATE",
			RawErr: err,
		}
	}

	return articleId, nil
}

func (r *ArticleRepository) EXISTS(id string) bool {
	redisKey := CreateKey(r.Prefix, id)
	return r.Client.Exists(redisKey).Val() == 1
}
