package Controllers

import (
	"log"

	"github.com/ariyn/Lcd/Models"
	"github.com/gin-gonic/gin"
)

type HTTPErrorHandler func(*gin.Context, string, string, error)
type HTTPError struct {
	handler HTTPErrorHandler
	path    string
	message string
	rawErr  error
}

type Handler struct {
	Path    string
	Method  string
	Handler gin.HandlerFunc
	UseAuth bool
}

type Controller struct {
	Path     string
	Handlers []Handler
}

var controllers = []Controller{
	UserController,
	ArticleController,
}

func InitController(r *gin.Engine) {
	for _, controller := range controllers {
		group := r.Group(controller.Path)

		for _, handler := range controller.Handlers {
			var method func(string, ...gin.HandlerFunc) gin.IRoutes

			switch handler.Method {
			case Models.GET:
				method = group.GET
			case Models.POST:
				method = group.POST
			case Models.DELETE:
				method = group.DELETE
			case Models.PUT:
				method = group.PUT
			}

			path := handler.Path
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
				err := httpError.(*HTTPError)
				log.Printf("[Error]controller %s: %s: %#v", err.path, err.message, err.rawErr)
				err.handler(c, err.path, err.message, err.rawErr)
			}
		}()
		h(c)
	}
}

// WARNING: this functions is not fully implemented and in wrong place.
// path: c.Request.URL.String()
func badRequest(c *gin.Context, path, message string, err error) {
	c.JSON(400, gin.H{
		"status":  "error",
		"message": "WrongRequest",
	})
}

func internalServerError(c *gin.Context, path, message string, err error) {
	c.JSON(500, gin.H{
		"status":  "error",
		"message": "InternalServerError",
	})
}
