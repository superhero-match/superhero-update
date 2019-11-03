package producer

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"

	"github.com/superhero-update/internal/producer/model"
)

// UpdateSuperhero publishes update for a Superhero on Kafka topic for it to be
// consumed by consumer and updated in DB and Elasticsearch.
func(p *Producer) UpdateSuperhero(s model.Superhero) error {
	var sb bytes.Buffer

	key := s.ID

	err := json.NewEncoder(&sb).Encode(s)
	if err != nil {
		return err
	}

	err = p.Producer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(key),
			Value: sb.Bytes(),
		},
	)
	if err != nil {
		return err
	}

	return nil
}
