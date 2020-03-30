package lixinsso

// User struct define a user model. Everything can be done by the username and password.
type User struct {
	username string
	password string
}

// New function can create a new user object, with parameters username and password
func New(username string, password string) *User {
	user := new(User)
	user.username = username
	user.password = password

	return user
}

// Login to the sso.
func (user *User) Login() bool {
	auth := NewAuthorize("ufsso_longmeng_portal_index", Token, "https://sso.lixin.edu.cn/index.html")

	if auth.Authorize(user) {
		return true
	}

	return false
}
