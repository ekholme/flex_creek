package frstr

import (
	"context"

	flexcreek "github.com/ekholme/flex_creek"
)

const (
	projectID      = "flex-creek"
	userCollection = "users"
)

type userService struct{}

//generator function
func NewUserService() flexcreek.UserService {
	return &userService{}
}

//placeholders until i implement actual code
func (us *userService) CreateUser(ctx context.Context, u *flexcreek.User) error {
	return nil
}

func (us *userService) GetUser(ctx context.Context, id string) (*flexcreek.User, error) {
	return nil, nil
}

func (us *userService) UpdateUser(ctx context.Context, id string) (*flexcreek.User, error) {
	return nil, nil
}

func (us *userService) DeleteUser(ctx context.Context, id string) error {
	return nil
}
