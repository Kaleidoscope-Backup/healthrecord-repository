package model

import (
	"gitlab.com/karte/healthrecord-repository/util"
	"gitlab.com/karte/mongo-lib/models"
)

//Prioroty ...
type Prioroty string

const (
	//ROUTINE ..
	ROUTINE Prioroty = "ROUTINE"
	//URGENT ..
	URGENT Prioroty = "URGENT"
	//ASAP ..
	ASAP Prioroty = "ASAP"
	//EMERGENCY ..
	EMERGENCY Prioroty = "EMERGENCY"
)

//SortBy ...
type SortBy string

const (
	//LAST_UPDATED ..
	LAST_UPDATED SortBy = "LAST_UPDATED"
)

//EntityType ...
type EntityType string

const (
	//ENTITY_CONSUMER ..
	ENTITY_CONSUMER EntityType = "ENTITY_CONSUMER"

	//ENTITY_PRACTITIONER ..
	ENTITY_PRACTITIONER EntityType = "ENTITY_PRACTITIONER"

	//ENTITY_ORGANIZATION ..
	ENTITY_ORGANIZATION EntityType = "ENTITY_ORGANIZATION"

	//ENTITY_HEALTHRECORD ..
	ENTITY_HEALTHRECORD EntityType = "ENTITY_HEALTHRECORD"

	//ENTITY_CONSENT ..
	ENTITY_CONSENT EntityType = "ENTITY_CONSENT"

	//ENTITY_QUESTIONNAIRE ..
	ENTITY_QUESTIONNAIRE EntityType = "ENTITY_QUESTIONNAIRE"

	//ENTITY_QUESTIONNAIRE_RESPONSE ..
	ENTITY_QUESTIONNAIRE_RESPONSE EntityType = "ENTITY_QUESTIONNAIRE_RESPONSE"

	//ENTITY_CLINICAL_TRIAL ..
	ENTITY_CLINICAL_TRIAL EntityType = "ENTITY_CLINICAL_TRIAL"

	//ENTITY_ORDER ..
	ENTITY_ORDER EntityType = "ENTITY_ORDER"

	//ENTITY_ORDER_EVENT ..
	ENTITY_ORDER_EVENT EntityType = "ENTITY_ORDER_EVENT"

	//ENTITY_PRODUCT ..
	ENTITY_PRODUCT EntityType = "ENTITY_PRODUCT"

	//ENTITY_NOTIFICATION ..
	ENTITY_NOTIFICATION EntityType = "ENTITY_NOTIFICATION"

	//ENTITY_ACKNOWLEDGEMENT ..
	ENTITY_ACKNOWLEDGEMENT EntityType = "ENTITY_ACKNOWLEDGEMENT"

	//ENTITY_SCHEDULE ..
	ENTITY_SCHEDULE EntityType = "ENTITY_SCHEDULE"

	//ENTITY_SLOT ..
	ENTITY_SLOT EntityType = "ENTITY_SLOT"

	//ENTITY_REFERRAL_REQUEST ..
	ENTITY_REFERRAL_REQUEST EntityType = "ENTITY_REFERRAL_REQUEST"

	//ENTITY_DEVICE ..
	ENTITY_DEVICE EntityType = "ENTITY_DEVICE"

	//ENTITY_DEVICE_METRIC ...
	ENTITY_DEVICE_METRIC EntityType = "ENTITY_DEVICE_METRIC"

	//ENTITY_APPLICATION ...
	ENTITY_APPLICATION EntityType = "ENTITY_APPLICATION"

	//ENTITY_MESSAGE ...
	ENTITY_MESSAGE EntityType = "ENTITY_MESSAGE"

	//ENTITY_REVIEW ...
	ENTITY_REVIEW EntityType = "ENTITY_REVIEW"

	//ENTITY_COMMENT ...
	ENTITY_COMMENT EntityType = "ENTITY_COMMENT"

	//ENTITY_RELATIONSHIP ...
	ENTITY_RELATIONSHIP EntityType = "ENTITY_RELATIONSHIP"

	//ENTITY_LOCATION ...
	ENTITY_LOCATION EntityType = "ENTITY_LOCATION"

	//ENTITY_HEALTHSCAREERVICE ...
	ENTITY_HEALTHSCAREERVICE EntityType = "ENTITY_HEALTHSCAREERVICE"
)

//AgeGroup ...
type AgeGroup string

const (
	//NEWBORN ..
	NEWBORN AgeGroup = "INFANT"

	//INFANT ..
	INFANT AgeGroup = "INFANT"

	//CHILDREN ..
	CHILDREN AgeGroup = "CHILDREN"

	//YOUNG_ADULT ..
	YOUNG_ADULT AgeGroup = "YOUNG_ADULT"

	//ADULT ..
	ADULT AgeGroup = "ADULT"

	//SENIOR ..
	SENIOR AgeGroup = "SENIOR"
)

