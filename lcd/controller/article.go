package controller

import (
	"database/sql"
	"github.com/ariyn/Lcd/lcd"
	"github.com/ariyn/Lcd/models"
	"github.com/ariyn/Lcd/util"
	"github.com/labstack/echo"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"log"
	"net/http"
	"reflect"
)

type Article struct {
	db *sql.DB
}

func NewArticle(db *sql.DB) (c Article) {
	c = Article{
		db: db,
	}

	return c
}

func (a Article) InitHandlers(e *echo.Echo) {
	articleGroup := e.Group("/articles", util.DefaultContentType("application/json"), util.ErrorLogger)
	articleGroup.POST("", a.CreateArticle, a.ParseArticle)
	articleGroup.GET("/:articleId", a.GetArticle, util.ParseParam("articleId", reflect.Int64, false))
	articleGroup.POST("/:articleId/connection", a.ConnectTo, util.ParseParam("articleId", reflect.Int64, false))
}

func (a Article) GetArticle(ctx echo.Context) (err error) {
	articleId, ok := ctx.Get("articleId").(int64)
	if !ok {
		return ctx.String(http.StatusBadRequest, "Bad Article Id")
	}

	articleDto, err := models.Articles(qm.Where("uid = ?", articleId), qm.Load(models.ArticleRels.ToArticleMaps)).OneG(ctx.Request().Context())
	if err != nil {
		if err.Error() == noRowsErr.Error() {
			return ctx.String(404, "No Such Article")
		}
		return err
	}

	article := lcd.Article{
		Uid:   int64(articleDto.UID),
		Title: articleDto.Title,
	}

	articles := make([]lcd.Article, 0)
	for _, am := range articleDto.R.ToArticleMaps {
		articles = append(articles, lcd.Article{
			Uid:   int64(am.R.ToArticle.UID),
			Title: am.R.ToArticle.Title,
		})
	}

	article.ConnectedArticles = articles

	err = ctx.JSON(http.StatusOK, article)
	return err
}

func (a Article) CreateArticle(ctx echo.Context) (err error) {
	article, ok := ctx.Get("article").(models.Article)
	if !ok {
		return ctx.String(400, "Bad Request")
	}

	log.Println(article)

	err = article.InsertG(ctx.Request().Context(), boil.Infer())
	if err != nil {
		return err
	}

	return ctx.JSON(200, article)
}

type ConnectTo struct {
	ToUid int64
}

func (a Article) ConnectTo(ctx echo.Context) (err error) {
	articleId, ok := ctx.Get("articleId").(int64)
	if !ok {
		return ctx.String(http.StatusBadRequest, "Bad Article Id")
	}

	var ct ConnectTo
	err = ctx.Bind(&ct)
	if err != nil {
		return
	}

	articleFrom, err := models.Articles(qm.Where("uid = ?", articleId)).OneG(ctx.Request().Context())
	if err != nil {
		return
	}

	articleTo, err := models.Articles(qm.Where("uid = ?", ct.ToUid)).OneG(ctx.Request().Context())
	if err != nil {
		return
	}

	am := models.ArticleMap{
		From: articleFrom.UID,
		To:   articleTo.UID,
	}
	err = am.InsertG(ctx.Request().Context(), boil.Infer())
	if err != nil {
		return
	}

	return ctx.String(200, "ok")
}

func (a Article) ParseArticle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// TODO: lcd.Article이 아니라, Fe에서 넘어올 ArticleDTO를 넣어줄 것.
		var article lcd.Article
		err := ctx.Bind(&article)
		if err != nil {
			return ctx.String(http.StatusBadRequest, "Bad Request")
		}

		dto := models.Article{
			UID:      int(article.Uid),
			Title:    article.Title,
			OwnerUID: int(article.Owner.UID),
		}
		ctx.Set("article", dto)
		return next(ctx)
	}
}
