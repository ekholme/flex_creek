package frstr

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"cloud.google.com/go/firestore"
	flexcreek "github.com/ekholme/flex_creek"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
)

const wodColl = "wods"

type wodService struct {
	Client *firestore.Client
}

func NewWodService(client *firestore.Client) flexcreek.WodService {
	return &wodService{
		Client: client,
	}
}

// implement methods
func (ws wodService) CreateWod(ctx context.Context, w *flexcreek.Wod) error {

	_, _, err := ws.Client.Collection(wodColl).Add(ctx, w)

	if err != nil {
		return err
	}

	return nil
}

func (ws wodService) GetAllWods(ctx context.Context) ([]*flexcreek.Wod, error) {
	docs, err := ws.Client.Collection(wodColl).Documents(ctx).GetAll()

	if err != nil {
		return nil, err
	}

	var wods []*flexcreek.Wod

	for _, doc := range docs {
		var w *flexcreek.Wod

		doc.DataTo(&w)

		wods = append(wods, w)
	}
	return wods, nil
}

func (ws wodService) GetWodByID(ctx context.Context, id string) (*flexcreek.Wod, error) {

	iter := ws.Client.Collection(wodColl).Limit(1).Documents(ctx)

	defer iter.Stop()

	doc, err := iter.Next()

	if err != nil {
		return nil, err
	}

	var wod *flexcreek.Wod

	doc.DataTo(&wod)

	return wod, nil
}

func (ws wodService) GetRandomWod(ctx context.Context) (*flexcreek.Wod, error) {
	rand.Seed(time.Now().UnixNano())

	ref := uuid.New().String()
	//try to find a random document 'less than' the random id we just created
	iter := ws.Client.Collection(wodColl).Where("ID", "<", ref).OrderBy("ID", 1).Limit(1).Documents(ctx)

	defer iter.Stop()

	doc, err := iter.Next()

	//try from other direction if nothing is available

	if err == iterator.Done {

		iter := ws.Client.Collection(wodColl).Where("ID", ">", ref).OrderBy("ID", 1).Limit(1).Documents(ctx)

		doc, err = iter.Next()

		if err == iterator.Done {
			return nil, errors.New("no wods available")
		}

	}

	var wod *flexcreek.Wod

	doc.DataTo(&wod)

	return wod, nil
}

// TODO
func (ws wodService) UpdateWod(ctx context.Context, id string) (*flexcreek.Wod, error) {
	return nil, nil
}

func (ws wodService) DeleteWod(ctx context.Context, id string) error {
	return nil
}
