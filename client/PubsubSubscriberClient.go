package client

import (
	"context"
	"log"

	pubsub "cloud.google.com/go/pubsub/apiv1"
)

// PubsubSubscriberClient is a wrapper of gcp pubsub client
type PubsubSubscriberClient struct {
	Client *pubsub.SubscriberClient
}

var pubsubClient *PubsubSubscriberClient
var pubsubCtx context.Context

// GetPubsubSubscriberClient : get instance of pubsub client, and the client will only be initialized once
func GetPubsubSubscriberClient() (*PubsubSubscriberClient, context.Context) {
	if pubsubClient == nil || pubsubCtx == nil {
		ctx := context.Background()
		client, newClientErr := pubsub.NewSubscriberClient(ctx)
		if newClientErr != nil {
			log.Fatal(newClientErr)
		}
		log.Println("Created new subscriber client")

		pubsubClient = &PubsubSubscriberClient{client}
		pubsubCtx = ctx
	}
	return pubsubClient, pubsubCtx
}
