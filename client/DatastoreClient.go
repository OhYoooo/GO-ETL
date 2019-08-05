package client

import (
	"common-etl/config"
	"context"
	"log"
	"os"

	"cloud.google.com/go/datastore"
)

// DatastoreClient : contains Client instance and Namespace
type DatastoreClient struct {
	Client    *datastore.Client
	Namespace string
}

var datastoreClient *DatastoreClient
var datastoreCtx context.Context

// GetDatastoreClient : get an common instance of datastore client, and the client will only be initialized once
func GetDatastoreClient() (*DatastoreClient, context.Context) {
	if datastoreClient == nil || datastoreCtx == nil {
		environment := os.Getenv("ENVIRONMENT")
		projectID := os.Getenv("GCLOUD_PROJECT_ID")
		if environment == "" || projectID == "" {
			log.Fatal("Environment variables are not set: ENVIRONMENT, GCLOUD_PROJECT_ID")
		}
		ctx := context.Background()
		client, err := datastore.NewClient(ctx, "")
		if err != nil {
			log.Fatal(err)
		}

		datastoreClient = &DatastoreClient{client, config.Env[environment]}
		datastoreCtx = ctx
		log.Println("Created new datastore client")
	}
	return datastoreClient, datastoreCtx
}
