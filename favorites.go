package flexcreek

import "context"

type FavoriteService interface {
	CreateFavoriteWod(ctx context.Context, userId string, w *Wod) error
	DeleteFavoriteWod(ctx context.Context, userId string, wodId string) error
	GetAllFavoriteWods(ctx context.Context, userId string) ([]*Wod, error)
}
