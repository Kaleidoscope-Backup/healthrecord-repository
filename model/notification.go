package model

import (
	"gitlab.com/karte/healthrecord-repository/util"
	"gitlab.com/karte/mongo-lib/models"
)

//NotificationStatus ...
type NotificationStatus string

const (
	//ACTIVE_NOTIFICATION ..
	ACTIVE_NOTIFICATION NotificationStatus = "ACTIVE_NOTIFICATION"
	//ACKNOWLEDGED_NOTIFICATION ..
	ACKNOWLEDGED_NOTIFICATION NotificationStatus = "ACKNOWLEDGED_NOTIFICATION"
)

//NotificationCreate ..
type NotificationCreate struct {
	Name        string               `json:"name"`
	Category    string               `json:"category"`
	Description *string              `json:"description"`
	Created     util.Time            `json:"created,omitempty"`
	Updated     *util.Time           `json:"updated,omitempty"`
	ConsumerID  string               `json:"consumerID"`
	Reference   ReferenceEntityInput `json:"reference"`
	AckOptions  []string             `json:"ackOptions"`
}

//Notification ..
type Notification struct {
	Id          string             `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Category    string             `json:"category" bson:"category"`
	Description *string            `json:"description" bson:"description"`
	Status      NotificationStatus `json:"status" bson:"status"`
	Created     util.Time          `json:"created,omitempty" bson:"created,omitempty"`
	Updated     *util.Time         `json:"updated,omitempty" bson:"updated,omitempty"`
	ConsumerID  string             `json:"consumerID" bson:"consumerID"`
	Reference   ReferenceEntity    `json:"reference" bson:"reference"`
	AckOptions  []string           `json:"ackOptions" bson:"ackOptions"`
	Meta        *models.Meta       //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
