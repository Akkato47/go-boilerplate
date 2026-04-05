package kafka

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(brokers []string, topic string) *Producer {
	w := &kafka.Writer{
		Addr:         kafka.TCP(brokers...),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireOne,
	}
	return &Producer{
		writer: w,
	}
}

func (p *Producer) Publish(ctx context.Context, key, value []byte) error {
	return p.writer.WriteMessages(ctx, kafka.Message{
		Key:   key,
		Value: value,
	})
}

func (p *Producer) PublishJSON(ctx context.Context, key []byte, value any) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return p.Publish(ctx, key, data)
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
