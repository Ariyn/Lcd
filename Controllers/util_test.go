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

	"github.com/ariyn/Lcd/Models"
	"github.com/ariyn/Lcd/Repositories"
	"github.com/gin-gonic/gin"
)

var (
	AnonymousRoleUser = &Models.User{Role: Models.RoleAnonymous}
	UserRoleUser      = &Models.User{Role: Models.RoleUser}
	EditorRoleUser    = &Models.User{Role: Models.RoleEditor}
	AdminRoleUser     = &Models.User{Role: Models.RoleAdmin}
)

type request struct {
	method      Models.Method
	path        string
	article     *Models.Article
	user        *Models.User
	authUser    *Models.User
	body        string
	contentType string
	authToken   string
	query       map[string]string
}

func (r request) getBodyReader() *bytes.Buffer {
	if r.article != nil {
		data, _ := json.Marshal(r.article)
		return bytes.NewBuffer(data)
	}

	if r.user != nil {
		data, _ := json.Marshal(r.user)
		return bytes.NewBuffer(data)
	}

	return bytes.NewBufferString(r.body)
}

func (r request) getContentType() string {
	if r.contentType == "" {
		return "application/json"
	}
	return r.contentType
}

func (r request) getAuthUser() *Models.User {
	if r.authUser != nil {
		return r.authUser
	} else if r.user != nil {
		return r.user
	}
	return nil
}

func authorizingMiddleware(c *gin.Context) {
	user := &Models.User{}
	authHeader := c.GetHeader("testAuth")

	log.Println(authHeader)
	if authHeader == "" {
		user = &Models.User{Role: Models.RoleAnonymous}
	} else if err := json.Unmarshal([]byte(authHeader), user); err != nil {
		panic("Invalid authorizing header " + err.Error())
	}

	c.Set("user", user)
	c.Next()
}

func createUser(user *Models.User) error {
	if user == nil {
		user = &Models.User{
			Account: _SAMPLE_USER_ACCOUNT,
		}
	}

	if user.Role == 0 {
		user.Role = Models.RoleUser
	}

	_, err := Repositories.User.CREATE(user)
	return err
}

func deleteUser(user *Models.User) {
	Repositories.User.DELETE(strconv.Itoa(user.ID))
}

func deleteManyUsers(ids []string) {
	for _, userID := range ids {
		user := &Models.User{}
		if err := getUser(userID, user); err != nil {
			log.Printf("can't delete such user %s, %#v", userID, err)
		} else {
			deleteUser(user)
		}
	}
}

func getUser(id string, user *Models.User) error {
	loadedUser, err := Repositories.User.READ(id)
	if loadedUser != nil {
		*user = *loadedUser
	}
	return err
}

func enableLog() {
	log.SetOutput(os.Stdout)
}

func disenableLog() {
	log.SetOutput(ioutil.Discard)
}

func doRequest(r request) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(string(r.method), r.path, r.getBodyReader())
	req.Header.Add("Content-Type", r.getContentType())
	req.Header.Add("Authorization", "Bearer "+r.authToken)

	if testAuth, err := json.Marshal(r.getAuthUser()); err == nil {
		req.Header.Add("testAuth", string(testAuth))
	}

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
