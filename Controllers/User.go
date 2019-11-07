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

const (
	_DEFAULT_PAGINATION_SIZE = 100
)

var UserController Controller = Controller{
	Path: "/user",
	Handlers: []Handler{
		Handler{Path: "", Method: Models.GET, Handler: getEntireUser},
		Handler{Path: "", Method: Models.POST, Handler: postUser},
		Handler{Path: "/:userID", Method: Models.GET, Handler: getUser, UseAuth: true},
		Handler{Path: "/:userID", Method: Models.PUT, Handler: putUser, UseAuth: true},
		Handler{Path: "/:userID", Method: Models.DELETE, Handler: deleteUser, UseAuth: true},
	},
}

type paging struct {
	Start int `form:"start, default=1"`
	Size  int `form:"size, default=100"`
}

type success struct {
	code    int
	message string
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

func getEntireUser(c *gin.Context) {
	var paging paging

	err := c.BindQuery(&paging)
	if err != nil {
		panic(&HTTPError{badRequest, c.Request.URL.Path, "Not enought paramters", err})
	}

	userIDs := []string{}

	// redis index starts from 1.
	// start 0, size 100 means 100 users.
	// so make sure result be correct size
	for i := paging.Start; i < paging.Start+paging.Size+1; i++ {
		userIDs = append(userIDs, strconv.Itoa(i))
	}

	if users, err := Repositories.User.MREAD_WITH_KEY(userIDs); err != nil {
		panic(&HTTPError{internalServerError, c.Request.URL.Path, "Can't read redis", err})
	} else {
		c.JSON(200, users)
	}
}

func putUser(c *gin.Context) {
	userID := c.Param("userID")
	user, err := Repositories.User.READ(userID)
	log.Printf("read user %#v, %#v", user, err)
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

	updateUser := &Models.User{}
	if err := getUserDTO(c, updateUser); err != nil {
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
	user := &Models.User{}
	if err := getUserDTO(c, user); err != nil {
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

func deleteUser(c *gin.Context) {
	userID := c.Param("userID")

	// TODO: check user role

	if Repositories.User.EXISTS(userID) == false {
		panic(&HTTPError{badRequest, c.Request.URL.Path, "No such user", nil})
	}

	if err := Repositories.User.DELETE(userID); err != nil {
		panic(&HTTPError{internalServerError, c.Request.URL.Path, "Can't remove user", err})
	}

	c.JSON(200, success{
		code:    200,
		message: "success to delete",
	})
}

func getUserDTO(c *gin.Context, user *Models.User) error {
	data, err := c.GetRawData()
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, user); err != nil {
		return err
	}

	return nil
}

func isNoSuchUser(err error) bool {
	_, ok := err.(Errors.NoSuchUser)
	return ok
}

func isRedisFailure(err error) bool {
	_, ok := err.(Errors.RedisFailure)
	return ok
}
