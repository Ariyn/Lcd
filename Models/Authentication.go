package Models

type Role string

type AuthRule struct {
	FullPath string
	Action   Action
	Role     Role
}

func (a AuthRule) IsAuthorizable(u *User) bool {
	if a.Role == u.Role {
		return true
	}

	return false
}