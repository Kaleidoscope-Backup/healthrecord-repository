package model

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
	"github.com/Kaleidoscope-Backup/mongo-lib/models"
)

// AllergyStatus ...
type AllergyStatus string

const (
	//ALLERGY_ACTIVE ..
	ALLERGY_ACTIVE AllergyStatus = "ALLERGY_ACTIVE"
	//ALLERGY_INACTIVE ..
	ALLERGY_INACTIVE AllergyStatus = "ALLERGY_INACTIVE"
	//ALLERGY_RESOLVED ..
	ALLERGY_RESOLVED AllergyStatus = "ALLERGY_RESOLVED"
)

// AllergyCategory ...
type AllergyCategory string

const (
	//ALLERGY_FOOD ..
	ALLERGY_FOOD AllergyCategory = "ALLERGY_FOOD"
	//ALLERGY_MEDICATION ..
	ALLERGY_MEDICATION AllergyCategory = "ALLERGY_MEDICATION"
	//ALLERGY_ENVIRONMENT ..
	ALLERGY_ENVIRONMENT AllergyCategory = "ALLERGY_ENVIRONMENT"
	//ALLERGY_BIOLOGIC ..
	ALLERGY_BIOLOGIC AllergyCategory = "ALLERGY_BIOLOGIC"
)

// AllergyCriticality ...
type AllergyCriticality string

const (
	//ALLERGY_LOW ..
	ALLERGY_LOW AllergyCriticality = "ALLERGY_LOW"
	//ALLERGY_HIGH ..
	ALLERGY_HIGH AllergyCriticality = "ALLERGY_HIGH"
	//ALLERGY_UNABLE_TO_ASSES ..
	ALLERGY_UNABLE_TO_ASSES AllergyCriticality = "ALLERGY_UNABLE_TO_ASSES"
)

// AllergyOnsetCreate ...
type AllergyOnsetCreate struct {
	OnsetDate *util.Time `json:"onsetDate"`
	OnsetAge  *string    `json:"onsetAge"`
	OnsetNote *string    `json:"onsetNote"`
}

// AllergyReactionInput ...
type AllergyReactionInput struct {
	Substance         string               `json:"substance"`
	SubstanceCode     *CodableConceptInput `json:"substanceCode"`
	Manifestation     string               `json:"manifestation"`
	ManifestationCode *CodableConceptInput `json:"manifestationCode"`
	ExposureRoute     string               `json:"exposureRoute"`
	ExposureRouteCode *CodableConceptInput `json:"exposureRouteCode"`
	Description       *string              `json:"description"`
	Severity          *Severity            `json:"severity"`
}

// AllergyReaction ...
type AllergyReaction struct {
	Id                string          `json:"id" bson:"_id"`
	Substance         string          `json:"substance" bson:"substance"`
	SubstanceCode     *CodableConcept `json:"substanceCode" bson:"substanceCode"`
	Manifestation     string          `json:"manifestation" bson:"manifestation"`
	ManifestationCode *CodableConcept `json:"manifestationCode" bson:"manifestationCode"`
	ExposureRoute     string          `json:"exposureRoute" bson:"exposureRoute"`
	ExposureRouteCode *CodableConcept `json:"exposureRouteCode" bson:"exposureRouteCode"`
	Description       *string         `json:"description" bson:"description"`
	Severity          *Severity       `json:"severity" bson:"severity"`
	Meta              *models.Meta    //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

// AllergyOnset ...
type AllergyOnset struct {
	Id        string       `json:"id" bson:"_id"`
	OnsetDate *util.Time   `json:"onsetDate" bson:"onsetDate"`
	OnsetAge  *string      `json:"onsetAge" bson:"onsetAge"`
	OnsetNote *string      `json:"onsetNote" bson:"onsetNote"`
	Meta      *models.Meta //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

// AllergyRecordCreate ...
type AllergyRecordCreate struct {
	HealthRecordCreate
	OnsetDate      *AllergyOnsetCreate     `json:"onsetDate"`
	LastOccurrence *util.Time              `json:"lastOccurrence"`
	Category       AllergyCategory         `json:"category"`
	Criticality    AllergyCriticality      `json:"criticality"`
	Status         AllergyStatus           `json:"status"`
	Code           *CodableConceptInput    `json:"code"`
	Reactions      *[]AllergyReactionInput `json:"reactions"`
}

// AllergyRecord ...
type AllergyRecord struct {
	HealthRecord
	Id             string             `json:"id" bson:"_id"`
	OnsetDate      *AllergyOnset      `json:"onsetDate" bson:"onsetDate"`
	LastOccurrence *util.Time         `json:"lastOccurrence" bson:"lastOccurrence"`
	Category       AllergyCategory    `json:"category" bson:"category"`
	Criticality    AllergyCriticality `json:"criticality" bson:"criticality"`
	Status         AllergyStatus      `json:"status" bson:"status"`
	Code           *CodableConcept    `json:"code" bson:"code"`
	Reactions      *[]AllergyReaction `json:"reactions" bson:"reactions"`
	Meta           *models.Meta       //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
