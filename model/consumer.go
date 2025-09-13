package model

import (
	"github.com/karte/healthrecord-repository/util"
	"github.com/karte/mongo-lib/models"
)

// Race ...
type Race string

const (
	//AMERICAN_INDIAN_OR_ALASKAN_NATIVE ..
	AMERICAN_INDIAN_OR_ALASKAN_NATIVE Race = "AMERICAN_INDIAN_OR_ALASKAN_NATIVE"

	//ASIAN ..
	ASIAN Race = "ASIAN"

	//BLACK_OR_AFRECIAN_AMERICAN ..
	BLACK_OR_AFRECIAN_AMERICAN Race = "BLACK_OR_AFRECIAN_AMERICAN"

	//HISPANIC_OR_LATINO ..
	HISPANIC_OR_LATINO Race = "HISPANIC_OR_LATINO"

	//NATIVE_HAWAIIAN_OR_OTHER_PACIFIC_ISLANDER ..
	NATIVE_HAWAIIAN_OR_OTHER_PACIFIC_ISLANDER Race = "NATIVE_HAWAIIAN_OR_OTHER_PACIFIC_ISLANDER"

	//WHITE ..
	WHITE Race = "WHITE"
)

// Gender ...
type Gender string

const (
	//MALE ..
	MALE Gender = "MALE"
	//FEMALE ..
	FEMALE Gender = "FEMALE"
)

func (t Gender) toDescription() string {
	switch t {
	case MALE:
		return "Male gender"
	case FEMALE:
		return "Female gender"
	}

	return ""
}

// MarritalStatus ...
type MarritalStatus string

const (
	//MARRIED ..
	MARRIED MarritalStatus = "MARRIED"

	//UNMARRIED ..
	UNMARRIED MarritalStatus = "UNMARRIED"

	//DIVORCED ..
	DIVORCED MarritalStatus = "DIVORCED"

	//WIDOWED ..
	WIDOWED MarritalStatus = "WIDOWED"

	//DOMESTIC_PARTNER ..
	DOMESTIC_PARTNER MarritalStatus = "DOMESTIC_PARTNER"

	//POLYGAMOUS ..
	POLYGAMOUS MarritalStatus = "POLYGAMOUS"

	//LEGALLY_SEPARATED ..
	LEGALLY_SEPARATED MarritalStatus = "LEGALLY_SEPARATED"

	//INTERLOCUTORY ..
	INTERLOCUTORY MarritalStatus = "INTERLOCUTORY"

	//ANNULLED ..
	ANNULLED MarritalStatus = "ANNULLED"
)

// LoginInfoInput ...
type LoginInfoInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ConsumerCreate ....
type ConsumerCreate struct {
	ActorCreate
	Password string `json:"password"`

	//Source ID ...
	SourceID *SourceConsumerIDInput `json:"sourceID"`
	//Primary address ..
	Address             *AddressInput   `json:"address"`
	Ethnicity           *string         `json:"ethnicity"`
	Gender              *Gender         `json:"gender"`
	Race                *Race           `json:"race"`
	MarritalStatus      *MarritalStatus `json:"marritalStatus"`
	DateOfBirth         *util.Time      `json:"dateOfBirth"`
	PrimaryContactType  *string         `json:"primaryContactType"`
	PrimaryContactValue *string         `json:"primaryContactValue"`
	Photo               *string         `json:"photo"`
}

// ConsumerUpdate ....
type ConsumerUpdate struct {
	Id                 string          `json:"id"`
	FirstName          *string         `json:"firstName"`
	LastName           *string         `json:"lastName"`
	Email              *string         `json:"email"`
	Address            *AddressInput   `json:"address"`
	Ethnicity          *string         `json:"ethnicity"`
	Gender             *Gender         `json:"gender"`
	Race               *Race           `json:"race"`
	LanguagePreference *string         `json:"languagePreference,omitempty"`
	MarritalStatus     *MarritalStatus `json:"marritalStatus"`
	Photo              *string         `json:"photo"`
}

// Consumer is an actor in our system representing a patient whose health record is of importance
type Consumer struct {
	Actor
	Id                 string              `json:"id" bson:"_id"`
	SourceIDs          *[]SourceConsumerID `json:"sourceIDs" bson:"sourceIDs"`
	Photo              *string             `json:"photo" bson:"photo"`
	Ethnicity          *string             `json:"ethnicity" bson:"ethnicity"`
	DateOfBirth        *util.Time          `json:"dateOfBirth" bson:"dateOfBirth"`
	MarritalStatus     *MarritalStatus     `json:"marritalStatus" bson:"marritalStatus"`
	Gender             *Gender             `json:"gender" bson:"gender"`
	Race               *Race               `json:"race" bson:"race"`
	PrimaryContact     *ContactPoint       `json:"primaryContact" bson:"primaryContact"`
	AdditionalContacts *[]ContactPoint     `json:"additionalContacts" bson:"additionalContacts"`
	EmergencyContacts  *[]Contact          `json:"emergencyContacts" bson:"emergencyContacts"`
	Address            *Address            `json:"address" bson:"address"`
	Meta               *models.Meta        //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
