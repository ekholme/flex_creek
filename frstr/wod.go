package frstr

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"cloud.google.com/go/firestore"
	flexcreek "github.com/ekholme/flex_creek"
	"github.com/google/uuid"
)

//define some constants
const (
	wodCollection = "wods"
)

type wodService struct{}

//generator function
func NewWodService() flexcreek.WodService {
	return &wodService{}
}

//placeholders until I can actually implement CRUD operations

func (ws *wodService) CreateWod(ctx context.Context, w *flexcreek.Wod) error {

	client, err := firestore.NewClient(ctx, projectID)

	if err != nil {
		return err
	}

	defer client.Close()

	_, _, err = client.Collection(wodCollection).Add(ctx, w)

	if err != nil {
		return err
	}

	return nil
}

func (ws *wodService) GetAllWods(ctx context.Context) ([]*flexcreek.Wod, error) {

	client, err := firestore.NewClient(ctx, projectID)

	if err != nil {
		return nil, err
	}

	defer client.Close()

	docs, err := client.Collection(wodCollection).Documents(ctx).GetAll()

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

func (ws *wodService) GetRandomWod(ctx context.Context) (*flexcreek.Wod, error) {

	//set the random seed
	rand.Seed(time.Now().UnixNano())

	//get a random uuid to use as a reference
	ref := uuid.New().String()

	client, err := firestore.NewClient(ctx, projectID)

	if err != nil {
		return nil, err
	}

	defer client.Close()

	//get a random document using the reference ID
	docs, err := client.Collection(wodCollection).Where("ID", ">=", ref).OrderBy("ID", 1).Limit(1).Documents(ctx).GetAll()

	if err != nil {
		return nil, err
	}

	//try from other direction if previous result returned nothing
	if len(docs) == 0 {
		docs, err = client.Collection(wodCollection).Where("ID", "<", ref).OrderBy("ID", 1).Limit(1).Documents(ctx).GetAll()

		if err != nil {
			return nil, err
		}
	}

	//return an error if we can't retrieve any wods
	if len(docs) == 0 {

		err = errors.New("unable to retrieve any wods")

		return nil, err
	}

	var w *flexcreek.Wod

	//the above should only return 1 element, but will still return as a slice, so just get the first element
	docs[0].DataTo(&w)

	return w, nil
}

func (ws *wodService) GetWodByID(ctx context.Context, id string) (*flexcreek.Wod, error) {

	client, err := firestore.NewClient(ctx, projectID)

	if err != nil {
		return nil, err
	}

	defer client.Close()

	docs, err := client.Collection(wodCollection).Where("ID", "==", id).Limit(1).Documents(ctx).GetAll()

	if err != nil {
		return nil, err
	}

	if len(docs) == 0 {

		err = errors.New("unable to retrieve any wods")

		return nil, err
	}

	var w *flexcreek.Wod

	docs[0].DataTo(&w)

	return w, nil
}

//placeholder for now
func (ws *wodService) UpdateWod(ctx context.Context, id string) (*flexcreek.Wod, error) {
	return nil, nil
}

//placeholder for now
func (ws *wodService) DeleteWod(ctx context.Context, id string) error {
	return nil
}
