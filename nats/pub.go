package nats

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

const (
	natsURL     = "nats://localhost:4222" // NATS server URL
	subject     = "latency.test"         // Subject to publish/subscribe
	numMessages = 100000                 // Number of messages to publish/subscribe
)

func PublishMessages(subject string) {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatalf("Error connecting to NATS: %v", err)
	}
	defer nc.Close()

	start := time.Now()
	payload := []byte(start.String())

	for i := 0; i < numMessages; i++ {
		if err := nc.Publish(subject, payload); err != nil {
			log.Fatalf("Error publishing message: %v", err)
		}
	}

	defer nc.Flush()

	fmt.Printf("Published %d messages in %v\n", numMessages, time.Since(start))
}
