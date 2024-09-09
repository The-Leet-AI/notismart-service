package queue

import (
    "context"
    "log"
    "github.com/segmentio/kafka-go"
)

func ConsumeNotifications(ctx context.Context, topic string, brokers []string) {
    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers: brokers,
        Topic:   topic,
        GroupID: "notification-group",
    })

    for {
        msg, err := r.ReadMessage(ctx)
        if err != nil {
            log.Printf("Error reading message: %v", err)
            continue
        }

        // Handle the notification message
        log.Printf("Received notification: %s", string(msg.Value))
    }
}
