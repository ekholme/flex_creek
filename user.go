package flexcreek

import "context"

type User struct {
	ID        string   `json:"id"`
	Username  string   `json:"username"`
	Password  string   `json:"password"`
	Email     string   `json:"email"`
	Favorites []string `json:"favorites"`
}

type UserService interface {
	CreateUser(ctx context.Context, u *User) error
	GetUser(ctx context.Context, id string) (*User, error)
	UpdateUser(ctx context.Context, id string) (*User, error)
	DeleteUser(ctx context.Context, id string) error
}
