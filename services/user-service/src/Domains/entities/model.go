package entities

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID       string
	RoleId   int    `json:"roleId" db:"role_id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type RegisteredUser struct {
	ID   string
	Name string
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
