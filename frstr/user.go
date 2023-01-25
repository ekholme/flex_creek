package frstr

import (
	"context"

	"cloud.google.com/go/firestore"
	flexcreek "github.com/ekholme/flex_creek"
)

//TODO
//implement UserService methods

const userColl = "users"

type userService struct {
	Client *firestore.Client
}

func NewUserService(client *firestore.Client) flexcreek.UserService {
	return &userService{
		Client: client,
	}
}

// add methods
func (us userService) CreateUser(ctx context.Context, u *flexcreek.User) error {
	_, _, err := us.Client.Collection(userColl).Add(ctx, u)

	if err != nil {
		return err
	}

	return nil
}

func (us userService) GetAllUsers(ctx context.Context) ([]*flexcreek.User, error) {

	var users []*flexcreek.User

	docs, err := us.Client.Collection(userColl).Documents(ctx).GetAll()

	if err != nil {
		return nil, err
	}

	for _, v := range docs {
		var user *flexcreek.User

		v.DataTo(&user)

		users = append(users, user)
	}

	return users, nil
}

func (us userService) GetUserByID(ctx context.Context, id string) (*flexcreek.User, error) {
	return nil, nil
}

func (us userService) UpdateUser(ctx context.Context, id string, u *flexcreek.User) (*flexcreek.User, error) {
	return nil, nil
}

func (us userService) DeleteUser(ctx context.Context, id string) error {
	return nil
}

func (us userService) GetUserFavoriteWods(ctx context.Context, id string) ([]*flexcreek.Wod, error) {
	return nil, nil
}

// this probably isn't correct
func (us userService) Login(ctx context.Context, username string, pw string) (*flexcreek.User, error) {
	return nil, nil
}
