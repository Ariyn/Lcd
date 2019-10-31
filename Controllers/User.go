package Controllers

import (
	"encoding/json"
	"io/ioutil"

	"github.com/ariyn/Lcd/Models"
	"github.com/ariyn/Lcd/Repositories"
	"github.com/gin-gonic/gin"
)

var UserController Controller = Controller{
	Path: "user",
	Handlers: []Handler{
		Handler{Path: "/", Method: GET, Handler: UserGetHandler},
	},
}

func UserGetHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"test": 1,
	})
}

func UserPostHandler(c *gin.Context) {
	bodyReader, err := c.Request.GetBody()
	if err != nil {
		panic(&HTTPError{BadRequest, c.Request.URL.String(), "can't read post body", err})
	}

	body, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		panic(&HTTPError{BadRequest, c.Request.URL.String(), "can't read post body", err})
	}

	var user *Models.User
	err = json.Unmarshal(body, user)
	if err != nil {
		panic(&HTTPError{BadRequest, c.Request.URL.String(), "can't parse body to user", err})
	}

	userID, err := Repositories.User.CREATE(user)
	if err != nil {
		panic(&HTTPError{InternalServerError, c.Request.URL.String(), "can't save user", err})
	}

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "user saved",
		"userID":  userID,
	})
}
