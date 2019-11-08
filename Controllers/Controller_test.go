package Controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ariyn/Lcd/Models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetUserOrAnonymous(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	expected := &Models.User{Account: "test"}
	c.Set("user", expected)

	actual := getUserOrAnonymous(c)
	assert.Equal(t, expected, actual)
}

func TestGetUserOrAnonymousWithNil(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// c.Get("user") will return nil. therefore getUserOrAnonymous returns anonymous user
	actual := getUserOrAnonymous(c)
	assert.Equal(t, Models.RoleAnonymous, actual.Role)
}

func TestGetUserOrAnonymousWithInvalidData(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	path := "/test"
	c.Request, _ = http.NewRequest("get", path, nil)

	c.Set("user", 1)
	defer func() {
		recovered := recover()
		err, ok := recovered.(*HTTPError)
		assert.True(t, ok, "Not *HTTPError")

		assert.Equal(t, path, err.path)
		assert.Equal(t, "Can't load user authentication", err.message)
	}()

	getUserOrAnonymous(c)
}
