package frstr

import (
	"context"

	firestore "cloud.google.com/go/firestore/apiv1"
	flexcreek "github.com/ekholme/flex_creek"
)

type favoriteWodsService struct {
	client *firestore.Client
}

func NewFavoriteWodsService(c *firestore.Client) flexcreek.FavoriteWodsService {
	return &favoriteWodsService{
		client: c,
	}
}

//methods

func (fs favoriteWodsService) CreateFavoriteWod(ctx context.Context, userId string, w *flexcreek.Wod) error {
	return nil
}

func (fs favoriteWodsService) DeleteFavoriteWod(ctx context.Context, userId string, wodId string) error {
	return nil
}

func (fs favoriteWodsService) GetAllFavoriteWods(ctx context.Context, userId string) ([]*flexcreek.Wod, error) {
	return nil, nil
}
