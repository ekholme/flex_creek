package flexcreek

import "context"

type Wod struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Length      string `json:"length"`
	Description string `json:"description"`
	Difficulty  int64  `json:"difficulty" validate:"gte=1,lte=10"`
	AddedBy     string `json:"addedby"` //may want this to be a *User in the future
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
