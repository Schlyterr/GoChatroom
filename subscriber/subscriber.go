package main

import (
	"bufio"
	"log"
	"os"

	"fmt"

	"strings"

	"encoding/json"

	"cloud.google.com/go/pubsub"
	"github.com/GoogleCloudPlatform/golang-samples/appengine_flexible/chatroom/chatRoom"
	"golang.org/x/net/context"
)

func main() {
	ctx := context.Background()
	projectID := "liuchat-165710"
	reader := bufio.NewReader(os.Stdin)

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed creating client: %v", err)
	}

	fmt.Print("Enter chat room: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")

	topic := chatRoom.Create(ctx, client, input)
	sub := chatRoom.Join(ctx, client, "Sub-"+input, topic)

	fmt.Print("Waiting for messages\n")

	err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		var msg chatRoom.Msg
		err := json.Unmarshal(m.Data, &msg)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		log.Printf("%v: %v", msg.Username, msg.Message)
		m.Ack()
	})
	if err != nil {
		log.Fatalf("Some error occured while receiving")
	}

}
