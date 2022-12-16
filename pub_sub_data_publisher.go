package pubsubmit

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/pubsub"
)

type PubSubDataPublisher[T any] struct {
	configuration *Configuration
}

func NewPubSubDataPublisher[T any](configuration *Configuration) *PubSubDataPublisher[T] {
	this := new(PubSubDataPublisher[T])

	this.configuration = configuration

	return this
}

func (this PubSubDataPublisher[T]) Publish(data T) {
	ctx := context.Background()

	client, _ := pubsub.NewClient(ctx, this.configuration.ProjectID)

	defer client.Close()

	marshalledData, _ := json.Marshal(data)

	t := client.Topic(this.configuration.CollectionName + "-" + this.configuration.Operation + "-submit")
	result := t.Publish(ctx, &pubsub.Message{
		Data: marshalledData,
	})

	_, _ = result.Get(ctx)
}
