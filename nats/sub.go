package nats

import (
	"fmt"
	"log"
	"sync"

	"github.com/nats-io/nats.go"
)

func SubscribeMessages(wg *sync.WaitGroup, subject string) {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatalf("Error connecting to NATS: %v", err)
	}
	defer nc.Close()

	msgCount := 0
	sub, err := nc.Subscribe(subject, func(msg *nats.Msg) {
		msgCount++
		fmt.Printf("Received message %d\n", msgCount)
	})
	if err != nil {
		log.Fatalf("Error subscribing to subject: %v", err)
	}
	sub.SetPendingLimits(-1, -1)

	fmt.Println("Subscriber is ready and waiting for messages...")

	// Keep the subscriber running
	defer nc.Flush()
	select {}
}
