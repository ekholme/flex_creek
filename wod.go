package flexcreek

import "context"

type Wod struct {
	ID          string `json:"id"`
	Name        string `json:"name" validate:"required,alphanum"`
	Type        string `json:"type" validate:"required"`
	Length      string `json:"length" validate:"required,oneof=Short Medium Long"`
	Description string `json:"description" validate:"required"`
	Difficulty  int64  `json:"difficulty" validate:"required,oneof=Easy Moderate Hard Death"`
	AddedBy     string `json:"addedby"` //this should end up being a User.Username
}

type WodService interface {
	CreateWod(ctx context.Context, w *Wod) error
	GetAllWods(ctx context.Context) ([]*Wod, error)
	GetWodsbyType(ctx context.Context, t string) ([]*Wod, error)
	GetWodByID(ctx context.Context, id string) (*Wod, error)
	GetRandomWod(ctx context.Context) (*Wod, error)
	UpdateWod(ctx context.Context, id string, w *Wod) (*Wod, error)
	DeleteWod(ctx context.Context, id string) error
}
