package model

import "github.com/Kaleidoscope-Backup/healthrecord-repository/util"

// ReferralRequestStatus ...
type ReferralRequestStatus string

const (
	//REFERRALREQUEST_DRAFT ..
	REFERRALREQUEST_DRAFT ReferralRequestStatus = "REFERRALREQUEST_DRAFT"
	//REFERRALREQUEST_ACTIVE ..
	REFERRALREQUEST_ACTIVE ReferralRequestStatus = "REFERRALREQUEST_ACTIVE"
	//REFERRALREQUEST_SUSPENDED ..
	REFERRALREQUEST_SUSPENDED ReferralRequestStatus = "REFERRALREQUEST_SUSPENDED"
	//REFERRALREQUEST_CANCELLED ..
	REFERRALREQUEST_CANCELLED ReferralRequestStatus = "REFERRALREQUEST_CANCELLED"
	//REFERRALREQUEST_COMPLETED ..
	REFERRALREQUEST_COMPLETED ReferralRequestStatus = "REFERRALREQUEST_COMPLETED"
	//REFERRALREQUEST_ENTERED_IN_ERROR ..
	REFERRALREQUEST_ENTERED_IN_ERROR ReferralRequestStatus = "REFERRALREQUEST_ENTERED_IN_ERROR"
	//REFERRALREQUEST_UNKNOWN ..
	REFERRALREQUEST_UNKNOWN ReferralRequestStatus = "REFERRALREQUEST_UNKNOWN"
)

// ReferralRequestCreate ...
type ReferralRequestCreate struct {
	BasedOn     *[]ReferenceHealthRecordInput `json:"basedOn"`
	Status      ReferralRequestStatus         `json:"status"`
	StatusCode  *ClinicalCodeInput            `json:"statusCode"`
	Subject     ReferenceActorInput           `json:"subject"`
	Requester   ReferenceActorInput           `json:"requester"`
	Recipient   ReferenceActorInput           `json:"recipient"`
	Description *string                       `json:"description"`
	Occurence   util.Time                     `json:"occurence"`
}

// ReferralRequest ...
type ReferralRequest struct {
	Id          string                   `json:"id" bson:"_id"`
	BasedOn     *[]ReferenceHealthRecord `json:"basedOn" bson:"basedOn"`
	Status      ReferralRequestStatus    `json:"status" bson:"status"`
	StatusCode  *ClinicalCode            `json:"statusCode" bson:"statusCode"`
	Subject     ReferenceActor           `json:"subject" bson:"subject"`
	Requester   ReferenceActor           `json:"requester" bson:"requester"`
	Recipient   ReferenceActor           `json:"recipient" bson:"recipient"`
	Description *string                  `json:"description" bson:"description"`
	Occurence   util.Time                `json:"occurence" bson:"occurence"`
}
