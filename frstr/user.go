package frstr

import (
	"context"

	"cloud.google.com/go/firestore"
	flexcreek "github.com/ekholme/flex_creek"
	"golang.org/x/crypto/bcrypt"
)

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
	err := hashPw(u)

	if err != nil {
		return err
	}
	_, _, err = us.Client.Collection(userColl).Add(ctx, u)

	if err != nil {
		return err
	}

	return nil
}

func (us userService) GetAllUsers(ctx context.Context) ([]*flexcreek.User, error) {

	docs, err := us.Client.Collection(userColl).Documents(ctx).GetAll()

	if err != nil {
		return nil, err
	}

	users := make([]*flexcreek.User, len(docs))

	for i, v := range docs {
		var user *flexcreek.User

		v.DataTo(&user)

		users[i] = user
	}

	return users, nil
}

func (us userService) GetUserByID(ctx context.Context, id string) (*flexcreek.User, error) {

	ref, err := getFirestoreDocId(us.Client.Collection(userColl), ctx, id)

	if err != nil {
		return nil, err
	}

	var user *flexcreek.User

	doc, err := us.Client.Collection(userColl).Doc(ref).Get(ctx)

	if err != nil {
		return nil, err
	}

	doc.DataTo(&user)

	return user, nil
}

func (us userService) UpdateUser(ctx context.Context, id string, u *flexcreek.User) (*flexcreek.User, error) {
	ref, err := getFirestoreDocId(us.Client.Collection(userColl), ctx, id)

	if err != nil {
		return nil, err
	}

	u.ID = id

	//i think this should work?
	_, err = us.Client.Collection(userColl).Doc(ref).Set(ctx, u)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (us userService) DeleteUser(ctx context.Context, id string) error {
	ref, err := getFirestoreDocId(us.Client.Collection(userColl), ctx, id)

	if err != nil {
		return err
	}

	_, err = us.Client.Collection(userColl).Doc(ref).Delete(ctx)

	if err != nil {
		return err
	}

	//I don't think this should return anything -- I can propogate a generic message via the handler if the error is nil
	return nil
}

// util function to hash a password
func hashPw(u *flexcreek.User) error {
	p, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(p)

	return nil
}
