package middleware

import (
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/ariyn/Lcd/Models"
	"github.com/ariyn/Lcd/Models/Errors"
	"github.com/ariyn/Lcd/Repositories"
	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string
	Password string
}

const SECRET_KEY = "SECRET"

var identityKey = "id"

var AuthRules = []Models.AuthRule{}

func InitJwtMiddleware() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:      "develop",
		Key:        []byte(SECRET_KEY),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(Models.User); ok {
				return jwt.MapClaims{
					identityKey: v.Account,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals Login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			user, err := Repositories.User.FIND_WITH_ID(loginVals.Username)
			if user.MatchPassword(loginVals.Password) {
				return user, nil
			}

			err = Errors.NotCorrectLoginInfo{
				Name: "JWT.Authenticator",
			}
			return nil, err
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
}
