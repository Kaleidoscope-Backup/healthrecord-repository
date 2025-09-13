package model

import "gitlab.com/karte/mongo-lib/models"

//ClinicalTrialCreate ...
type ClinicalTrialCreate struct {
	NCT    string `json:"nctCode"`
	Period Period `json:"period"`
	Reason string `json:"reason"`
}

//ClinicalTrial ...
type ClinicalTrial struct {
	Id         string        `json:"id" bson:"_id"`
	NCT        string        `json:"nctCode" bson:"nctCode"`
	Period     Period        `json:"period" bson:"period"`
	Reason     string        `json:"reason" bson:"reason"`
	ReasonCode *ClinicalCode `json:"reasonCode" bson:"reasonCode"`
	Meta       *models.Meta  //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
