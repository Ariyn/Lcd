package middleware_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/ariyn/Lcd/Models"
	"github.com/ariyn/Lcd/Repositories"
	"github.com/ariyn/Lcd/Service/middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type request struct {
	method        Models.Method
	path          string
	body          string
	loginBody     *middleware.Login
	contentType   string
	authorization string
}

func (r request) getBodyReader() *bytes.Buffer {
	if r.loginBody != nil {
		buf, _ := json.Marshal(r.loginBody)
		return bytes.NewBuffer(buf)
	}
	return bytes.NewBufferString(r.body)
}

func (r request) getContentType() string {
	if r.contentType == "" {
		return "application/json"
	}
	return r.contentType
}

var _GIN_ENGINE *gin.Engine
var _JWT_MIDDLEWARE *jwt.GinJWTMiddleware

const (
	_TEST_DB    = 15
	authGroup   = "/auth"
	authAction  = "/authAction"
	anonyAction = "/anonyAction"
)
const (
	_SAMPLE_USER_ACCOUNT  = "test_user"
	_SAMPLE_USER_PASSWORD = "test_password"
)

func TestMain(m *testing.M) {
	client := Repositories.InitRedis(_TEST_DB)
	Repositories.Initialize(client)

	_GIN_ENGINE = gin.Default()

	_JWT_MIDDLEWARE, _ = middleware.InitJwtMiddleware()
	_GIN_ENGINE.Use(_JWT_MIDDLEWARE.MiddlewareFunc())

	_GIN_ENGINE.POST("/login", _JWT_MIDDLEWARE.LoginHandler)
	middleware.AddAuthRules(&Models.AuthRule{
		FullPath:       "/login",
		Method:         Models.POST,
		Role:           Models.RoleUser,
		AllowAnonymous: true,
	})

	group := _GIN_ENGINE.Group(authGroup)
	group.GET(authAction, func(c *gin.Context) {
		c.JSON(200, "")
	})
	middleware.AddAuthRules(&Models.AuthRule{
		FullPath:       authGroup + authAction,
		Method:         Models.GET,
		Role:           Models.RoleUser,
		AllowAnonymous: false,
	})

	group.GET(anonyAction, func(c *gin.Context) {
		c.JSON(200, "")
	})
	middleware.AddAuthRules(&Models.AuthRule{
		FullPath:       authGroup + anonyAction,
		Method:         Models.GET,
		Role:           Models.RoleUser,
		AllowAnonymous: true,
	})

	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestJwtLogin(t *testing.T) {
	loginData, err := createUser()
	if err != nil {
		assert.Fail(t, "Can't create user", err)
		return
	}
	defer deleteUser(loginData.Username)

	w := doRequest(request{
		method:    Models.POST,
		path:      "/login",
		loginBody: loginData,
	})

	assert.Equal(t, 200, w.Code)
}

func TestJwtLoginWithInvalidBody(t *testing.T) {
	loginData, err := createUser()
	if err != nil {
		assert.Fail(t, "Can't create user", err)
		return
	}
	defer deleteUser(loginData.Username)

	w := doRequest(request{
		method: Models.POST,
		path:   "/login",
		body:   "",
	})

	assert.Equal(t, 401, w.Code)
}

func TestJwtLoginWithInvalidLoginInfo(t *testing.T) {
	loginData, err := createUser()
	if err != nil {
		assert.Fail(t, "Can't create user", err)
		return
	}
	defer deleteUser(loginData.Username)
	loginData.Password = "WRONG PASSWORD"

	w := doRequest(request{
		method:    Models.POST,
		path:      "/login",
		loginBody: loginData,
	})

	assert.Equal(t, 401, w.Code)
}

func TestJwtAuthorizationWithInvalidToken(t *testing.T) {
	token := "invalid token"

	w := doRequest(request{
		path:          authGroup + authAction,
		method:        Models.GET,
		authorization: token,
	})

	assert.Equal(t, 401, w.Code)
}

func TestJwtIdentityHandler(t *testing.T) {
	token, err, deleteFunc := getJwtToken()
	if err != nil {
		assert.Fail(t, "Can't generate token", err)
		return
	}
	defer deleteFunc()

	_GIN_ENGINE.GET("/test", func(c *gin.Context) {
		assert.Equal(t, _SAMPLE_USER_ACCOUNT, jwt.ExtractClaims(c)["id"])
	})

	w := doRequest(request{
		path:          authGroup + authAction,
		method:        Models.GET,
		authorization: token,
	})

	assert.Equal(t, 200, w.Code)
}

func TestJwtAnonyActionWithLogin(t *testing.T) {
	token, err, deleteFunc := getJwtToken()
	if err != nil {
		assert.Fail(t, "Can't generate token", err)
		return
	}
	defer deleteFunc()

	w := doRequest(request{
		method:        Models.GET,
		path:          authGroup + anonyAction,
		authorization: token,
	})

	assert.Equal(t, 200, w.Code)
}

func TestJwtAnonyActionWithoutLogin(t *testing.T) {
	w := doRequest(request{
		method: Models.GET,
		path:   authGroup + anonyAction,
	})

	assert.Equal(t, 200, w.Code)
}

func createUser() (*middleware.Login, error) {
	user := Models.User{
		Account:  _SAMPLE_USER_ACCOUNT,
		Password: _SAMPLE_USER_PASSWORD,
	}

	_, err := Repositories.User.CREATE(&user)

	return &middleware.Login{
		Username: user.Account,
		Password: user.Password,
	}, err
}

func deleteUser(id string) {
	user, _ := Repositories.User.FIND_WITH_ID(id)

	Repositories.User.DELETE(id)
	Repositories.User.DELETE(strconv.Itoa(user.ID))
}

func getJwtToken() (string, error, func()) {
	user := &Models.User{
		Account:  _SAMPLE_USER_ACCOUNT,
		Password: _SAMPLE_USER_PASSWORD,
	}

	Repositories.User.CREATE(user)

	token, _, err := _JWT_MIDDLEWARE.TokenGenerator(user)

	return token, err, func() {
		deleteUser(_SAMPLE_USER_ACCOUNT)
	}
}

func doRequest(r request) *httptest.ResponseRecorder {
	body := r.getBodyReader()
	req, _ := http.NewRequest(string(r.method), r.path, body)
	req.Header.Add("Content-Type", r.getContentType())
	req.Header.Add("Authorization", "Bearer "+r.authorization)

	w := httptest.NewRecorder()
	_GIN_ENGINE.ServeHTTP(w, req)

	return w
}
