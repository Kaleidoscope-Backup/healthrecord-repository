package model

import (
	"github.com/karte/mongo-lib/models"
)

// OrganizationType ...
type OrganizationType string

const (
	//HOSPITAL ..
	HOSPITAL OrganizationType = "HOSPITAL"
	//CLINICAL_SERVICE_PROVIDER ..
	CLINICAL_SERVICE_PROVIDER OrganizationType = "CLINICAL_SERVICE_PROVIDER"
	//INSURANCE_PROVIDER ..
	INSURANCE_PROVIDER OrganizationType = "INSURANCE_PROVIDER"
	//GOVERNMENT ..
	GOVERNMENT OrganizationType = "GOVERNMENT"
	//EDUCATIONAL_INSTITUTE ..
	EDUCATIONAL_INSTITUTE OrganizationType = "EDUCATIONAL_INSTITUTE"
)

// OrganizationQueryParam ...
type OrganizationQueryParam struct {
	Name  *string           `json:"name"`
	Email *string           `json:"email"`
	Type  *OrganizationType `json:"type"`
}

// OrganizationCreate ....
type OrganizationCreate struct {
	Name         string               `json:"name"`
	Email        *string              `json:"email"`
	Type         OrganizationType     `json:"type"`
	SourceID     *string              `json:"sourceID"`
	SourceIDType *string              `json:"sourceIDType"`
	PartOf       *string              `json:"partOf"`
	Photo        *string              `json:"photo"`
	Contacts     *[]ContactPointInput `json:"contacts"`
	Address      *[]AddressInput      `json:"address"`
}

// SourceOrganizationID ...
type SourceOrganizationID struct {
	Id       string       `json:"id" bson:"_id"`
	SourceID string       `json:"sourceID" bson:"sourceID"`
	Type     *string      `json:"type" bson:"type"`
	Meta     *models.Meta //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

// Organization represents a business entity
type Organization struct {
	Id       string                `json:"id" bson:"_id"`
	Name     string                `json:"name" bson:"name"`
	Email    *string               `json:"email" bson:"email"`
	SourceID *SourceOrganizationID `json:"sourceID" bson:"sourceID"`
	Address  *[]Address            `json:"address" bson:"address"`
	Type     OrganizationType      `json:"type" bson:"type"`
	PartOf   *string               `json:"partOf" bson:"partOf"`
	Photo    *string               `json:"photo" bson:"photo"`
	Contacts *[]ContactPoint       `json:"contacts" bson:"contacts"`
	Meta     *models.Meta          //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
