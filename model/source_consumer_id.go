package model

import (
	"gitlab.com/karte/healthrecord-repository/util"
	"gitlab.com/karte/mongo-lib/models"
)

//SourceConsumerIDInput is for input ...
type SourceConsumerIDInput struct {
	System   string                `json:"system" bson:"system"`
	Value    string                `json:"qualification" bson:"qualification"`
	Assigner string                `json:"assigner" bson:"assigner"`
	Use      *SourceConsumerIDUse  `json:"use,omitempty" bson:"use,omitempty"`
	Type     *SourceConsumerIDType `json:"type,omitempty" bson:"type,omitempty"`
	Start    *util.Time            `json:"start,omitempty" bson:"start,omitempty"`
	End      *util.Time            `json:"end,omitempty" bson:"end,omitempty"`
}

// SourceConsumerID represents the Id of patient/ practitioner in source system from where data is imported
type SourceConsumerID struct {
	Id       string                `json:"id" bson:"_id"`
	System   string                `json:"system" bson:"system"`
	Value    string                `json:"qualification" bson:"qualification"`
	Assigner string                `json:"assigner" bson:"assigner"`
	Use      *SourceConsumerIDUse  `json:"use,omitempty" bson:"use,omitempty"`
	Type     *SourceConsumerIDType `json:"type,omitempty" bson:"type,omitempty"`
	Start    *util.Time            `json:"start,omitempty" bson:"start,omitempty"`
	End      *util.Time            `json:"end,omitempty" bson:"end,omitempty"`
	Meta     *models.Meta          //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

//SourceConsumerIDUse ...
type SourceConsumerIDUse string

const (
	//USUAL ...
	USUAL SourceConsumerIDUse = "USUAL"
	//OFFICIAL ...
	OFFICIAL SourceConsumerIDUse = "OFFICIAL"
	//TEMP ...
	TEMP SourceConsumerIDUse = "TEMP"
	//SECONDARY ...
	SECONDARY SourceConsumerIDUse = "SECONDARY"
)

//SourceConsumerIDType ...
type SourceConsumerIDType string

const (
	//DL ...
	DL SourceConsumerIDType = "DL"
	//PPN ...
	PPN SourceConsumerIDType = "PPN"
	//BRN ...
	BRN SourceConsumerIDType = "BRN"
	//MR ...
	MR SourceConsumerIDType = "MR"
	//MCN ...
	MCN SourceConsumerIDType = "MCN"
	//EN ...
	EN SourceConsumerIDType = "EN"
	//TAX ...
	TAX SourceConsumerIDType = "TAX"
	//NIIP ...
	NIIP SourceConsumerIDType = "NIIP"
	//PRN ...
	PRN SourceConsumerIDType = "PRN"
	//MD
	MD SourceConsumerIDType = "MD"
	//DR ...
	DR SourceConsumerIDType = "DR"
	//ACSN ...
	ACSN SourceConsumerIDType = "ACSN"
	//AUTH_0
	AUTH_0 SourceConsumerIDType = "AUTH_0"
	//ETHEREUM_ID
	ETHEREUM_ID SourceConsumerIDType = "ETHEREUM_ID"
)

func (t SourceConsumerIDType) toDescription() string {
	switch t {
	case DL:
		return "Driver's license number"
	case PPN:
		return "Passport number"
	case BRN:
		return "Breed Registry Number"
	case MR:
		return "Medical record number"
	case MCN:
		return "Microchip Number"
	case EN:
		return "Employer number"
	case TAX:
		return "Tax ID number"
	case NIIP:
		return "National Insurance Payor Identifier (Payor)"
	case PRN:
		return "Provider number"
	case DR:
		return "Donor Registration Number"
	case ACSN:
		return "Accession ID"
	}

	return ""
}
