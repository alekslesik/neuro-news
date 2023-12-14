package model


type User struct {
    ID       int
    Username string
    Email    string
}

type UserRepository interface {
    GetUserByID(id int) (*User, error)
}
