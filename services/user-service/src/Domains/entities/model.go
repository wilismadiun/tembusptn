package entities

type User struct {
	ID       string
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
