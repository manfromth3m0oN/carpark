package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/Shopify/sarama"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

func CreateConsumer() (sarama.Consumer, error) {
	log.Println("Creating kafka consumer")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	conn, err := sarama.NewConsumer([]string{"localhost:9092"}, config)

	if err != nil {
		return nil, err
	}

	return conn, err
}

func CreateProducer() (sarama.SyncProducer, error) {
	log.Println("Creating kafka producer")

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func PublishMessage(ctx context.Context, topic string, eventType string, data interface{}) error {
	if prodInter := ctx.Value("producer"); prodInter != nil {
		producer := prodInter.(sarama.SyncProducer)
		event := cloudevents.NewEvent()
		event.SetSource("server")
		event.SetType(eventType)
		event.SetData(cloudevents.ApplicationJSON, data)

		bytes, err := json.Marshal(event)
		if err != nil {
			return err
		}

		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(bytes),
		}

		partition, offset, err := producer.SendMessage(msg)

		if err != nil {
			return err
		}

		log.Printf("Message sent on topic %s with offset: %d and partition %d", topic, partition, offset)

		return nil
	}
	return errors.New("No kafka producer in context")
}
