package lib

// User struct define.
type User struct {
	UserName    string            `json:"username"`
	Password    string            `json:"password"`
	Email       string            `json:"email"`
	Phone       string            `json:"phone"`
	Description string            `json:"description"`
	DisplayName string            `json:"displayName"`
	ExtraInfo   map[string]string `json:"extraInfo"`
}

// New a user.
func New() *User {
	return &User{}
}

func (u *User) create() {

}

func (u *User) assume(uuid string) {

}
