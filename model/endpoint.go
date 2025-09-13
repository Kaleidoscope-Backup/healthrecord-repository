package model

import "gitlab.com/karte/mongo-lib/models"

//EndpointStatus ...
type EndpointStatus string

const (
	//ENDPOINT_ACTIVE ...
	ENDPOINT_ACTIVE EndpointStatus = "ENDPOINT_ACTIVE"

	//ENDPOINT_SUSPENDED ...
	ENDPOINT_SUSPENDED EndpointStatus = "ENDPOINT_SUSPENDED"

	//ENDPOINT_ERROR ...
	ENDPOINT_ERROR EndpointStatus = "ENDPOINT_ERROR"

	//ENDPOINT_OFF ...
	ENDPOINT_OFF EndpointStatus = "ENDPOINT_OFF"

	//ENDPOINT_ENTERED_IN_ERROR ...
	ENDPOINT_ENTERED_IN_ERROR EndpointStatus = "ENDPOINT_ENTERED_IN_ERROR"

	//ENDPOINT_TEST ...
	ENDPOINT_TEST EndpointStatus = "ENDPOINT_TEST"
)

//EndpointConnectionType ...
type EndpointConnectionType string

const (
	//IHE_XCPD ...
	IHE_XCPD EndpointConnectionType = "IHE_XCPD"

	//IHE_XCA ...
	IHE_XCA EndpointConnectionType = "IHE_XCA"

	//IHE_XDR ...
	IHE_XDR EndpointConnectionType = "IHE_XDR"

	//IHE_XDS ...
	IHE_XDS EndpointConnectionType = "IHE_XDS"

	//IHE_IID ...
	IHE_IID EndpointConnectionType = "IHE_IID"

	//DICOM_WADO_RS ...
	DICOM_WADO_RS EndpointConnectionType = "DICOM_WADO_RS"

	//DICOM_QIDO_RS ...
	DICOM_QIDO_RS EndpointConnectionType = "DICOM_QIDO_RS"

	//DICOM_STOW_RS ...
	DICOM_STOW_RS EndpointConnectionType = "DICOM_STOW_RS"

	//DICOM_WADO_URI ...
	DICOM_WADO_URI EndpointConnectionType = "DICOM_WADO_URI"

	//HL7_FHIR_REST ...
	HL7_FHIR_REST EndpointConnectionType = "HL7_FHIR_REST"

	//HL7_FHIR_MSG ...
	HL7_FHIR_MSG EndpointConnectionType = "HL7_FHIR_MSG"

	//HL7V2_MLLP ...
	HL7V2_MLLP EndpointConnectionType = "HL7V2_MLLP"

	//SECURE_EMAIL ...
	SECURE_EMAIL EndpointConnectionType = "SECURE_EMAIL"

	//DIRECT_PROJECT ...
	DIRECT_PROJECT EndpointConnectionType = "DIRECT_PROJECT"
)

//EndpointInput ...
type EndpointInput struct {
	Status               EndpointStatus         `json:"status" bson:"status"`
	Name                 string                 `json:"name" bson:"name"`
	ConnectionType       EndpointConnectionType `json:"connectionType" bson:"connectionType"`
	ManagingOrganization ReferenceActorInput    `json:"managingOrganization" bson:"managingOrganization"`
	Contact              *[]ContactPointInput   `json:"contact" bson:"contact"`
	Period               *PeriodInput           `json:"period" bson:"period"`
	PayloadType          *[]string              `json:"payloadType" bson:"payloadType"`
	PayloadTypeCode      *[]CodableConceptInput `json:"payloadTypeCode" bson:"payloadTypeCode"`
	PlayloadMimeType     *[]string              `json:"playloadMimeType" bson:"playloadMimeType"`
	PlayloadMimeTypeCode *[]CodableConceptInput `json:"playloadMimeTypeCode" bson:"playloadMimeTypeCode"`
	Address              string                 `json:"address" bson:"address"`
	Header               *[]string              `json:"header" bson:"header"`
}

//Endpoint ...
type Endpoint struct {
	Id                   string                 `json:"id" bson:"_id"`
	Status               EndpointStatus         `json:"status" bson:"status"`
	Name                 string                 `json:"name" bson:"name"`
	ConnectionType       EndpointConnectionType `json:"connectionType" bson:"connectionType"`
	ManagingOrganization ReferenceActor         `json:"managingOrganization" bson:"managingOrganization"`
	Contact              *[]ContactPoint        `json:"contact" bson:"contact"`
	Period               *Period                `json:"period" bson:"period"`
	PayloadType          *[]string              `json:"payloadType" bson:"payloadType"`
	PayloadTypeCode      *[]CodableConcept      `json:"payloadTypeCode" bson:"payloadTypeCode"`
	PlayloadMimeType     *[]string              `json:"playloadMimeType" bson:"playloadMimeType"`
	PlayloadMimeTypeCode *[]CodableConcept      `json:"playloadMimeTypeCode" bson:"playloadMimeTypeCode"`
	Address              string                 `json:"address" bson:"address"`
	Header               *[]string              `json:"header" bson:"header"`
	Meta                 *models.Meta           //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
