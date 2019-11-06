package Controllers_test

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

	. "github.com/ariyn/Lcd/Controllers"
	"github.com/ariyn/Lcd/Models"
	"github.com/ariyn/Lcd/Repositories"
	"github.com/stretchr/testify/assert"
)

const (
	_SAMPLE_USER_ACCOUNT  = "account"
	_SAMPLE_USER_PASSWORD = "password"
)

func TestGetUser(t *testing.T) {
	id, err := createUser()
	if err != nil {
		assert.Fail(t, "Can't create user", err)
		return
	}
	defer deleteUser(getUser(id))
	path := UserController.Path + "/" + id

	w := requestUser(request{
		method: Models.GET,
		path:   path,
	})

	assert.Equal(t, 200, w.Code)
}

func TestGetUserWithNotExistsUserID(t *testing.T) {
	path := UserController.Path + "/" + "-1"

	w := requestUser(request{
		method: Models.GET,
		path:   path,
	})

	assert.Equal(t, 400, w.Code)
}

// func TestGetUserWithME(t *testing.T) {
// 	return
// 	id, err := createUser()
// 	if err != nil {
// 		assert.Fail(t, "Can't create user", err)
// 		return
// 	}
// 	user := getUser(id)
// 	log.Println(user)
// 	defer deleteUser(user)
// 	token, err := getJwtToken(user)
// 	if err != nil {
// 		assert.Fail(t, "Can't generate jwt token", err)
// 		return
// 	}
// 	path := UserController.Path + "/" + "me"
//
// 	w := requestUser(request{
// 		method:    Models.GET,
// 		path:      path,
// 		authToken: token,
// 	})
//
// 	assert.Equal(t, 200, w.Code)
// }

func TestPutUser(t *testing.T) {
	userID, err := createUser()
	if err != nil {
		assert.Fail(t, "Can't create user", err)
		return
	}
	user := getUser(userID)
	defer deleteUser(user)

	expected := Models.User(*user)
	expected.Nickname = "updated user nickname"

	path := UserController.Path + "/" + userID
	w := requestUser(request{
		method: Models.PUT,
		path:   path,
		user:   &expected,
	})

	assert.Equal(t, 200, w.Code)
	actual, _ := Repositories.User.READ(strconv.Itoa(user.ID))
	assert.Equal(t, &expected, actual)
}

func TestPutUserWithNotExistsUser(t *testing.T) {
	user := &Models.User{}

	path := UserController.Path + "/" + "-1"
	w := requestUser(request{
		method: Models.PUT,
		path:   path,
		user:   user,
	})

	assert.Equal(t, 400, w.Code)
}

func TestPutUserWithInvalidJson(t *testing.T) {
	userID, err := createUser()
	if err != nil {
		assert.Fail(t, "Can't create user", err)
		return
	}
	defer deleteUser(getUser(userID))

	path := UserController.Path + "/" + userID
	w := requestUser(request{
		method: Models.PUT,
		path:   path,
		body:   `invalid json`,
	})
	// TODO: body:`{"test":"Test"}` is also invalid json, but server will not return invelid json error

	assert.Equal(t, 400, w.Code)
}

func TestPostUser(t *testing.T) {
	var actual *Models.User
	defer func() {
		deleteUser(actual)
	}()

	expected := &Models.User{
		Account: _SAMPLE_USER_ACCOUNT,
	}

	path := UserController.Path
	w := requestUser(request{
		path:   path,
		method: Models.POST,
		user:   expected,
	})

	assert.Equal(t, 200, w.Code)

	var err error
	actual, err = unmarshalResponseUser(w.Body)
	assert.NoError(t, err, "parse response")

	expected.ID = actual.ID
	assert.Equal(t, expected, actual)
}

func TestPostUserWithInvalidJson(t *testing.T) {
	path := UserController.Path
	w := requestUser(request{
		path:   path,
		method: Models.POST,
		body:   "invalid json",
	})

	assert.Equal(t, 400, w.Code)
}

func createUser() (string, error) {
	user := &Models.User{
		Account: _SAMPLE_USER_ACCOUNT,
	}
	id, err := Repositories.User.CREATE(user)
	return strconv.Itoa(id), err
}

func deleteUser(user *Models.User) {
	Repositories.User.DELETE(strconv.Itoa(user.ID))
}

func getUser(id string) *Models.User {
	user, _ := Repositories.User.READ(id)
	return user
}

func getJwtToken(user *Models.User) (string, error) {
	token, _, err := JWT_MIDDLEWARE.TokenGenerator(user)
	return token, err
}

func requestUser(r request) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(r.method, r.path, r.getBodyReader())
	req.Header.Add("Content-Type", r.getContentType())
	req.Header.Add("Authorization", "Bearer "+r.authToken)

	w := httptest.NewRecorder()
	GIN_ENGINE.ServeHTTP(w, req)

	return w
}

func unmarshalResponseUser(body *bytes.Buffer) (*Models.User, error) {
	user := &Models.User{}

	err := json.Unmarshal(body.Bytes(), user)
	return user, err
}

func enableLog() {
	log.SetOutput(os.Stdout)
}

func disenableLog() {
	log.SetOutput(ioutil.Discard)
}
