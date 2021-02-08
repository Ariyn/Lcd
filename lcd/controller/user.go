package controller

import (
	"github.com/ariyn/Lcd/lcd"
	"github.com/ariyn/Lcd/util"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"reflect"
)

type User struct {
	uRepo lcd.UserRepository
}

func NewUser(uRepo lcd.UserRepository) (c User) {
	c = User{
		uRepo: uRepo,
	}

	return c
}

func (c User) InitHandlers(e *echo.Echo) {
	userGroup := e.Group("/users", util.DefaultContentType("application/json"), util.ErrorLogger)
	userGroup.GET("/:userId", c.GetUser, util.ParseParam("userId", reflect.Int64))
	userGroup.POST("", c.CreateUser, c.ParseUser)
}

func (c User) GetUser(ctx echo.Context) (err error) {
	userId, ok := ctx.Get("userId").(int64)
	if !ok {
		return ctx.String(http.StatusBadRequest, "Bad User Id")
	}

	user, err := c.uRepo.GetUserByUid(userId)
	if err == util.NoResultErr {
		return ctx.String(404, "")
	}

	if err != nil {
		return err
	}

	err = ctx.JSON(http.StatusOK, user)
	return err
}

func (c User) CreateUser(ctx echo.Context) (err error) {
	user, ok := ctx.Get("user").(lcd.User)
	log.Println(user)
	if !ok {
		return ctx.String(400, "Bad Request")
	}

	createdUser, err := c.uRepo.CreateUser(user.Id, user.Name)
	if err != nil {
		log.Println(createdUser, err)
		return err
	}

	return ctx.JSON(200, createdUser)
}

func (c User) ParseUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var user User
		err := ctx.Bind(&user)
		if err != nil {
			return ctx.String(http.StatusBadRequest, "Bad Request")
		}

		ctx.Set("user", user)
		return next(ctx)
	}
}
