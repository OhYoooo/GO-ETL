package readers

import (
	"common-etl/client"
	"common-etl/models"
	"encoding/json"
	"fmt"
	"log"
	"os"

	pubsubpb "google.golang.org/genproto/googleapis/pubsub/v1"
)

// Reader : struct contains a flag to notify reader goroutine to stop
type Reader struct {
	stopRead chan bool
}

// NewReader : get an instance of reader
func NewReader() *Reader {
	return &Reader{make(chan bool)}
}

// Extract data from google cloud pubsub
func (p *Reader) Extract(outputChan chan []models.PubsubModel) {

	projectID := os.Getenv("GCLOUD_PROJECT_ID")
	subscriptionName := os.Getenv("SUBSCRIPTION_NAME")
	logLocation := os.Getenv("LOG_LOCATION")

	logFile, openFileErr := os.OpenFile(logLocation, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if (projectID == "") || (subscriptionName == "") || (logLocation == "") {
		log.Fatal("Environment variables GCLOUD_PROJECT_ID, SUBSCRIPTION_NAME, LOG_LOCATION are not set")
	}

	if openFileErr != nil {
		log.Fatal(openFileErr)
	}
	log.Println("Opened log file")
	defer logFile.Close()

	pubsubClient, ctx := client.GetPubsubSubscriberClient()
	defer pubsubClient.Client.Close()

	subscriptionID := fmt.Sprintf("projects/%s/subscriptions/%s", projectID, subscriptionName)
	request := pubsubpb.PullRequest{
		Subscription: subscriptionID,
		MaxMessages:  5,
	}

	defer close(outputChan)

	for {
		select {
		case <-p.stopRead:
			log.Println("Reader is terminated")
			return
		default:
			response, err := pubsubClient.Client.Pull(ctx, &request)

			if err != nil {
				log.Fatal(err)
			}

			if len(response.ReceivedMessages) != 0 {

				var pubsubModels []models.PubsubModel

				for _, receivedMessage := range response.ReceivedMessages {
					log.Printf("Fetched from Pubsub: %s\n", string(receivedMessage.Message.Data))
					logFile.Write(append(receivedMessage.Message.Data, "\n"...))

					ackError := pubsubClient.Client.Acknowledge(ctx, &pubsubpb.AcknowledgeRequest{
						Subscription: subscriptionID,
						AckIds:       []string{receivedMessage.AckId},
					})
					if ackError != nil {
						log.Fatal(ackError)
					}

					var pubsubModel models.PubsubModel
					err := json.Unmarshal(receivedMessage.Message.Data, &pubsubModel)

					if err != nil {
						log.Fatal(err)
						continue
					}

					pubsubModels = append(pubsubModels, pubsubModel)
				}
				if len(pubsubModels) > 0 {
					outputChan <- pubsubModels
				}
			}
		}
	}
}

// Stop reader goroutine
func (p *Reader) Stop() {
	p.stopRead <- true
}
