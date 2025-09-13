package model

import (
	"gitlab.com/karte/healthrecord-repository/util"
	"gitlab.com/karte/mongo-lib/models"
)

//DocumentReferenceStatus ...
type DocumentReferenceStatus string

const (
	//DOCUMENT_CURRENT ..
	DOCUMENT_CURRENT DocumentReferenceStatus = "DOCUMENT_CURRENT"

	//DOCUMENT_SUSPENDED ..
	DOCUMENT_SUSPENDED DocumentReferenceStatus = "DOCUMENT_SUSPENDED"

	//DOCUMENT_ENTERED_IN_ERROR ..
	DOCUMENT_ENTERED_IN_ERROR DocumentReferenceStatus = "DOCUMENT_ENTERED_IN_ERROR"
)

//CompositionStatus ...
type CompositionStatus string

const (
	//COMPOSITIONSTATUS_PRELIMINARY ..
	COMPOSITIONSTATUS_PRELIMINARY CompositionStatus = "COMPOSITIONSTATUS_PRELIMINARY"

	//COMPOSITIONSTATUS_FINAL ..
	COMPOSITIONSTATUS_FINAL CompositionStatus = "COMPOSITIONSTATUS_FINAL"

	//COMPOSITIONSTATUS_APPENDED ..
	COMPOSITIONSTATUS_APPENDED CompositionStatus = "COMPOSITIONSTATUS_APPENDED"

	//COMPOSITIONSTATUS_AMENDED ..
	COMPOSITIONSTATUS_AMENDED CompositionStatus = "COMPOSITIONSTATUS_AMENDED"

	//COMPOSITIONSTATUS_ENTERED_IN_ERROR ..
	COMPOSITIONSTATUS_ENTERED_IN_ERROR CompositionStatus = "COMPOSITIONSTATUS_ENTERED_IN_ERROR"
)

//DocumentReferenceQueryParam ...
type DocumentReferenceQueryParam struct {
	Type      *string  `json:"type"`
	Language  Language `json:"language"`
	Class     string   `json:"class"`
	Custodian string   `json:"custodian"`
}

//DocumentContentInput ...
type DocumentContentInput struct {
	Content    string           `json:"content"`
	Attachment *AttachmentInput `json:"attachment"`
}

//DocumentContent ...
type DocumentContent struct {
	Id         string       `json:"id" bson:"_id"`
	Content    string       `json:"content" bson:"content"`
	Attachment *Attachment  `json:"attachment" bson:"attachment"`
	Meta       *models.Meta //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

//DocumentReferenceInput ...
type DocumentReferenceInput struct {
	Language          Language                `json:"language"`
	Status            DocumentReferenceStatus `json:"status"`
	CompositionStatus CompositionStatus       `json:"compositionStatus"`
	Type              string                  `json:"type"`
	TypeCode          *ClinicalCodeInput      `json:"typeCode"`
	Class             string                  `json:"class"`
	ClassCode         *ClinicalCodeInput      `json:"classCode"`
	Created           util.Time               `json:"created"`
	Author            *ReferenceActorInput    `json:"author"`
	Authenticator     *ReferenceActorInput    `json:"authenticator"`
	Custodian         ReferenceActorInput     `json:"custodian"`
	Description       *string                 `json:"description"`
	SecurityLabel     *ClinicalCodeInput      `json:"securityLabel"`
	Context           *ReferenceEntityInput   `json:"context"`
	Content           *[]DocumentContentInput `json:"content"`
}

//DocumentReference ...
type DocumentReference struct {
	Id                string                  `json:"id" bson:"_id"`
	Language          Language                `json:"language" bson:"language"`
	Status            DocumentReferenceStatus `json:"status" bson:"status"`
	CompositionStatus CompositionStatus       `json:"compositionStatus" bson:"compositionStatus"`
	Type              string                  `json:"type" bson:"type"`
	TypeCode          *ClinicalCode           `json:"typeCode" bson:"typeCode"`
	Class             string                  `json:"class" bson:"class"`
	ClassCode         *ClinicalCode           `json:"classCode" bson:"classCode"`
	Created           util.Time               `json:"created" bson:"created"`
	Author            *ReferenceActor         `json:"author" bson:"author"`
	Authenticator     *ReferenceActor         `json:"authenticator" bson:"authenticator"`
	Custodian         ReferenceActor          `json:"custodian" bson:"custodian"`
	Description       *string                 `json:"description" bson:"description"`
	SecurityLabel     *ClinicalCode           `json:"securityLabel" bson:"securityLabel"`
	Context           *ReferenceEntity        `json:"context" bson:"context"`
	Content           *[]DocumentContent      `json:"content" bson:"content"`
	Meta              *models.Meta            //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
