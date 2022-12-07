package flexcreek

import "context"

// other fields to include?
type User struct {
	ID           string `json:"id"`
	Username     string `json:"username" binding:"required"`
	Password     string `json:"password" binding:"required"`
	Email        string `json:"email" validate:"email" binding:"required"`
	FavoriteWods []*Wod `json:"favoriteWods"`
}

type UserService interface {
	CreateUser(ctx context.Context, u *User) error
	GetAllUsers(ctx context.Context) ([]*User, error)
	GetUserByID(ctx context.Context, id string) (*User, error)
	UpdateUser(ctx context.Context, id string) (*User, error)
	DeleteUser(ctx context.Context, id string) error
}
