package chatRoom

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

type Msg struct {
	Username string
	Message  string
}

func Create(ctx context.Context, client *pubsub.Client, name string) *pubsub.Topic {
	topic := client.Topic(name)
	ok, err := topic.Exists(ctx)
	if err != nil {
		log.Fatalf("Some topic error occured: %v", err)
	}
	if !ok {
		topic, err := client.CreateTopic(ctx, name)
		if err != nil {
			log.Fatalf("failed to create topic: %v", err)
		}
		return topic
	}
	return topic
}

func Join(ctx context.Context, client *pubsub.Client, name string, topic *pubsub.Topic) *pubsub.Subscription {
	sub := client.Subscription(name)
	ok, err := sub.Exists(ctx)
	if err != nil {
		log.Fatalf("Some sub error occured: %v", err)
	}
	if !ok {
		sub, err := client.CreateSubscription(ctx, name, topic, 0, nil)
		if err != nil {
			log.Fatalf("Failed to create subscription: %v", err)
		}
		return sub
	}
	return sub
}

func Message(ctx context.Context, topic *pubsub.Topic, msg string, user string) {

	data := Msg{
		Username: user,
		Message:  msg,
	}
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Print("Json Error")
	}
	topic.Publish(ctx, &pubsub.Message{Data: b})
	//topic.Stop()
}