//MimeType ...
type MimeType string

const (
	//TEXT_PLAIN ..
	TEXT_PLAIN MimeType = "TEXT_PLAIN"

	//TEXT_HTML ..
	TEXT_HTML MimeType = "TEXT_HTML"

	//IMAGE ..
	IMAGE MimeType = "IMAGE"

	//VIDEO ..
	VIDEO MimeType = "VIDEO"

	//PDF ..
	PDF MimeType = "PDF"
)

//Severity ...
type Severity string

const (
	//SEVERITY_MILD ..
	SEVERITY_MILD Severity = "SEVERITY_MILD"

	//SEVERITY_MEDIUM ..
	SEVERITY_MEDIUM Severity = "SEVERITY_MEDIUM"

	//SEVERITY_HIGH ..
	SEVERITY_HIGH Severity = "SEVERITY_HIGH"
)

//Priority ...
type Priority string

const (
	//PRIORITY_HIGH ..
	PRIORITY_HIGH Priority = "PRIORITY_HIGH"

	//PRIORITY_MEDIUM ..
	PRIORITY_MEDIUM Priority = "PRIORITY_MEDIUM"

	//PRIORITY_LOW ..
	PRIORITY_LOW Priority = "PRIORITY_LOW"
)

// Result ...
type Result struct {
	Success bool `json:"success"`
}

//ReferenceEntityInput ...
type ReferenceEntityInput struct {
	EntityType EntityType `json:"entityType"`
	EntityID   string     `json:"entityID"`
}

//ReferenceEntity ...
type ReferenceEntity struct {
	Id         string     `json:"id" bson:"_id"`
	EntityType EntityType `json:"entityType" bson:"entityType"`
	EntityID   string     `json:"entityID" bson:"entityID"`
}

//AttachmentInput ...
type AttachmentInput struct {
	ContentType MimeType  `json:"contentType" bson:"contentType"`
	Language    *string   `json:"language" bson:"language"`
	URL         string    `json:"url" bson:"url"`
	Size        *int32    `json:"size" bson:"size"`
	Title       string    `json:"title" bson:"title"`
	CreatedOn   util.Time `json:"createdOn" bson:"createdOn"`
}

//Attachment ...
type Attachment struct {
	Id          string       `json:"id" bson:"_id"`
	ContentType MimeType     `json:"contentType" bson:"contentType"`
	Language    *string      `json:"language" bson:"language"`
	URL         string       `json:"url" bson:"url"`
	Size        *int32       `json:"size" bson:"size"`
	Title       string       `json:"title" bson:"title"`
	CreatedOn   util.Time    `json:"createdOn" bson:"createdOn"`
	Meta        *models.Meta //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

//Contact represents the emergency contact
type Contact struct {
	Id           string       `json:"id" bson:"_id"`
	Name         string       `json:"name" bson:"name"`
	Relationship string       `json:"relationship" bson:"relationship"`
	Phone        string       `json:"phone" bson:"phone"`
	Email        string       `json:"email" bson:"email"`
	Meta         *models.Meta //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

//SymptomInput ...
type SymptomInput struct {
	Name string               `json:"name"`
	Code *CodableConceptInput `json:"code"`
}

//Symptom ...
type Symptom struct {
	Id   string          `json:"id" bson:"_id"`
	Name string          `json:"name" bson:"name"`
	Code *CodableConcept `json:"code" bson:"code"`
	Meta *models.Meta    //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

//OnsetInput ...
type OnsetInput struct {
	Date *util.Time `json:"date" bson:"date"`
	Age  *string    `json:"age" bson:"age"`
	Note *string    `json:"note" bson:"note"`
}

//Onset ...
type Onset struct {
	Id   string       `json:"id" bson:"_id"`
	Date *util.Time   `json:"date" bson:"date"`
	Age  *string      `json:"age" bson:"age"`
	Note *string      `json:"note" bson:"note"`
	Meta *models.Meta //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

//AbatementInput ...
type AbatementInput struct {
	Abatement *bool      `json:"abatement" bson:"abatement"`
	Date      *util.Time `json:"date" bson:"date"`
	Age       *string    `json:"age" bson:"age"`
	Note      *string    `json:"note" bson:"note"`
}

//Abatement ...
type Abatement struct {
	Id        string       `json:"id" bson:"_id"`
	Abatement *bool        `json:"abatement" bson:"abatement"`
	Date      *util.Time   `json:"date" bson:"date"`
	Age       *string      `json:"age" bson:"age"`
	Note      *string      `json:"note" bson:"note"`
	Meta      *models.Meta //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
