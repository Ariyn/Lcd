package controller

import (
	"github.com/ariyn/Lcd/lcd"
	"github.com/ariyn/Lcd/util"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"reflect"
)

type Article struct {
	aRepo lcd.ArticleRepository
}

func NewArticle(aRepo lcd.ArticleRepository) (c Article) {
	c = Article{
		aRepo: aRepo,
	}

	return c
}

func (a Article) InitHandlers(e *echo.Echo) {
	userGroup := e.Group("/articles", util.DefaultContentType("application/json"), util.ErrorLogger)
	userGroup.GET("/:articleId", a.GetArticle, util.ParseParam("articleId", reflect.Int64))
	userGroup.POST("", a.CreateArticle, a.ParseArticle)
}

func (a Article) GetArticle(ctx echo.Context) (err error) {
	articleId, ok := ctx.Get("articleId").(int64)
	if !ok {
		return ctx.String(http.StatusBadRequest, "Bad Article Id")
	}

	article, err := a.aRepo.GetArticleByUid(articleId)
	if err == util.NoResultErr {
		return ctx.String(404, "")
	}

	if err != nil {
		return err
	}

	err = ctx.JSON(http.StatusOK, article)
	return err
}

func (a Article) CreateArticle(ctx echo.Context) (err error) {
	article, ok := ctx.Get("article").(lcd.Article)
	if !ok {
		return ctx.String(400, "Bad Request")
	}

	createdUser, err := a.aRepo.CreateArticle(article.Owner, article.Title)
	if err != nil {
		log.Println(createdUser, err)
		return err
	}

	return ctx.JSON(200, createdUser)
}

func (a Article) ParseArticle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// TODO: lcd.Article이 아니라, Fe에서 넘어올 ArticleDTO를 넣어줄 것.
		var article lcd.Article
		err := ctx.Bind(&article)
		if err != nil {
			return ctx.String(http.StatusBadRequest, "Bad Request")
		}

		ctx.Set("article", article)
		return next(ctx)
	}
}
