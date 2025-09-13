package model

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
	"github.com/Kaleidoscope-Backup/mongo-lib/models"
)

// LocationMode ...
type LocationMode string

const (
	//LOCATION_KIND ..
	LOCATION_KIND LocationMode = "LOCATION_KIND"

	//LOCATION_INSTANCE ..
	LOCATION_INSTANCE LocationMode = "LOCATION_INSTANCE"
)

// LocationPhysicalType ...
type LocationPhysicalType string

const (
	//LOCATION_SITE ..
	LOCATION_SITE LocationPhysicalType = "LOCATION_SITE"

	//LOCATION_BUILDING ..
	LOCATION_BUILDING LocationPhysicalType = "LOCATION_BUILDING"

	//LOCATION_WING ..
	LOCATION_WING LocationPhysicalType = "LOCATION_WING"

	//LOCATION_WARD ..
	LOCATION_WARD LocationPhysicalType = "LOCATION_WARD"

	//LOCATION_LEVEL ..
	LOCATION_LEVEL LocationPhysicalType = "LOCATION_LEVEL"

	//LOCATION_CORRIDOR ..
	LOCATION_CORRIDOR LocationPhysicalType = "LOCATION_CORRIDOR"

	//LOCATION_ROOM ..
	LOCATION_ROOM LocationPhysicalType = "LOCATION_ROOM"

	//LOCATION_BED ..
	LOCATION_BED LocationPhysicalType = "LOCATION_BED"

	//LOCATION_VEHICLE ..
	LOCATION_VEHICLE LocationPhysicalType = "LOCATION_VEHICLE"

	//LOCATION_HOUSE ..
	LOCATION_HOUSE LocationPhysicalType = "LOCATION_HOUSE"

	//LOCATION_CABINET ..
	LOCATION_CABINET LocationPhysicalType = "LOCATION_CABINET"

	//LOCATION_ROAD ..
	LOCATION_ROAD LocationPhysicalType = "LOCATION_ROAD"

	//LOCATION_AREA ..
	LOCATION_AREA LocationPhysicalType = "LOCATION_AREA"

	//LOCATION_JURISDICTION ..
	LOCATION_JURISDICTION LocationPhysicalType = "LOCATION_JURISDICTION"
)

// LocationInput ...
type LocationInput struct {
	Name                 string                 `json:"name" bson:"name"`
	Alias                *[]string              `json:"alias" bson:"alias"`
	Description          *string                `json:"description" bson:"description"`
	Mode                 LocationMode           `json:"mode" bson:"mode"`
	Type                 *[]string              `json:"type" bson:"type"`
	TypeCode             *[]CodableConceptInput `json:"typeCode" bson:"typeCode"`
	Telecom              *[]ContactPointInput   `json:"telecom" bson:"telecom"`
	Address              AddressInput           `json:"address" bson:"address"`
	PhysicalType         *LocationPhysicalType  `json:"physicalType" bson:"physicalType"`
	ManagingOrganization *ReferenceActorInput   `json:"managingOrganization" bson:"managingOrganization"`
	Position             *[]GeoLocationInput    `json:"position" bson:"position"`
	PartOf               *ReferenceEntityInput  `json:"partOf" bson:"partOf"`
	AllDay               *bool                  `json:"allDay" bson:"allDay"`
	DaysOfWeek           *DaysOfWeek            `json:"daysOfWeek" bson:"daysOfWeek"`
	OpeningTime          *util.Time             `json:"openingTime" bson:"openingTime"`
	ClosingTime          *util.Time             `json:"closingTime" bson:"closingTime"`
}

// Location ...
type Location struct {
	Id                   string                `json:"id" bson:"_id"`
	Name                 string                `json:"name" bson:"name"`
	Alias                *[]string             `json:"alias" bson:"alias"`
	Description          *string               `json:"description" bson:"description"`
	Mode                 LocationMode          `json:"mode" bson:"mode"`
	Type                 *[]string             `json:"type" bson:"type"`
	TypeCode             *[]CodableConcept     `json:"typeCode" bson:"typeCode"`
	Telecom              *[]ContactPoint       `json:"telecom" bson:"telecom"`
	Address              Address               `json:"address" bson:"address"`
	PhysicalType         *LocationPhysicalType `json:"physicalType" bson:"physicalType"`
	ManagingOrganization *ReferenceActor       `json:"managingOrganization" bson:"managingOrganization"`
	Position             *[]GeoLocation        `json:"position" bson:"position"`
	PartOf               *ReferenceEntity      `json:"partOf" bson:"partOf"`
	AllDay               *bool                 `json:"allDay" bson:"allDay"`
	DaysOfWeek           *DaysOfWeek           `json:"daysOfWeek" bson:"daysOfWeek"`
	OpeningTime          *util.Time            `json:"openingTime" bson:"openingTime"`
	ClosingTime          *util.Time            `json:"closingTime" bson:"closingTime"`
	Meta                 *models.Meta          //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
