package model

//SlotStatus ...
type SlotStatus string

const (
	//SLOT_BUSY ..
	SLOT_BUSY SlotStatus = "SLOT_BUSY"
	//SLOT_FREE ..
	SLOT_FREE SlotStatus = "SLOT_FREE"
	//SLOT_BUSY_UNAVAILABLE ..
	SLOT_BUSY_UNAVAILABLE SlotStatus = "SLOT_BUSY_UNAVAILABLE"
	//SLOT_BUSY_TENTATIVE ..
	SLOT_BUSY_TENTATIVE SlotStatus = "SLOT_BUSY_TENTATIVE"
	//SLOT_ENTERED_IN_ERROR ..
	SLOT_ENTERED_IN_ERROR SlotStatus = "SLOT_ENTERED_IN_ERROR"
)

//SlotCreate ...
type SlotCreate struct {
	Status          SlotStatus           `json:"status"`
	StatusCode      *ClinicalCodeInput   `json:"statusCode"`
	OverBooked      *bool                `json:"overBooked"`
	Speciality      *[]string            `json:"speciality"`
	SpecialityCode  *[]ClinicalCodeInput `json:"specialityCode"`
	ServiceType     *[]string            `json:"serviceType" bson:"serviceType"`
	ServiceTypeCode *[]ClinicalCodeInput `json:"serviceTypeCode"`
	Schedule        ReferenceEntityInput `json:"schedule"`
	Comment         *string              `json:"comment" bson:"comment"`
	Period          *PeriodInput         `json:"period"`
}

//Slot ...
type Slot struct {
	Id              string          `json:"id" bson:"_id"`
	Status          SlotStatus      `json:"status" bson:"status"`
	StatusCode      *ClinicalCode   `json:"statusCode" bson:"statusCode"`
	OverBooked      *bool           `json:"overBooked" bson:"overBooked"`
	Speciality      *[]string       `json:"speciality" bson:"speciality"`
	SpecialityCode  *[]ClinicalCode `json:"specialityCode" bson:"specialityCode"`
	ServiceType     *[]string       `json:"serviceType" bson:"serviceType"`
	ServiceTypeCode *[]ClinicalCode `json:"serviceTypeCode" bson:"serviceTypeCode"`
	Schedule        ReferenceEntity `json:"schedule" bson:"schedule"`
	Comment         *string         `json:"comment" bson:"comment"`
	Period          *Period         `json:"period" bson:"period"`
}
