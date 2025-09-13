package model

import (
	"gitlab.com/karte/healthrecord-repository/util"
	"gitlab.com/karte/mongo-lib/models"
)

//AdministrationRoute ...
type AdministrationRoute string

const (
	//ORAL_ADMINISTRATION ..
	ORAL_ADMINISTRATION AdministrationRoute = "ORAL_ADMINISTRATION"
	//INJECTION ..
	INJECTION AdministrationRoute = "INJECTION"
	//INTRAMASCULER_INJECTION ..
	INTRAMASCULER_INJECTION AdministrationRoute = "INTRAMASCULER_INJECTION"
	//SUBCUTENOUS_INJECTION ..
	SUBCUTENOUS_INJECTION AdministrationRoute = "SUBCUTENOUS_INJECTION"
	//INHALATION ..
	INHALATION AdministrationRoute = "INHALATION"
)

//AdministrationRoute ...
type MedicationStatus string

const (
	//MEDICATION_STATUS_ACTIVE ..
	MEDICATION_STATUS_ACTIVE MedicationStatus = "MEDICATION_STATUS_ACTIVE"
	//MEDICATION_STATUS_INACTIVE ..
	MEDICATION_STATUS_INACTIVE MedicationStatus = "MEDICATION_STATUS_INACTIVE"
	//MEDICATION_STATUS_ONHOLD ..
	MEDICATION_STATUS_ONHOLD MedicationStatus = "MEDICATION_STATUS_ONHOLD"
)

//Dosage ..
type Dosage struct {
	Id        string       `json:"id" bson:"_id"`
	Value     int32        `json:"value" bson:"value"`
	Frequency string       `json:"frequency" bson:"frequency"`
	Unit      string       `json:"unit" bson:"unit"`
	Meta      *models.Meta //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

//Strength ..
type Strength struct {
	Id     string       `json:"id" bson:"_id"`
	Number int32        `json:"number" bson:"number"`
	Unit   string       `json:"unit" bson:"unit"`
	Meta   *models.Meta //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

//MedicationCreate input..
type MedicationCreate struct {
	MedicationStatus MedicationStatus     `json:"medicationStatus"`
	Code             *CodableConceptInput `json:"code" bson:"code"`
	ProductName      string               `json:"productName"`
	IsOverTheCounter bool                 `json:"isOverTheCounter"`
	Route            AdministrationRoute  `json:"route" bson:"route"`
	Instructions     string               `json:"instructions"`
	DosageValue      int32                `json:"dosageValue"`
	DosageFrequency  string               `json:"dosageFrequency"`
	DosageUnit       string               `json:"dosageUnit"`
	RefillsRemaining *int32               `json:"refillsRemaining,omitempty"`
	RefillsTotal     *int32               `json:"refillsTotal,omitempty"`
	StrengthNumber   int32                `json:"number"`
	StrengthUnit     string               `json:"unit"`
	Start            util.Time            `json:"start,omitempty"`
	End              *util.Time           `json:"end,omitempty"`
}

//Medication ..
type Medication struct {
	Id               string              `json:"id" bson:"_id"`
	Code             *CodableConcept     `json:"code" bson:"code"`
	MedicationStatus MedicationStatus    `json:"medicationStatus" bson:"medicationStatus"`
	ProductName      string              `json:"productName" bson:"productName"`
	IsOverTheCounter bool                `json:"isOverTheCounter" bson:"isOverTheCounter"`
	Route            AdministrationRoute `json:"route" bson:"route"`
	Instructions     string              `json:"instructions" bson:"instructions"`
	Dosage           *Dosage             `json:"dosage,omitempty" bson:"dosage,omitempty"`
	RefillsRemaining *int32              `json:"refillsRemaining,omitempty" bson:"refillsRemaining,omitempty"`
	RefillsTotal     *int32              `json:"refillsTotal,omitempty" bson:"refillsTotal,omitempty"`
	Strength         *Strength           `json:"strength,omitempty" bson:"strength,omitempty"`
	Start            util.Time           `json:"start,omitempty" bson:"start,omitempty"`
	End              *util.Time          `json:"end,omitempty" bson:"end,omitempty"`
	Meta             *models.Meta        //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
