//go:build integrationtyped
// +build integrationtyped

package pubsubmit

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/stretchr/testify/assert"
)

var project = os.Getenv("PROJECT")
var appUrl = os.Getenv("APP_URL")
var collectionName = os.Getenv("COLLECTION_NAME")
var operation = os.Getenv("OPERATION")

func TestDataMatchingModelIsPublished(t *testing.T) {
	ctx := context.Background()
	client, _ := pubsub.NewClient(ctx, project)
	defer client.Close()

	topicID := fmt.Sprintf("%s-%s-submit", collectionName, operation)

	topic, _ := client.CreateTopic(ctx, topicID)

	subscription, _ := client.CreateSubscription(ctx, "test-subscription", pubsub.SubscriptionConfig{
		Topic:               topic,
		RetainAckedMessages: false,
	})

	body := make(map[string]interface{})
	body["prop1"] = "abc"
	body["prop2"] = 123
	body["prop3"] = "def"

	bodyJson, _ := json.Marshal(body)

	bodyBuffer := bytes.NewBuffer(bodyJson)

	http.Post(appUrl, "application/json", bodyBuffer)

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	messagesReceived := 0
	var messageDataBytes []byte

	subscription.Receive(ctx, func(_ context.Context, message *pubsub.Message) {
		messageDataBytes = message.Data
		messagesReceived++
		message.Ack()
	})

	assert.Equal(t, 1, messagesReceived)

	var messageData map[string]interface{}
	json.Unmarshal(messageDataBytes, &messageData)

	assert.Equal(t, "abc", messageData["prop1"])
	assert.Equal(t, float64(123), messageData["prop2"])
	assert.Equal(t, nil, messageData["prop3"])
}
