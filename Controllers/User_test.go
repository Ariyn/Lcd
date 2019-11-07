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

const _MANY_USER_SIZE = 100

func TestGetUser(t *testing.T) {
	user := &Models.User{}
	if err := createUser(user); err != nil {
		assert.Fail(t, "Can't create user", err)
		return
	}
	defer deleteUser(user)
	path := UserController.Path + "/" + strconv.Itoa(user.ID)

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

func TestPutUser(t *testing.T) {
	user := &Models.User{}
	if err := createUser(user); err != nil {
		assert.Fail(t, "Can't create user", err)
		return
	}
	defer deleteUser(user)

	expected := Models.User(*user)
	expected.Nickname = "updated user nickname"

	path := UserController.Path + "/" + strconv.Itoa(user.ID)
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
	user := &Models.User{}
	if err := createUser(user); err != nil {
		assert.Fail(t, "Can't create user", err)
		return
	}
	defer deleteUser(user)

	path := UserController.Path + "/" + strconv.Itoa(user.ID)
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

func TestGetEntireUser(t *testing.T) {
	ids := createManyUsers()
	defer deleteManyUsers(ids)
	size := 100

	path := UserController.Path

	w := requestUser(request{
		path:   path,
		method: Models.GET,
		query: map[string]string{
			"size": strconv.Itoa(size),
		},
	})

	assert.Equal(t, 200, w.Code)

	var users map[string]Models.User
	if err := parseUsers(w.Body, &users); err != nil {
		assert.Fail(t, "Can't parse users", err, w.Body.String())
		return
	}

	assert.Equal(t, _MANY_USER_SIZE, len(users))
	for i := 1; i <= size; i++ {
		index := strconv.Itoa(i)

		user, ok := users[index]
		assert.True(t, ok)
		assert.Equal(t, index, user.Nickname)
	}
}

func TestDeleteUser(t *testing.T) {
	user := &Models.User{}
	if err := createUser(user); err != nil {
		log.Printf("Can't create user at %d: %#v\n", user.ID, err)
		return
	}
	defer deleteUser(user)

	w := requestUser(request{
		path:   UserController.Path + "/" + strconv.Itoa(user.ID),
		method: Models.DELETE,
	})

	assert.Equal(t, 200, w.Code)
	err := getUser(strconv.Itoa(user.ID), nil)
	assert.EqualError(t, err, "no such user User::"+strconv.Itoa(user.ID))
}

func TestDeleteUserWithNotExistsUser(t *testing.T) {
	w := requestUser(request{
		path:   UserController.Path + "/" + "-1",
		method: Models.DELETE,
	})

	assert.Equal(t, 400, w.Code)
}

func parseUsers(body *bytes.Buffer, users *map[string]Models.User) error {
	return json.Unmarshal(body.Bytes(), users)
}

func createManyUsers() []string {
	ids := []string{}

	for i := 1; i <= _MANY_USER_SIZE; i++ {
		user := &Models.User{
			Account:  _SAMPLE_USER_ACCOUNT,
			Nickname: strconv.Itoa(i),
		}
		err := createUser(user)
		if err != nil {
			log.Printf("Can't create user at %d, %d: %#v\n", i, user.ID, err)
			break
		}

		ids = append(ids, strconv.Itoa(user.ID))
	}

	log.Printf("created %d users", len(ids))
	if len(ids) != _MANY_USER_SIZE {
		deleteManyUsers(ids)
		ids = []string{}
	}

	return ids
}

func deleteManyUsers(ids []string) {
	for i, userID := range ids {
		if err := deleteUserWithID(userID); err != nil {
			log.Printf("Can't delete %d user, %#v", i, err)
		}
	}
}

func createUser(user *Models.User) error {
	if user == nil {
		user = &Models.User{
			Account: _SAMPLE_USER_ACCOUNT,
		}
	}

	_, err := Repositories.User.CREATE(user)
	return err
}

func deleteUser(user *Models.User) {
	Repositories.User.DELETE(strconv.Itoa(user.ID))
}

func deleteUserWithID(id string) error {
	return Repositories.User.DELETE(id)
}

func getUser(id string, user *Models.User) error {
	loadedUser, err := Repositories.User.READ(id)
	if loadedUser != nil {
		*user = *loadedUser
	}
	return err
}

func requestUser(r request) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(string(r.method), r.path, r.getBodyReader())
	req.Header.Add("Content-Type", r.getContentType())
	req.Header.Add("Authorization", "Bearer "+r.authToken)

	q := req.URL.Query()
	for key, value := range r.query {
		log.Println(key, value)
		q.Add(key, value)
	}

	req.URL.RawQuery = q.Encode()

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
