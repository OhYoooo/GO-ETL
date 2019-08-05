package processors

import (
	"common-etl/models"
	"log"
	"time"
)

// Processor : void struct
type Processor struct{}

// NewProcessor : expose processor instance
func NewProcessor() *Processor {
	return &Processor{}
}

// ProcessData : Parse models to sample datastore entities
func (p *Processor) ProcessData(inputChan chan []models.PubsubModel, outputChan chan []models.DatastoreEntity) {
	defer close(outputChan)

	for pubsubModels := range inputChan {
		log.Printf("Fetched by Processor: %+v\n", pubsubModels)
		var entitys []models.DatastoreEntity

		for _, data := range pubsubModels {
			entity := models.DatastoreEntity{}

			// The layout is constructed from a fixed reference time: Mon Jan 2 15:04:05 -0700 MST 2006
			layout := "2006-01-02T03:04:05.000+00:00"
			timestamp, err := time.Parse(layout, data.Timestamp)
			if err != nil {
				log.Fatal(err)
			}

			entity.Timestamp = timestamp.UnixNano() / 1000000
			entity.UserID = data.UserID
			entity.Password = data.Password

			// Do some other transaction

			entitys = append(entitys, entity)
		}
		if len(entitys) > 0 {
			outputChan <- entitys
		}
	}
	log.Println("Processor is terminated")
}
