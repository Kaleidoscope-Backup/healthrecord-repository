package model

import (
	"github.com/Kaleidoscope-Backup/mongo-lib/models"
)

// PractitionerUpdate ...
type PractitionerUpdate struct {
	Id                 string               `json:"_id"`
	Qualification      *string              `json:"qualification"`
	Speciality         *string              `json:"speciality"`
	Photo              *string              `json:"photo"`
	LanguagePreference *string              `json:"languagePreference,omitempty"`
	Contacts           *[]ContactPointInput `json:"contacts"`
}

// PractitionerCreate ...
type PractitionerCreate struct {
	ActorCreate
	Password      string               `json:"password"`
	Organization  string               `json:"organization"`
	Qualification string               `json:"qualification"`
	Speciality    string               `json:"speciality"`
	Photo         *string              `json:"photo"`
	Contacts      *[]ContactPointInput `json:"contacts"`
}

// Practitioner is an actor in our system representing any clinical service provider
type Practitioner struct {
	Actor
	Id            string          `json:"id" bson:"_id"`
	Speciality    string          `json:"speciality" bson:"speciality"`
	Qualification string          `json:"qualification" bson:"qualification"`
	Organization  string          `json:"organization" bson:"organization"`
	Photo         *string         `json:"photo" bson:"photo"`
	Contacts      *[]ContactPoint `json:"contacts" bson:"contacts"`
	Meta          *models.Meta    //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
