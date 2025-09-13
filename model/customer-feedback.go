package model

import (
	"gitlab.com/karte/healthrecord-repository/util"
	"gitlab.com/karte/mongo-lib/models"
)

//CustomerFeedbackInput ...
type CustomerFeedbackInput struct {
	By          ReferenceActorInput  `json:"by"`
	Application ReferenceEntityInput `json:"application"`
	Subject     string               `json:"subject"`
	Description string               `json:"description"`
	Type        *string              `json:"type"`
	Images      *[]AttachmentInput   `json:"images"`
}

//CustomerFeedback ...
type CustomerFeedback struct {
	Id          string          `json:"id" bson:"_id"`
	By          ReferenceActor  `json:"by" bson:"by"`
	Application ReferenceEntity `json:"application" bson:"application"`
	Subject     string          `json:"subject" bson:"subject"`
	Description string          `json:"description" bson:"description"`
	Type        *string         `json:"type" bson:"type"`
	Images      *[]Attachment   `json:"images" bson:"images"`
	Comments    *[]Comment      `json:"comments" bson:"comments"`
	CreatedAt   util.Time       `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	Meta        *models.Meta    //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
