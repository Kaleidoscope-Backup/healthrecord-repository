package config

import (
	"encoding/json"
	logging "log"

	"github.com/karte/healthrecord-repository/model"
	"github.com/streadway/amqp"
)

// AMQPPublisher ...
type AMQPPublisher struct {
}

// Send ...
func (p *AMQPPublisher) Send(msg model.HealthRecordMessage) error {
	conn, err := amqp.Dial(AMQPUri)
	if err != nil {
		logging.Printf("Error dialing AMQP: %s", err.Error())
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		logging.Printf("Error creating AMQP Channel: %s", err.Error())
		return err
	}

	q, err := ch.QueueDeclare(Cfg.AMQP.HealthRecordTopic, true, false, false, false, nil)
	if err != nil {
		logging.Printf("Error declaring AMQP Queue: %s", err.Error())
		return err
	}

	b, _ := json.Marshal(msg)
	logging.Printf("Message to send %s", string(b))
	if err := ch.Publish("", q.Name, false, false, amqp.Publishing{ContentType: "plain/text", Body: b}); err != nil {
		logging.Printf("Error Publishing meesage: %s", err.Error())
		return err
	}
	return nil
}
