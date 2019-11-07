package Models

// AuthRule for controller handler
// each controller handler has at least one AuthRule.
type AuthRule struct {
	FullPath       string
	Method         Method
	AllowAnonymous bool
	Role           Role
}

// Role for users
type Role int

// Role is actual role enum
const (
	RoleAnonymous Role = iota + 1
	RoleUser
	RoleOwner
	RoleEditor
	RoleAdmin
)

// IsHighRole returns role is higher then other.
// higher role means much permission and authorized more.
func (r Role) IsHighRole(other Role) bool {
	return other < r
}
