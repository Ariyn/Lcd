package controller

import (
	"database/sql"
	errors2 "errors"
	"github.com/ariyn/Lcd/lcd"
	"github.com/ariyn/Lcd/models"
	"github.com/ariyn/Lcd/util"
	"github.com/labstack/echo"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"log"
	"net/http"
	"reflect"
)

var noRowsErr = errors2.New("sql: no rows in result set")

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) (c User) {
	c = User{
		db: db,
	}

	return c
}

func (c User) InitHandlers(e *echo.Echo) {
	userGroup := e.Group("/users", util.DefaultContentType("application/json"), util.ErrorLogger)
	userGroup.GET("/:userId", c.GetUser, util.ParseParam("userId", reflect.Int64, false))
	userGroup.POST("", c.CreateUser, c.ParseUser)
}

func (c User) GetUser(ctx echo.Context) (err error) {
	userId, ok := ctx.Get("userId").(int64)
	if !ok {
		return ctx.String(http.StatusBadRequest, "Bad User Id")
	}

	user, err := models.Users(Load(models.UserRels.OwnerUIDArticles), Where("uid = ?", userId)).OneG(ctx.Request().Context())
	if err != nil {
		if err.Error() == noRowsErr.Error() {
			return ctx.String(404, "No Such User")
		}
		return err
	}

	u := lcd.User{
		UID:  int64(user.UID),
		Id:   user.ID,
		Name: user.Name,
	}

	articles := make([]lcd.Article, 0)
	for _, a := range user.R.OwnerUIDArticles {
		articles = append(articles, lcd.Article{
			Uid:   int64(a.UID),
			Owner: u,
			Title: a.Title,
		})
	}

	u.Articles = articles

	err = ctx.JSON(http.StatusOK, u)
	return err
}

func (c User) CreateUser(ctx echo.Context) (err error) {
	user, ok := ctx.Get("user").(models.User)
	log.Println(user)
	if !ok {
		return ctx.String(400, "Bad Request")
	}

	err = user.InsertG(ctx.Request().Context(), boil.Infer())
	if err != nil {
		log.Println(user, err)
		return err
	}

	return ctx.JSON(200, user)
}

func (c User) ParseUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var user models.User
		err := ctx.Bind(&user)
		if err != nil {
			return ctx.String(http.StatusBadRequest, "Bad Request")
		}

		ctx.Set("user", user)
		return next(ctx)
	}
}
