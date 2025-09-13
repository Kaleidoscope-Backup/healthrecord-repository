package model

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
	"github.com/Kaleidoscope-Backup/mongo-lib/models"
)

// ListStatus ...
type ListStatus string

const (
	//LIST_RETIRED ..
	LIST_RETIRED ListStatus = "LIST_RETIRED"

	//LIST_ACTIVE ..
	LIST_ACTIVE ListStatus = "LIST_ACTIVE"

	//LIST_ENTERED_IN_ERROR ..
	LIST_ENTERED_IN_ERROR ListStatus = "LIST_ENTERED_IN_ERROR"
)

// ListMode ...
type ListMode string

const (
	//LIST_WORKING ..
	LIST_WORKING ListMode = "LIST_WORKING"

	//LIST_SNAPSHOT ..
	LIST_SNAPSHOT ListMode = "LIST_SNAPSHOT"

	//LIST_CHANGES ..
	LIST_CHANGES ListMode = "LIST_CHANGES"
)

// ListEntry ...
type ListEntry struct {
	Id      string       `json:"id" bson:"_id"`
	Date    *util.Time   `json:"date" bson:"date"`
	Deleted *bool        `json:"deleted" bson:"deleted"`
	Entry   *[]Attribute `json:"entry" bson:"entry"`
	Meta    *models.Meta //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

// ListEntryInput ...
type ListEntryInput struct {
	Date    *util.Time        `json:"date"`
	Deleted *bool             `json:"deleted"`
	Entry   *[]AttributeInput `json:"entry"`
}

// ListInput ...
type ListInput struct {
	Status  ListStatus            `json:"status"`
	Mode    ListMode              `json:"mode"`
	Title   string                `json:"title"`
	Code    *CodableConceptInput  `json:"code"`
	Subject *ReferenceEntityInput `json:"subject"`
	Owner   ReferenceActorInput   `json:"owner"`
	Source  *ReferenceEntityInput `json:"source"`
	Note    *string               `json:"note"`
	Items   *[]ListEntryInput     `json:"items"`
}

// List ...
type List struct {
	Id      string           `json:"id" bson:"_id"`
	Status  ListStatus       `json:"status" bson:"status"`
	Mode    ListMode         `json:"mode" bson:"mode"`
	Title   string           `json:"title" bson:"title"`
	Code    *CodableConcept  `json:"code" bson:"code"`
	Subject *ReferenceEntity `json:"subject" bson:"subject"`
	Owner   ReferenceActor   `json:"owner" bson:"owner"`
	Source  *ReferenceEntity `json:"source" bson:"source"`
	Note    *string          `json:"note" bson:"note"`
	Items   *[]ListEntry     `json:"items" bson:"items"`
	Meta    *models.Meta     //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
