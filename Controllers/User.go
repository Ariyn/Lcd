package Controllers

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/ariyn/Lcd/Models"
	"github.com/ariyn/Lcd/Models/Errors"
	"github.com/ariyn/Lcd/Repositories"
	"github.com/gin-gonic/gin"
)

var UserController Controller = Controller{
	Path: "/user",
	Handlers: []Handler{
		Handler{Path: "", Method: Models.POST, Handler: postUser},
		Handler{Path: "/:userID", Method: Models.GET, Handler: getUser, UseAuth: true},
		Handler{Path: "/:userID", Method: Models.PUT, Handler: putUser, UseAuth: true},
	},
}

func getUser(c *gin.Context) {
	userID := c.Param("userID")
	p, ok := c.Get("JWT_PAYLOAD")
	log.Printf("%#v %#v\n", p, ok)

	p, ok = c.Get("JWT_TOKEN")
	log.Printf("%#v %#v\n", p, ok)

	if strings.ToLower(userID) == "me" {
		claims := jwt.ExtractClaims(c)
		log.Println(c.Request.Header.Get("Authorization"))
		log.Println(claims)
	}

	user, err := Repositories.User.FIND_WITH_ID(userID)
	if err != nil {
		url := c.Request.URL.Path
		if _, ok := err.(Errors.RedisFailure); ok {
			panic(&HTTPError{internalServerError, url, "Redis error", err})
		}
		if _, ok := err.(Errors.NoSuchUser); ok {
			panic(&HTTPError{badRequest, url, "No such user", err})
		}
		panic(&HTTPError{internalServerError, url, "Unknown error", err})
	}

	c.JSON(200, user)
}

func putUser(c *gin.Context) {
	userID := c.Param("userID")
	user, err := Repositories.User.READ(userID)
	if err != nil {
		url := c.Request.URL.Path
		if isNoSuchUser(err) {
			panic(&HTTPError{badRequest, url, "No such user", err})
		} else if isRedisFailure(err) {
			panic(&HTTPError{internalServerError, url, "Internal server error", err})
		} else {
			panic(&HTTPError{internalServerError, url, "Unknown redis error", err})
		}
	}

	updateUser, err := getUserDTO(c)
	if err != nil {
		panic(&HTTPError{badRequest, c.Request.URL.Path, "Bad request", err})
	}

	updateUser.ID, err = strconv.Atoi(userID)
	if err != nil {
		panic(&HTTPError{badRequest, c.Request.URL.Path, "Invalid user id", err})
	}

	err = Repositories.User.UPDATE(updateUser)
	if err != nil {
		url := c.Request.URL.Path
		if isRedisFailure(err) {
			panic(&HTTPError{internalServerError, url, "Internal server error", err})
		}
		panic(&HTTPError{internalServerError, url, "Unknown redis error", err})
	}

	c.JSON(200, user)
}

func postUser(c *gin.Context) {
	user, err := getUserDTO(c)
	if err != nil {
		panic(&HTTPError{badRequest, c.Request.URL.Path, "Bad request", err})
	}

	userID, err := Repositories.User.CREATE(user)
	if err != nil {
		url := c.Request.URL.Path
		if isRedisFailure(err) {
			panic(&HTTPError{internalServerError, url, "Internal server error", err})
		} else {
			panic(&HTTPError{internalServerError, url, "Unknown redis error", err})
		}
	}

	user, _ = Repositories.User.READ(strconv.Itoa(userID))
	c.JSON(200, user)
}

func getUserDTO(c *gin.Context) (*Models.User, error) {
	user := &Models.User{}

	data, err := c.GetRawData()
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, user); err != nil {
		return nil, err
	}

	return user, nil
}

func isNoSuchUser(err error) bool {
	_, ok := err.(Errors.NoSuchUser)
	return ok
}

func isRedisFailure(err error) bool {
	_, ok := err.(Errors.RedisFailure)
	return ok
}
