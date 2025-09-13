package model

import (
	"github.com/karte/healthrecord-repository/util"
	"github.com/karte/mongo-lib/models"
)

// MessageInput ...
type MessageInput struct {
	From        ReferenceActorInput           `json:"from"`
	To          ReferenceActorInput           `json:"to"`
	Message     string                        `json:"message"`
	Attachments *[]AttachmentInput            `json:"attachments"`
	Records     *[]ReferenceHealthRecordInput `json:"records"`
	CreatedAt   util.Time                     `json:"createdAt"`
}

// Message ...
type Message struct {
	Id          string                   `json:"id" bson:"_id"`
	From        ReferenceActor           `json:"from" bson:"from"`
	To          ReferenceActor           `json:"to" bson:"to"`
	Message     string                   `json:"message" bson:"message"`
	Attachments *[]Attachment            `json:"attachments" bson:"attachments"`
	Records     *[]ReferenceHealthRecord `json:"records" bson:"records"`
	CreatedAt   util.Time                `json:"createdAt" bson:"createdAt"`
	Meta        *models.Meta             //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
