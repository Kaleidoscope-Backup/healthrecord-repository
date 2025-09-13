package model

import (
	"gitlab.com/karte/healthrecord-repository/util"
	"gitlab.com/karte/mongo-lib/models"
)

//CommentInput ...
type CommentInput struct {
	CommentText string               `json:"commentText"`
	Context     ReferenceEntityInput `json:"context"`
	Attachments *[]AttachmentInput   `json:"attachments"`
	CommentedBy ReferenceActorInput  `json:"commentedBy"`
	Location    *GeoLocationInput    `json:"location"`
}

//CommentOnCommentInput ...
type CommentOnCommentInput struct {
	ExternalID  string              `json:"externalID"`
	CommentText string              `json:"commentText"`
	Attachments *[]AttachmentInput  `json:"attachments"`
	CommentedBy ReferenceActorInput `json:"commentedBy"`
}

//Comment ...
type Comment struct {
	Id          string          `json:"id" bson:"_id"`
	ExternalID  string          `json:"externalID" bson:"externalID"`
	CreatedAt   util.Time       `json:"createdAt" bson:"createdAt"`
	Context     ReferenceEntity `json:"context" bson:"context"`
	CommentText string          `json:"commentText" bson:"commentText"`
	CommentedBy ReferenceActor  `json:"commentedBy" bson:"commentedBy"`
	Comments    *[]Comment      `json:"comments" bson:"comments"`
	Attachments *[]Attachment   `json:"attachments" bson:"attachments"`
	Location    *GeoLocation    `json:"location" bson:"location"`
	Meta        *models.Meta    //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
