package model

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
	"github.com/Kaleidoscope-Backup/mongo-lib/models"
)

// OrderEventType ...
type OrderEventType string

const (
	//ORDER_PLACED_EVENT ..
	ORDER_PLACED_EVENT OrderEventType = "ORDER_PLACED_EVENT"

	//ORDER_SHIPPED_EVENT ..
	ORDER_SHIPPED_EVENT OrderEventType = "ORDER_SHIPPED_EVENT"

	//ORDER_RECEIVED_EVENT ..
	ORDER_RECEIVED_EVENT OrderEventType = "ORDER_RECEIVED_EVENT"

	//ORDER_TRANSPORT_EVENT ..
	ORDER_TRANSPORT_EVENT OrderEventType = "ORDER_TRANSPORT_EVENT"

	//ORDER_NOT_RECEIVED_EVENT ..
	ORDER_NOT_RECEIVED_EVENT OrderEventType = "ORDER_NOT_RECEIVED_EVENT"

	//SAMPLE_COLLECTED_EVENT
	SAMPLE_COLLECTED_EVENT OrderEventType = "SAMPLE_COLLECTED_EVENT"
)

// OrderEventQueryParam ...
type OrderEventQueryParam struct {
	OrderID    *string         `json:"orderID"`
	ExternalID *string         `json:"externalID"`
	FromID     *string         `json:"fromID"`
	ToID       *string         `json:"toID"`
	Type       *OrderEventType `json:"Type"`
}

// OrderEventCreate ..
type OrderEventCreate struct {
	OrderID        string              `json:"orderID"`
	ExternalID     *string             `json:"externalID"`
	Type           OrderEventType      `json:"type"`
	Code           *string             `json:"code"`
	From           ReferenceActorInput `json:"from"`
	To             ReferenceActorInput `json:"to"`
	AdditionalData *[]AttributeInput   `json:"additionalData"`
	TimeStamp      util.Time           `json:"timeStamp"`
}

// OrderEvent ..
type OrderEvent struct {
	Id             string         `json:"id" bson:"_id"`
	OrderID        string         `json:"orderID" bson:"orderID"`
	ExternalID     *string        `json:"externalID" bson:"externalID"`
	Type           OrderEventType `json:"type" bson:"type"`
	Code           *string        `json:"code" bson:"code"`
	From           ReferenceActor `json:"from" bson:"from"`
	To             ReferenceActor `json:"to" bson:"to"`
	AdditionalData *[]Attribute   `json:"additionalData" bson:"additionalData"`
	TimeStamp      util.Time      `json:"timeStamp" bson:"timeStamp"`
	Meta           *models.Meta   //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
