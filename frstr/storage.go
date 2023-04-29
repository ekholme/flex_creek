package frstr

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

const projectID = "flex-creek"

// just need a func to create a client, then this gets passed to services
// can initiate this in main
func NewFirestoreClient() *firestore.Client {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, projectID)

	if err != nil {
		log.Fatal("Couldn't create firestore client")
	}

	return client
}

// helper func to get firestore doc id
func getFirestoreDocId(cr *firestore.CollectionRef, ctx context.Context, id string) (string, error) {
	iter := cr.Where("ID", "==", id).Limit(1).Documents(ctx)

	defer iter.Stop()

	doc, err := iter.Next()

	if err != nil {
		return "", err
	}

	//get the firestore id of the document
	ref := doc.Ref.ID

	return ref, nil

}
