package model

import (
	"github.com/karte/healthrecord-repository/util"
	"github.com/karte/mongo-lib/models"
)

// ImmunizationRecordCreate ...
type ImmunizationRecordCreate struct {
	HealthRecordCreate
	Code             *CodableConceptInput `json:"code"`
	Vaccine          string               `json:"vaccine"`
	NotGiven         *bool                `json:"notGiven"`
	AdministeredDate *util.Time           `json:"administeredDate"`
	AdministeredBy   *string              `json:"administeredBy"`
	Route            *AdministrationRoute `json:"route"`
	RouteCode        *CodableConceptInput `json:"routeCode"`
	Reaction         *string              `json:"reaction"`
	ReactionCode     *CodableConceptInput `json:"reactionCode"`
	Manufacturer     *string              `json:"manufacturer"`
	ExperiationDate  *util.Time           `json:"experiationDate"`
}

// ImmunizationRecord ...
type ImmunizationRecord struct {
	HealthRecord
	Id               string               `json:"id" bson:"_id"`
	Code             *CodableConcept      `json:"code" bson:"code"`
	Vaccine          string               `json:"vaccine" bson:"vaccine"`
	NotGiven         *bool                `json:"notGiven" bson:"notGiven"`
	AdministeredDate *util.Time           `json:"administeredDate" bson:"administeredDate"`
	AdministeredBy   *string              `json:"administeredBy" bson:"administeredBy"`
	Route            *AdministrationRoute `json:"route" bson:"route"`
	RouteCode        *CodableConcept      `json:"routeCode" bson:"routeCode"`
	Reaction         *string              `json:"reaction" bson:"reaction"`
	ReactionCode     *CodableConcept      `json:"reactionCode" bson:"reactionCode"`
	Manufacturer     *string              `json:"manufacturer" bson:"manufacturer"`
	ExperiationDate  *util.Time           `json:"experiationDate" bson:"experiationDate"`
	Meta             *models.Meta         //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
