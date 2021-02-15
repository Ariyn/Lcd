package util

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

func GetString(ctx echo.Context, key string) (value string) {
	v := ctx.Get(key)
	if v == nil {
		return
	}

	value, _ = v.(string)
	return
}

func DefaultContentType(ct string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			header := ctx.Request().Header
			requestedContentType := header.Get("Content-Type")
			if requestedContentType == "" {
				header.Set("Content-Type", ct)
			}

			return next(ctx)
		}
	}
}

func ParseParam(key string, kind reflect.Kind, emptyOk bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			value := ctx.Param(key)
			if value == "" && !emptyOk {
				return next(ctx)
			}

			switch kind {
			case reflect.Int64:
				v, err := strconv.Atoi(value)
				if err != nil {
					return ctx.String(http.StatusBadRequest, "Bad User Id")
				}

				ctx.Set(key, int64(v))
			case reflect.Int:
				v, err := strconv.Atoi(value)
				if err != nil && value != "" {
					return ctx.String(http.StatusBadRequest, "Bad Request")
				}

				ctx.Set(key, v)
			}

			return next(ctx)
		}
	}
}

func ErrorLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		err = next(ctx)
		if err != nil {
			log.Println(err)
		}

		return err
	}
}
