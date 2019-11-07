package Models

// User is model of user
type User struct {
	ID       int
	Account  string
	Password string
	Nickname string
	Role     Role
}

// MatchPassword checks User.Password matches with input password
// input password will be hashed in function.
func (u *User) MatchPassword(password string) bool {
	hashedPassword := hashPassword(password)
	return hashedPassword == u.Password
}

// hashPassword return hashed Password And it's not implemented YET!
// TODO: implement HMAC and PBKDF2
// https://d2.naver.com/helloworld/318732
// https://en.wikipedia.org/wiki/PBKDF2
// https://golang.org/pkg/hash/
// https://minwan1.github.io/2018/05/28/2018-05-28-HMAC/
// https://golang.org/pkg/crypto/sha256/
func hashPassword(p string) string {
	return p
}

func (u User) GetArticles() {

}

// IsAuthrizable compare user's role and AuthRule.
// it will return true if user is authorized.
func (u User) IsAuthrizable(roles []Role) bool {
	minRole := RoleAnonymous
	for _, role := range roles {
		if minRole.IsHighRole(role) {
			minRole = role
		}
	}

	return u.Role.IsHighRole(minRole)
}
