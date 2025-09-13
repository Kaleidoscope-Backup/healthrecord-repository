package model

import "github.com/karte/mongo-lib/models"

// ApplicationMetadataType meta data types
type ApplicationMetadataType string

const (
	//APPMETADATA_PROFILE variables ...
	APPMETADATA_PROFILE ApplicationMetadataType = "APPMETADATA_PROFILE"

	//APPMETADATA_QUESTIONNAIRE variables ...
	APPMETADATA_QUESTIONNAIRE ApplicationMetadataType = "APPMETADATA_QUESTIONNAIRE"

	//APPMETADATA_RECOMMENDATION variables ...
	APPMETADATA_RECOMMENDATION ApplicationMetadataType = "APPMETADATA_RECOMMENDATION"
)

// ApplicationType enum ...
type ApplicationType string

const (
	//MOBILE_NATIVE_APP ...
	MOBILE_NATIVE_APP ApplicationType = "MOBILE_NATIVE_APP"
	//MOBILE_HYBRID_APP ...
	MOBILE_HYBRID_APP ApplicationType = "MOBILE_HYBRID_APP"
	//WEB_APP ...
	WEB_APP ApplicationType = "WEB_APP"
)

// ApplicationCreate ...
type ApplicationCreate struct {
	Name            string              `json:"name" bson:"name"`
	CallbackURL     *string             `json:"callbackURL"`
	SupportEmail    *string             `json:"supportEmail"`
	DefaultLanguage *Language           `json:"defaultLanguage"`
	Logo            *string             `json:"logo" bson:"logo"`
	Description     *string             `json:"description" bson:"description"`
	Owner           ReferenceActorInput `json:"owner" bson:"owner"`
	Type            *ApplicationType    `json:"type" bson:"type"`
	Attributes      *[]MetaDataInput    `json:"attributes" bson:"attributes"`
}

// Application ...
type Application struct {
	Id              string           `json:"id" bson:"_id"`
	Name            string           `json:"name" bson:"name"`
	CallbackURL     *string          `json:"callbackURL" bson:"callbackURL"`
	SupportEmail    *string          `json:"supportEmail" bson:"supportEmail"`
	DefaultLanguage *Language        `json:"defaultLanguage" bson:"defaultLanguage"`
	Logo            *string          `json:"logo" bson:"logo"`
	Description     *string          `json:"description" bson:"description"`
	Owner           ReferenceActor   `json:"owner" bson:"owner"`
	Type            *ApplicationType `json:"type" bson:"type"`
	Attributes      *[]MetaData      `json:"attributes" bson:"attributes"`
	Meta            *models.Meta     //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}

// ApplicationProfile ...
type ApplicationProfile struct {
	Name       string      `json:"name"`
	Value      string      `json:"value"`
	Attributes *[]MetaData `json:"attributes"`
}
