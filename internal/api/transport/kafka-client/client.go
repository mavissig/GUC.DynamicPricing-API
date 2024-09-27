package kafka_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/mavissig/GUC.DynamicPricing-API/internal/api/domain"
	"github.com/mavissig/GUC.DynamicPricing-API/internal/api/transport"
	"log"
)

type Client struct {
	cfg *transport.Config
	p   *kafka.Producer
}

func New(cfg *transport.Config) (*Client, error) {
	p, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers": "kafka:9092",
		})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("[API][TRANSPORT][KAFKA-CLIENT][NEW][ERROR]: %s", err))
	}

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					log.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	return &Client{
		cfg: cfg,
		p:   p,
	}, nil
}

func (c *Client) AddData(data *domain.Data) error {
	topic := "outData"

	b, err := json.Marshal(data)
	if err != nil {
		log.Println("[API][TRANSPORT][KAFKA-CLIENT][AddData][ERROR]: ", err)
		return err
	}

	err = c.p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: b,
	}, nil)
	if err != nil {
		log.Println("[API][TRANSPORT][KAFKA-CLIENT][AddData][ERROR]: ", err)
		return err
	}

	return nil
}
