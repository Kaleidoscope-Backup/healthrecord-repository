package model

//AppointmentType ...
type AppointmentType string

const (
	//APPOINTMENT_CHECKUP ..
	APPOINTMENT_CHECKUP AppointmentType = "APPOINTMENT_CHECKUP"
	//APPOINTMENT_EMERGENCY ..
	APPOINTMENT_EMERGENCY AppointmentType = "APPOINTMENT_EMERGENCY"
	//APPOINTMENT_FOLLOWUP ..
	APPOINTMENT_FOLLOWUP AppointmentType = "APPOINTMENT_FOLLOWUP"
	//APPOINTMENT_ROUTINE ..
	APPOINTMENT_ROUTINE AppointmentType = "APPOINTMENT_ROUTINE"
	//APPOINTMENT_WALKIN ..
	APPOINTMENT_WALKIN AppointmentType = "APPOINTMENT_WALKIN"
)

//AppointmentStatus ...
type AppointmentStatus string

const (
	//APPOINTMENT_PROPOSED ..
	APPOINTMENT_PROPOSED AppointmentStatus = "APPOINTMENT_PROPOSED"
	//APPOINTMENT_PENDING ..
	APPOINTMENT_PENDING AppointmentStatus = "APPOINTMENT_PENDING"
	//APPOINTMENT_BOOKED ..
	APPOINTMENT_BOOKED AppointmentStatus = "APPOINTMENT_BOOKED"
	//APPOINTMENT_ARRIVED ..
	APPOINTMENT_ARRIVED AppointmentStatus = "APPOINTMENT_ARRIVED"
	//APPOINTMENT_FULFILLED ..
	APPOINTMENT_FULFILLED AppointmentStatus = "APPOINTMENT_FULFILLED"
	//APPOINTMENT_CANELLED ..
	APPOINTMENT_CANELLED AppointmentStatus = "APPOINTMENT_CANELLED"
	//APPOINTMENT_NOSHOW ..
	APPOINTMENT_NOSHOW AppointmentStatus = "APPOINTMENT_NOSHOW"
	//APPOINTMENT_ENTERED_IN_ERROR ..
	APPOINTMENT_ENTERED_IN_ERROR AppointmentStatus = "APPOINTMENT_ENTERED_IN_ERROR"
)

//AppointmentRecordCreate ...
type AppointmentRecordCreate struct {
	HealthRecordCreate
	Status              AppointmentStatus             `json:"status"`
	StatusCode          *ClinicalCodeInput            `json:"statusCode"`
	Speciality          *[]string                     `json:"speciality"`
	SpecialityCode      *[]ClinicalCodeInput          `json:"specialityCode"`
	AppointmentType     AppointmentType               `json:"appointmentType"`
	ServiceCategory     *[]string                     `json:"serviceCategory"`
	ServiceCategoryCode *[]ClinicalCodeInput          `json:"serviceCategoryCode"`
	Reason              *[]string                     `json:"reason"`
	ReasonCode          *[]ClinicalCodeInput          `json:"reasonCode"`
	Indication          *[]ReferenceHealthRecordInput `json:"indication"`
	IncomingReferral    *[]ReferenceEntityInput       `json:"incomingReferral"`
	Priority            *Priority                     `json:"priority"`
	MinutesDuration     *int32                        `json:"minutesDuration"`
	Slot                *[]ReferenceEntityInput       `json:"slot"`
	Comment             *string                       `json:"comment"`
	Participants        *[]ReferenceActorInput        `json:"participants"`
	RequestedPeriod     PeriodInput                   `json:"requestedPeriod"`
}

//AppointmentRecord ...
type AppointmentRecord struct {
	HealthRecord
	Id                  string                   `json:"id" bson:"_id"`
	Status              AppointmentStatus        `json:"status" bson:"status"`
	StatusCode          *ClinicalCode            `json:"statusCode" bson:"statusCode"`
	Speciality          *[]string                `json:"speciality" bson:"speciality"`
	SpecialityCode      *[]ClinicalCode          `json:"specialityCode" bson:"specialityCode"`
	AppointmentType     AppointmentType          `json:"appointmentType" bson:"appointmentType"`
	ServiceCategory     *[]string                `json:"serviceCategory" bson:"serviceCategory"`
	ServiceCategoryCode *[]ClinicalCode          `json:"serviceCategoryCode" bson:"serviceCategoryCode"`
	Reason              *[]string                `json:"reason" bson:"reason"`
	ReasonCode          *[]ClinicalCode          `json:"reasonCode" bson:"reasonCode"`
	Indication          *[]ReferenceHealthRecord `json:"indication" bson:"indication"`
	IncomingReferral    *[]ReferenceEntity       `json:"incomingReferral" bson:"incomingReferral"`
	Priority            *Priority                `json:"priority" bson:"priority"`
	MinutesDuration     *int32                   `json:"minutesDuration" bson:"minutesDuration"`
	Slot                *[]ReferenceEntity       `json:"slot" bson:"slot"`
	Comment             *string                  `json:"comment" bson:"comment"`
	Participants        *[]ReferenceActor        `json:"participants" bson:"participants"`
	RequestedPeriod     Period                   `json:"requestedPeriod" bson:"requestedPeriod"`
}
