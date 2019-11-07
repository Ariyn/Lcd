package middleware

import (
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/ariyn/Lcd/Models"
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

func AddAuthRules(rule *Models.AuthRule) bool {
	// TODO: check duplicated rule
	AuthRules = append(AuthRules, *rule)

	return true
}

func InitJwtMiddleware() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:         "develop",
		Key:           []byte(SECRET_KEY),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		DisabledAbort: true,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*Models.User); ok {
				return jwt.MapClaims{
					identityKey: v.Account,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			log.Printf("identity handler %#v, %#v, %#v\n", claims, identityKey, claims[identityKey])
			return &Models.User{
				Account: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals Login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			user, err := Repositories.User.FIND_WITH_ID(loginVals.Username)
			if err == nil && user.MatchPassword(loginVals.Password) {
				return user, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			if allowAnonyRequest(c) && isEmptyError(message) {
				log.Println(message)
				identity := "anonymous"
				c.Set("JWT_TOKEN", "")
				c.Set("JWT_PAYLOAD", jwt.MapClaims{identityKey: identity})
				c.Set(identityKey, identity)

				c.Next()
			} else {
				c.Abort()
				c.JSON(code, gin.H{
					"code":    code,
					"message": message,
				})
			}
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
}

func allowAnonyRequest(c *gin.Context) bool {
	expect := &Models.AuthRule{
		Method:   Models.Method(c.Request.Method),
		FullPath: c.FullPath(),
	}

	isAllow := false
	for _, rule := range AuthRules {
		if !rule.AllowAnonymous {
			continue
		}

		if rule.FullPath == expect.FullPath && rule.Method == expect.Method {
			isAllow = true
			break
		}
	}

	return isAllow
}

func isEmptyError(message string) bool {
	return message == jwt.ErrEmptyCookieToken.Error() || message == jwt.ErrEmptyAuthHeader.Error() || message == jwt.ErrEmptyParamToken.Error() || message == jwt.ErrEmptyQueryToken.Error()
}
