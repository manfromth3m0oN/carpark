package main

import (
	"context"
	"log"

	"github.com/manfromth3m0oN/carpark/db"
	"github.com/manfromth3m0oN/carpark/http"
	"github.com/manfromth3m0oN/carpark/kafka"
)

func main() {
	// TODO: Recive messages from a kafka topic and allocate space in the car park
	log.Println("Server starting up")

	ctx := context.Background()
	mongoClient, err := db.ConnectToDB(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	ctx = context.WithValue(context.Background(), "db", mongoClient)

	consumer, err := kafka.CreateConsumer()
	if err != nil {
		log.Fatalf("Failed to create kafka consumer: %v", err)
	}

	ctx = context.WithValue(ctx, "consumer", consumer)

	producer, err := kafka.CreateProducer()
	if err != nil {
		log.Fatalf("Failed to create kafka producer: %v", err)
	}

	ctx = context.WithValue(ctx, "producer", producer)

	http.CreateHttpServer()
}
