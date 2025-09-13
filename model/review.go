package model

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
	"github.com/Kaleidoscope-Backup/mongo-lib/models"
)

// Emotion ...
type Emotion string

const (
	//LIKE ..
	LIKE Emotion = "LIKE"

	//DISLIKE ..
	DISLIKE Emotion = "DISLIKE"

	//HAPPY ..
	HAPPY Emotion = "HAPPY"

	//SAD ..
	SAD Emotion = "SAD"
)

// ReviewInput ..
type ReviewInput struct {
	Context   ReferenceEntityInput `json:"context"`
	By        ReferenceActorInput  `json:"by"`
	Comment   string               `json:"comment"`
	Rating    *Rating              `json:"rating"`
	Emotion   *Emotion             `json:"emotion" bson:"emotion"`
	Images    *[]AttachmentInput   `json:"images"`
	CreatedAt util.Time            `json:"createdAt"`
}

// Review ..
type Review struct {
	Id        string          `json:"id" bson:"_id"`
	Context   ReferenceEntity `json:"context" bson:"context"`
	By        ReferenceActor  `json:"by" bson:"by"`
	Comment   string          `json:"comment" bson:"comment"`
	Rating    *Rating         `json:"rating" bson:"rating"`
	Emotion   *Emotion        `json:"emotion" bson:"emotion"`
	Images    *[]Attachment   `json:"images" bson:"images"`
	CreatedAt util.Time       `json:"createdAt" bson:"image"`
	Comments  *[]Comment      `json:"comments" bson:"comments"`
	Meta      *models.Meta    //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
