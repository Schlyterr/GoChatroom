// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Sample helloworld is a basic App Engine flexible app.
package main

import (
	"bufio"
	"strings"

	"fmt"
	"log"
	"os"

	"github.com/GoogleCloudPlatform/golang-samples/appengine_flexible/chatroom/chatRoom"

	"cloud.google.com/go/pubsub"
	"golang.org/x/net/context"
)

func main() {
	ctx := context.Background()
	projectID := "liuchat-165710"
	reader := bufio.NewReader(os.Stdin)

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	fmt.Print("Pick a username: ")
	username, _ := reader.ReadString('\n')
	fmt.Print("Enter chatroom: ")
	chatroom, _ := reader.ReadString('\n')

	username = strings.TrimSuffix(username, "\n")
	chatroom = strings.TrimSuffix(chatroom, "\n")

	topic := chatRoom.Create(ctx, client, chatroom)

	for true {
		fmt.Print("Enter message: ")
		message, _ := reader.ReadString('\n')
		chatRoom.Message(ctx, topic, message, username)
	}
}
