package Controllers

import (
	"log"
	"strings"

	"github.com/ariyn/Lcd/Models"
	"github.com/ariyn/Lcd/Service/middleware"
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
		case "Owner":
			role = Models.RoleOwner
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

			path := handler.Path
			method(path, preHandler(&handler))

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
func preHandler(h *Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			httpError := recover()
			if httpError != nil {
				err := httpError.(*HTTPError)
				log.Printf("[Error]controller %s: %s: %#v", err.path, err.message, err.rawErr)
				err.handler(c, err.path, err.message, err.rawErr)
			}
		}()

		checkUserRole(c, h)
		h.Handler(c)
	}
}

func checkUserRole(c *gin.Context, h *Handler) bool {
	if !h.UseAuth {
		return true
	}

	user, ok := c.Get("user")
	if !ok {
		url := c.Request.URL.Path
		panic(&HTTPError{internalServerError, url, "[Error]Can't find user", nil})
	}

	log.Println(user)
	h.GetRoles()
	return true
}

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
