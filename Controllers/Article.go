package Controllers

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/ariyn/Lcd/Models"
	"github.com/ariyn/Lcd/Models/Errors"
	"github.com/ariyn/Lcd/Repositories"
	"github.com/gin-gonic/gin"
)

var ArticleController Controller = Controller{
	Path: "/article",
	Handlers: []Handler{
		Handler{Path: "", Method: Models.POST, Handler: postArticle, UseAuth: true},
		Handler{Path: "/:articleID", Method: Models.GET, Handler: getArticle},
		Handler{Path: "/:articleID", Method: Models.DELETE, Handler: deleteArticle, UseAuth: true},
		Handler{Path: "/:articleID", Method: Models.PUT, Handler: putArticle, UseAuth: true},
	},
}

// TODO: GetArticle? getArticle? ArticleController.GetArticle?
func getArticle(c *gin.Context) {
	articleID := c.Param("articleID")
	log.Printf("[INFO]controller reading article Id '%s'\n", articleID)

	article, err := Repositories.Article.READ(articleID)
	if err != nil {
		if _, ok := err.(Errors.InvalidJson); ok {
			panic(&HTTPError{internalServerError, c.Request.URL.Path, "Internal Server Error", err})
		}
		if _, ok := err.(Errors.NoSuchArticle); ok {
			panic(&HTTPError{badRequest, c.Request.URL.Path, "No such article", err})
		}
		panic(&HTTPError{internalServerError, c.Request.URL.Path, "Unknown error", err})
	}

	c.JSON(200, article)
}

func postArticle(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		panic(&HTTPError{badRequest, c.Request.URL.Path, "Can't read body", err})
	}
	if len(body) == 0 {
		panic(&HTTPError{badRequest, c.Request.URL.Path, "Empty body", nil})
	}

	var article Models.Article
	err = json.Unmarshal(body, &article)
	if err != nil {
		panic(&HTTPError{badRequest, c.Request.URL.Path, "Can't parse body", err})
	}

	_, err = Repositories.Article.CREATE(&article)
	if err != nil {
		if _, ok := err.(Errors.RedisFailure); ok {
			panic(&HTTPError{internalServerError, c.Request.URL.Path, "Internal Server error", err})
		}
	}

	c.JSON(200, article)
}

func deleteArticle(c *gin.Context) {
	articleID := c.Param("articleID")
	log.Printf("[INFO]controller deleting article Id '%s'\n", articleID)

	if Repositories.Article.EXISTS(articleID) == false {
		panic(&HTTPError{badRequest, c.Request.URL.Path, "No such article", nil})
	}

	err := Repositories.Article.DELETE(articleID)
	if err != nil {
		if _, ok := err.(Errors.RedisFailure); ok {
			panic(&HTTPError{internalServerError, c.Request.URL.Path, "Internal Server error", err})
		}
	}

	c.JSON(200, struct{ status string }{status: "success"})
}

func putArticle(c *gin.Context) {
	articleID := c.Param("articleID")
	log.Printf("[INFO]controller updating article Id '%s'\n", articleID)

	body, err := c.GetRawData()
	if err != nil {
		panic(&HTTPError{badRequest, c.Request.URL.Path, "Can't read body", err})
	}

	var article Models.Article
	err = json.Unmarshal(body, &article)
	if err != nil {
		panic(&HTTPError{badRequest, c.Request.URL.Path, "Can't parse body", err})
	}

	article.ID, _ = strconv.Atoi(articleID)

	_, err = Repositories.Article.UPDATE(&article)
	if err != nil {
		if _, ok := err.(Errors.NoSuchArticle); ok {
			panic(&HTTPError{badRequest, c.Request.URL.Path, "No such article", err})
		}
	}

	updatedArticle, _ := Repositories.Article.READ(articleID)
	log.Printf("%#v\n", article)
	log.Printf("%#v\n", updatedArticle)

	c.JSON(200, article)
}
