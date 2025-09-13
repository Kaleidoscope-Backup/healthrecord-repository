package model

import (
	"gitlab.com/karte/healthrecord-repository/util"
	"gitlab.com/karte/mongo-lib/models"
)

//UDIType enum ...
type UDIType string

const (
	//DEVICE_BARCODE ...
	DEVICE_BARCODE UDIType = "DEVICE_BARCODE"

	//DEVICE_RFID ...
	DEVICE_RFID UDIType = "DEVICE_RFID"

	//DEVICE_MANUAL ...
	DEVICE_MANUAL UDIType = "DEVICE_MANUAL"

	//DEVICE_CARD ...
	DEVICE_CARD UDIType = "DEVICE_CARD"

	//DEVICE_SELFREPORTED ...
	DEVICE_SELFREPORTED UDIType = "DEVICE_SELFREPORTED"

	//DEVICE_UNKNOWN ...
	DEVICE_UNKNOWN UDIType = "DEVICE_UNKNOWN"
)

//DeviceSafety enum ...
type DeviceSafety string

const (
	//C106046 ...
	C106046 DeviceSafety = "C106046"

	//C106045 ...
	C106045 DeviceSafety = "C106045"

	//C106047 ...
	C106047 DeviceSafety = "C106047"

	//C113844 ...
	C113844 DeviceSafety = "C113844"

	//C101673 ...
	C101673 DeviceSafety = "C101673"

	//C106038 ...
	C106038 DeviceSafety = "C106038"
)

//DeviceStatus enum ...
type DeviceStatus string

const (
	//DEVICE_STATUS_ACTIVE ...
	DEVICE_STATUS_ACTIVE DeviceStatus = "DEVICE_STATUS_ACTIVE"

	//DEVICE_STATUS_INACTIVE ...
	DEVICE_STATUS_INACTIVE DeviceStatus = "DEVICE_STATUS_INACTIVE"

	//DEVICE_STATUS_ENTEREDINERROR ...
	DEVICE_STATUS_ENTEREDINERROR DeviceStatus = "DEVICE_STATUS_ENTEREDINERROR"

	//DEVICE_STATUS_UNKNOWN ...
	DEVICE_STATUS_UNKNOWN DeviceStatus = "DEVICE_STATUS_UNKNOWN"
)

//DeviceUniqueIdentifierInput ...
type DeviceUniqueIdentifierInput struct {
	DeviceIdentifier *string  `json:"deviceIdentifier"`
	Name             *string  `json:"name"`
	Jurisdiction     *string  `json:"jurisdiction"`
	CarrierCRF       *string  `json:"carrierCRF"`
	Issuer           *string  `json:"issuer"`
	EntryType        *UDIType `json:"entryType"`
}

//DeviceUniqueIdentifier ...
type DeviceUniqueIdentifier struct {
	DeviceIdentifier *string      `json:"deviceIdentifier" bson:"deviceIdentifier"`
	Name             *string      `json:"name" bson:"name"`
	Jurisdiction     *string      `json:"jurisdiction" bson:"jurisdiction"`
	CarrierCRF       *string      `json:"carrierCRF" bson:"carrierCRF"`
	Issuer           *string      `json:"issuer" bson:"issuer"`
	EntryType        *UDIType     `json:"entryType" bson:"entryType"`
	Meta             *models.Meta //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

//DeviceCreate ...
type DeviceCreate struct {
	Udi              *DeviceUniqueIdentifierInput `json:"udi"`
	Status           DeviceStatus                 `json:"status"`
	Type             string                       `json:"type"`
	TypeCode         *ClinicalCodeInput           `json:"typeCode"`
	LotNumber        *string                      `json:"lotNumber"`
	Manufacturer     *string                      `json:"manufacturer"`
	ManufacturerDate *util.Time                   `json:"manufacturerDate"`
	ExpirationDate   *util.Time                   `json:"expirationDate"`
	Model            *string                      `json:"model"`
	Version          *string                      `json:"version"`
	Consumer         ReferenceActorInput          `json:"consumer"`
	Owner            *ReferenceActorInput         `json:"owner"`
	Contact          *ContactPointInput           `json:"contact"`
}

//Device ...
type Device struct {
	Id               string                  `json:"id" bson:"_id"`
	Udi              *DeviceUniqueIdentifier `json:"udi" bson:"udi"`
	Status           DeviceStatus            `json:"status" bson:"status"`
	Type             string                  `json:"type" bson:"type"`
	TypeCode         *ClinicalCode           `json:"typeCode" bson:"typeCode"`
	LotNumber        *string                 `json:"lotNumber" bson:"lotNumber"`
	Manufacturer     *string                 `json:"manufacturer" bson:"manufacturer"`
	ManufacturerDate *util.Time              `json:"manufacturerDate" bson:"manufacturerDate"`
	ExpirationDate   *util.Time              `json:"expirationDate" bson:"expirationDate"`
	Model            *string                 `json:"model" bson:"model"`
	Version          *string                 `json:"version" bson:"version"`
	Consumer         ReferenceActor          `json:"consumer" bson:"consumer"`
	Owner            *ReferenceActor         `json:"owner" bson:"owner"`
	Contact          *ContactPoint           `json:"contact" bson:"contact"`
	Meta             *models.Meta            //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
