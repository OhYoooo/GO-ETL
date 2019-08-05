package writers

import (
	"common-etl/client"
	"common-etl/models"
	"log"

	"cloud.google.com/go/datastore"
)

// Writer : void struct
type Writer struct {
}

// NewWriter : Returns an instance of Writer
func NewWriter() *Writer {
	return &Writer{}
}

// Write : Write to google cloud Datastore
func (p *Writer) Write(outputChan chan []models.DatastoreEntity) {
	datastoreClient, ctx := client.GetDatastoreClient()
	defer datastoreClient.Client.Close()

	for entitys := range outputChan {
		log.Printf("Fetched by Writer: %+v\n", entitys)
		var keys []*datastore.Key
		for range entitys {
			taskKey := datastore.IncompleteKey("Test", nil)
			taskKey.Namespace = datastoreClient.Namespace
			keys = append(keys, taskKey)
		}

		tx, err := datastoreClient.Client.NewTransaction(ctx)
		if err != nil {
			log.Fatalf("Datastore transaction failed at creating new transaction: %v", err)
		}

		if _, err := tx.PutMulti(keys, entitys); err != nil {
			tx.Rollback()
			log.Fatalf("Datastore transaction failed at PutMulti: %v", err)
		}
		if _, err = tx.Commit(); err != nil {
			log.Fatalf("Datastore transaction failed at Commit: %v", err)
		}
		log.Println("Entities saved successfully")
	}
	log.Println("Writer is terminated")
}
