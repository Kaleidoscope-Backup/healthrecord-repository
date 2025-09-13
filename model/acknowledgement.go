package model

import (
	"github.com/karte/healthrecord-repository/util"
	"github.com/karte/mongo-lib/models"
)

// AcknowledgementCreate ..
type AcknowledgementCreate struct {
	Created              util.Time `json:"created"`
	ConsumerID           string    `json:"consumerID"`
	RefrenceNotification string    `json:"refrenceNotification"`
	AckOption            string    `json:"ackOption"`
	Note                 *string   `json:"note"`
}

// Acknowledgement ..
type Acknowledgement struct {
	Id                   string       `json:"id" bson:"_id"`
	Created              util.Time    `json:"created" bson:"created"`
	ConsumerID           string       `json:"consumerID" bson:"consumerID"`
	RefrenceNotification string       `json:"refrenceNotification" bson:"refrenceNotification"`
	AckOption            string       `json:"ackOption" bson:"ackOption"`
	Note                 *string      `json:"note" bson:"note"`
	Meta                 *models.Meta //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
