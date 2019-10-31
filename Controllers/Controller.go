package Controllers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

const (
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
	PUT    = "PUT"
)

type HTTPErrorHandler func(*gin.Context, string, string, error)
type HTTPError struct {
	handler HTTPErrorHandler
	path    string
	comment string
	err     error
}

type Handler struct {
	Path    string
	Method  string
	Handler gin.HandlerFunc
}

type Controller struct {
	Path     string
	Handlers []Handler
}

var controllers = []Controller{
	UserController,
}

func InitController(r *gin.Engine) {
	for _, controller := range controllers {
		for _, handler := range controller.Handlers {
			var method func(string, ...gin.HandlerFunc) gin.IRoutes

			switch handler.Method {
			case GET:
				method = r.GET
			case POST:
				method = r.POST
			case DELETE:
				method = r.DELETE
			case PUT:
				method = r.PUT
			}

			path := fmt.Sprintf("%s/%s", controller.Path, handler.Path)
			method(path, preHandler(handler.Handler))
		}
	}
}

// preHandler will handle request first.
// this method will parse request body, urlParameters, etc...
func preHandler(h gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			httpError := recover()
			if httpError != nil {
				err := httpError.(HTTPError)
				err.handler(c, err.path, err.comment, err.err)
			}
		}()
		h(c)
	}
}

// WARNING: this functions is not fully implemented and in wrong place.
// path: c.Request.URL.String()
func BadRequest(c *gin.Context, path, message string, err error) {
	log.Printf("[Error]controller %s: %s: %s", path, message, err)
	c.JSON(400, gin.H{
		"status":  "error",
		"message": "WrongRequest",
	})
}

func InternalServerError(c *gin.Context, path, message string, err error) {
	log.Printf("[Error]controller %s: %s: %s", path, message, err)
	c.JSON(500, gin.H{
		"status":  "error",
		"message": "InternalServerError",
	})
}
