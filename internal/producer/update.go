/*
  Copyright (C) 2019 - 2022 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package producer

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"

	"github.com/superhero-match/superhero-update/internal/producer/model"
)

// UpdateSuperhero publishes update for a Superhero on Kafka topic for it to be
// consumed by consumer and updated in DB and Elasticsearch.
func (p *producer) UpdateSuperhero(s model.Superhero) error {
	return p.updateSuperhero(p.Producer, s)
}

// publishUpdateSuperhero publishes update for a Superhero on Kafka topic.
func publishUpdateSuperhero(producer *kafka.Writer, s model.Superhero) error {
	err := s.Validate()
	if err != nil {
		return err
	}

	var sb bytes.Buffer

	err = json.NewEncoder(&sb).Encode(s)
	if err != nil {
		return err
	}

	err = producer.WriteMessages(context.Background(),
		kafka.Message{
			Value: sb.Bytes(),
		},
	)
	if err != nil {
		return err
	}

	return nil
}
