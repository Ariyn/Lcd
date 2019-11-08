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
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const TEST_DATABASE = 15
const (
	SAMPLE_TITLE = "sample"
)
const EMPTY_BODY = ""

var NOT_EXISTING_ARTICLE_ID = strconv.Itoa(-1)
var GIN_ENGINE *gin.Engine

func TestMain(m *testing.M) {
	client := Repositories.InitRedis(TEST_DATABASE)
	defer client.Close()

	Repositories.Initialize(client)
	client.FlushDB()

	GIN_ENGINE = gin.Default()
	InitController(GIN_ENGINE)

	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestGetArticle(t *testing.T) {
	articleID, err := createArticle(SAMPLE_TITLE)
	if err != nil {
		assert.Fail(t, "Can't create article", err)
		return
	}
	defer deleteArticle(articleID)

	path := ArticleController.Path + "/" + articleID
	w := requestArticle(request{
		method: Models.GET,
		path:   path,
	})

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Header().Get("Content-Type"), "application/json")

	article, err := unmarshalResponseArticle(w.Body)
	assert.NoError(t, err, "can't decode response body to article")
	assert.Equal(t, SAMPLE_TITLE, article.Title)
}

func TestGetArticleWithNotExistsArticleID(t *testing.T) {
	path := ArticleController.Path + "/" + NOT_EXISTING_ARTICLE_ID
	w := requestArticle(request{
		method: Models.GET,
		path:   path,
	})

	assert.Equal(t, 400, w.Code)
}

func TestGetArticleWithInvalidAticle(t *testing.T) {
	id, err := createInvalidArticle()
	if err != nil {
		assert.Fail(t, "can't create article", err)
		return
	}
	defer deleteArticle(id)

	path := ArticleController.Path + "/" + id
	w := requestArticle(request{
		method: Models.GET,
		path:   path,
	})

	assert.Equal(t, 500, w.Code)
}

// TODO: add user to sampleArticle.
func TestPostArticle(t *testing.T) {
	var actual, expected *Models.Article
	defer func() {
		deleteArticle(strconv.Itoa(actual.ID))
	}()

	var sampleArticle = &Models.Article{
		Title: "test",
	}
	path := ArticleController.Path
	w := requestArticle(request{
		method:  Models.POST,
		path:    path,
		article: sampleArticle,
	})

	assert.Equal(t, 200, w.Code)

	actual, err := unmarshalResponseArticle(w.Body)
	assert.NoError(t, err, "parse response")

	expected, err = Repositories.Article.READ(strconv.Itoa(actual.ID))
	assert.NoError(t, err, "reading real article")
	assert.Equal(t, expected, actual)
}

func TestPostArticleWithIncorrectContentType(t *testing.T) {
	path := ArticleController.Path
	w := requestArticle(request{
		method:      Models.POST,
		path:        path,
		contentType: "plain/text",
	})

	assert.Equal(t, 400, w.Code)
}

func TestPostArticleWithInvalidBody(t *testing.T) {
	path := ArticleController.Path
	w := requestArticle(request{
		method: Models.POST,
		path:   path,
		body:   "invalid json",
	})

	assert.Equal(t, 400, w.Code)
}

func TestDeleteArticle(t *testing.T) {
	id, err := createArticle(SAMPLE_TITLE)
	if err != nil {
		assert.Fail(t, "can't create article", err)
		return
	}
	defer deleteArticle(id)

	path := ArticleController.Path + "/" + id
	w := requestArticle(request{
		method: Models.DELETE,
		path:   path,
	})

	assert.Equal(t, 200, w.Code)
	assert.False(t, Repositories.Article.EXISTS(id))
}

func TestDeleteArticleWithNotExistsArticleID(t *testing.T) {
	deleteArticle(NOT_EXISTING_ARTICLE_ID)

	path := ArticleController.Path + "/" + NOT_EXISTING_ARTICLE_ID
	w := requestArticle(request{
		method: Models.DELETE,
		path:   path,
	})

	assert.Equal(t, 400, w.Code)
}

func TestPutArticle(t *testing.T) {
	id, err := createArticle(SAMPLE_TITLE)
	if err != nil {
		assert.Fail(t, "Can't create article", err)
	}

	path := ArticleController.Path + "/" + id
	updatedTitle := SAMPLE_TITLE + "_POSTFIX"
	newArticle := &Models.Article{Title: updatedTitle}
	w := requestArticle(request{
		method:  Models.PUT,
		path:    path,
		article: newArticle,
	})

	assert.Equal(t, 200, w.Code)

	actual, err := Repositories.Article.READ(id)
	assert.NoError(t, err)
	assert.Equal(t, updatedTitle, actual.Title)
}

func TestPutArticleWithNotExistsArticleId(t *testing.T) {
	id := NOT_EXISTING_ARTICLE_ID

	path := ArticleController.Path + "/" + id
	updatedTitle := SAMPLE_TITLE + "_POSTFIX"
	newArticle := &Models.Article{Title: updatedTitle}
	w := requestArticle(request{
		method:  Models.PUT,
		path:    path,
		article: newArticle,
	})

	assert.Equal(t, 400, w.Code)
}

func TestPutArticleWithInvalidBody(t *testing.T) {
	id, err := createArticle(SAMPLE_TITLE)
	if err != nil {
		assert.Fail(t, "Can't create article", err)
		return
	}

	path := ArticleController.Path + "/" + id
	w := requestArticle(request{
		method: Models.PUT,
		path:   path,
		body:   "invalid json string",
	})

	assert.Equal(t, 400, w.Code)
}

func createArticle(title string) (string, error) {
	article := &Models.Article{
		Title: title,
	}

	return Repositories.Article.CREATE(article)
}

func deleteArticle(id string) {
	Repositories.Article.DELETE(id)
}

// TODO: i forgot to increase counter, and it breaks redis server.
// maybe i need to make redis.Set() which set and increase counter same time.
// and redis.Del() which delete and decrease counter
func createInvalidArticle() (string, error) {
	id := "test"
	redisKey := Repositories.CreateKey(Repositories.Article.Prefix, id)
	err := Repositories.Client.Set(redisKey, "invalid json string", 0).Err()
	Repositories.Client.Incr(Repositories.Article.Prefix + ":counter")

	return id, err
}

func unmarshalResponseArticle(body *bytes.Buffer) (*Models.Article, error) {
	var article Models.Article

	err := json.Unmarshal(body.Bytes(), &article)
	return &article, err
}

func requestArticle(r request) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(string(r.method), r.path, r.getBodyReader())
	req.Header.Add("Content-Type", r.getContentType())

	w := httptest.NewRecorder()
	GIN_ENGINE.ServeHTTP(w, req)

	return w
}
