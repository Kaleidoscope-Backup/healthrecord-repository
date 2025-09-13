package service

import (
	"encoding/json"

	"github.com/karte/healthrecord-repository/model"
)

// PostRecord ...
func PostRecord(record interface{}, consumerID string, recordID string) {
	msg := &model.HealthRecordMessage{}
	msg.Address = consumerID
	msg.RecordID = recordID
	e, err := json.Marshal(record)
	if err != nil {
		return
	}
	msg.Record = string(e)
	//go c.Send(*msg)
}
