package util

import (
	"database/sql"
	"github.com/ariyn/Lcd/config"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func MustConnectDB(conf config.DB) *sql.DB {
	if !strings.Contains(conf.ConnectionString, "loc=") {
		conf.ConnectionString += "&loc=Asia%2FSeoul"
	}

	db, err := sql.Open(conf.Driver, conf.ConnectionString)
	if err != nil {
		panic(err)
	}

	if conf.MaxIdle > 0 {
		db.SetMaxIdleConns(conf.MaxIdle)
	}
	if conf.MaxOpen > 0 {
		db.SetMaxOpenConns(conf.MaxOpen)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
