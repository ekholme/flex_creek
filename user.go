package flexcreek

import "context"

// other fields to include?
type User struct {
	ID       string `json:"id"`
	Username string `json:"username" binding:"required" validate:"required,alphanum"`
	Password string `json:"password" binding:"required" validate:"required,alphanum"`
	Email    string `json:"email" validate:"required,email" binding:"required"`
	Admin    bool   `json:"admin"`
}

type UserService interface {
	CreateUser(ctx context.Context, u *User) error
	CheckUserExists(ctx context.Context, un string) error
	GetAllUsers(ctx context.Context) ([]*User, error)
	GetUserByID(ctx context.Context, id string) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	UpdateUser(ctx context.Context, id string, u *User) (*User, error)
	DeleteUser(ctx context.Context, id string) error
}

type Login struct {
	Username string `json:"username" binding:"required" validate:"required,alphanum"`
	Password string `json:"password" binding:"required" validate:"required,alphanum"`
}
