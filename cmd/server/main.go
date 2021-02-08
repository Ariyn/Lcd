package main

import (
	"github.com/ariyn/Lcd/lcd"
	"github.com/ariyn/Lcd/lcd/controller"
	"github.com/ariyn/Lcd/util"
	"github.com/labstack/echo"
	"log"
)

func main() {
	conf := Load("config.json")
	db := util.MustConnectDB(conf.DB)

	userController := controller.NewUser(lcd.NewUserRepository(db))
	articleController := controller.NewArticle(lcd.NewArticleRepository(db))

	e := echo.New()
	userController.InitHandlers(e)
	articleController.InitHandlers(e)

	if err := e.Start(":8080"); err != nil {
		log.Println(err)
	}
}
