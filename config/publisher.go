package config

import "github.com/Kaleidoscope-Backup/healthrecord-repository/model"

// PublisherInterface ....
type PublisherInterface interface {
	Send(msg model.HealthRecordMessage) error
}

// Send ...
func Send(msg model.HealthRecordMessage) error {
	return PubAMQP.Send(msg)
}
