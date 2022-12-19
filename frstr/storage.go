package frstr

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

const projectID = "flex-creek"

// just need a func to create a client, then this gets passed to services
// can initiate this in main
func NewClient() *firestore.Client {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, projectID)

	if err != nil {
		log.Fatal("Couldn't create firestore client")
	}

	return client
}
