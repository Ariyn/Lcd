package Controllers

import (
	"log"
	"reflect"
	"runtime"
	"strings"

	"github.com/ariyn/Lcd/Models"
	"github.com/ariyn/Lcd/Service/middleware"
	"github.com/gin-gonic/gin"
)

type paging struct {
	Start int `form:"start, default=1"`
	Size  int `form:"size, default=30"`
}

type success struct {
	code    int
	message string
}

type HTTPErrorHandler func(*gin.Context, string, string, error)
type HTTPError struct {
	handler HTTPErrorHandler
	path    string
	message string
	rawErr  error
}

type Handler struct {
	Path    string
	Method  Models.Method
	Handler gin.HandlerFunc
	UseAuth bool
	Roles   string
}

// GetRoles return splited roles
func (h Handler) GetRoles() []Models.Role {
	roles := []Models.Role{}

	for _, roleName := range strings.Split(h.Roles, ",") {
		var role Models.Role
		roleName := strings.TrimSpace(roleName)

		switch roleName {
		case "User":
			role = Models.RoleUser
		case "Editor":
			role = Models.RoleEditor
		case "Admin":
			role = Models.RoleEditor
		}
		roles = append(roles, role)
	}

	if len(roles) == 0 {
		roles = append(roles, Models.RoleAnonymous)
	}

	return roles
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

			methodName := runtime.FuncForPC(reflect.ValueOf(method).Pointer()).Name()
			log.Printf("%s, %s, %#v", controller.Path+handler.Path, methodName, handler)

			path := handler.Path
			method(path, preHandler(handler))

			middleware.AddAuthRules(&Models.AuthRule{
				FullPath:       controller.Path + handler.Path,
				Method:         handler.Method,
				AllowAnonymous: !handler.UseAuth,
			})
			// log.Println(controller.Path+handler.Path, handler.Method, !handler.UseAuth)
		}
	}
}

// preHandler will handle request first.
// this method will parse request body, urlParameters, etc...
func preHandler(h Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			httpError := recover()
			if httpError != nil {
				err := httpError.(*HTTPError)
				log.Printf("[Error]controller %s %s: %s: %#v", string(h.Method), err.path, err.message, err.rawErr)
				err.handler(c, err.path, err.message, err.rawErr)
			}
		}()

		user := getUserOrAnonymous(c)
		if ahtorizeUserRole(user, &h) {
			h.Handler(c)
		} else {
			panic(&HTTPError{forbidden, c.Request.URL.Path, "Forbidden action", nil})
		}
	}
}

func getUserOrAnonymous(c *gin.Context) *Models.User {
	var user *Models.User

	tempUser, ok := c.Get("user")
	if !ok {
		user = &Models.User{
			Role: Models.RoleAnonymous,
		}
	} else {
		user, ok = tempUser.(*Models.User)
		if !ok {
			panic(&HTTPError{internalServerError, c.Request.URL.Path, "Can't load user authentication", nil})
		}
	}

	return user
}

func ahtorizeUserRole(user *Models.User, h *Handler) bool {
	if !h.UseAuth {
		return true
	}

	minRole := Models.RoleAdmin
	for _, role := range h.GetRoles() {
		if minRole.IsHighRole(role) {
			minRole = role
		}
	}
	return user.Role.IsHighRole(minRole)
}

func badRequest(c *gin.Context, path, message string, err error) {
	c.JSON(400, gin.H{
		"status":  "error",
		"message": "WrongRequest",
	})
}

func forbidden(c *gin.Context, path, message string, err error) {
	c.JSON(403, gin.H{
		"status":  "error",
		"message": "FrobiddenAction",
	})
}

func internalServerError(c *gin.Context, path, message string, err error) {
	c.JSON(500, gin.H{
		"status":  "error",
		"message": "InternalServerError",
	})
}
