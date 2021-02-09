package main

//go:generate sqlboiler --wipe mysql
import (
	"github.com/ariyn/Lcd/lcd/controller"
	"github.com/ariyn/Lcd/util"
	"github.com/labstack/echo"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"log"
)

func main() {
	conf := Load("config.json")
	db := util.MustConnectDB(conf.DB)

	boil.SetDB(db)

	userController := controller.NewUser(db)
	articleController := controller.NewArticle(db)

	e := echo.New()
	userController.InitHandlers(e)
	articleController.InitHandlers(e)

	if err := e.Start(":8080"); err != nil {
		log.Println(err)
	}
}
