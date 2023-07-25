package frstr

import (
	"context"
	"errors"

	firestore "cloud.google.com/go/firestore"
	flexcreek "github.com/ekholme/flex_creek"
)

const favColl = "FavoriteWods"

type favoriteService struct {
	Client *firestore.Client
}

func NewFavoriteService(client *firestore.Client) flexcreek.FavoriteService {
	return &favoriteService{
		Client: client,
	}
}

//methods

func (fs favoriteService) CreateFavoriteWod(ctx context.Context, userId string, wod *flexcreek.Wod) error {

	ref, err := getFirestoreDocId(fs.Client.Collection(userColl), ctx, userId)

	if err != nil {
		return err
	}

	if err = checkAlreadyFavorite(ctx, wod.ID, fs, ref); err != nil {
		return err
	}

	_, _, err = fs.Client.Collection(userColl).Doc(ref).Collection(favColl).Add(ctx, wod)

	if err != nil {
		return err
	}

	return nil
}

func (fs favoriteService) DeleteFavoriteWod(ctx context.Context, userId string, wodId string) error {

	//i feel like there should be a more straightforward way to do this?
	uref, err := getFirestoreDocId(fs.Client.Collection(userColl), ctx, userId)

	if err != nil {
		return err
	}

	wref, err := getFirestoreDocId(fs.Client.Collection(userColl).Doc(uref).Collection(favColl), ctx, wodId)

	if err != nil {
		return err
	}

	_, err = fs.Client.Collection(userColl).Doc(uref).Collection(favColl).Doc(wref).Delete(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (fs favoriteService) GetAllFavoriteWods(ctx context.Context, userId string) ([]*flexcreek.Wod, error) {

	uref, err := getFirestoreDocId(fs.Client.Collection(userColl), ctx, userId)

	if err != nil {
		return nil, err
	}

	docs, err := fs.Client.Collection(userColl).Doc(uref).Collection(favColl).Documents(ctx).GetAll()

	if err != nil {
		return nil, err
	}

	wods := make([]*flexcreek.Wod, len(docs))

	for i, doc := range docs {
		var wod *flexcreek.Wod

		doc.DataTo(&wod)

		wods[i] = wod
	}

	return wods, nil

}

// utility to check if a wod is already favorited
func checkAlreadyFavorite(ctx context.Context, wid string, fs favoriteService, uref string) error {

	docs, err := fs.Client.Collection(userColl).Doc(uref).Collection(favColl).Where("ID", "==", wid).Documents(ctx).GetAll()

	if err != nil {
		return err
	}

	if len(docs) > 0 {
		return errors.New("wod already favorited")
	}

	return nil
}
