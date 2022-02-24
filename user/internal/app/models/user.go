package models

type User struct {
	ID       string
	Login    string
	Password string
}

func NewUserSignIn(login, password string) *User {
	return &User{
		ID:       "",
		Login:    login,
		Password: password,
	}
}

func NewUserRegistration(login, password string) *User {
	return &User{
		ID:       "",
		Login:    login,
		Password: password,
	}
}
