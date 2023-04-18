package models

type User struct {
	Login    string `bson:"_id"`
	Username string `bson:"username"`
	Password string `bson:"password"`
	err      string
}

func NewUser(login, password string) *User {
	return &User{
		Login:    login,
		Password: password,
	}
}

func (u User) Error() string {
	return u.err
}

var UserDoesntExist error = &User{err: "user doesn't exist"}
var UserAlreadyExist error = &User{err: "user already exist"}
