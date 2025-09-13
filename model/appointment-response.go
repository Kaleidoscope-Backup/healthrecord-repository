package model

import "github.com/Kaleidoscope-Backup/healthrecord-repository/util"

// AppointmentResponseStatus ...
type AppointmentResponseStatus string

const (
	//APPOINTMENT_RESPONSE_STATUS_ACCEPTED ..
	APPOINTMENT_RESPONSE_STATUS_ACCEPTED AppointmentResponseStatus = "APPOINTMENT_RESPONSE_STATUS_ACCEPTED"
	//APPOINTMENT_RESPONSE_STATUS_DECLINED ..
	APPOINTMENT_RESPONSE_STATUS_DECLINED AppointmentResponseStatus = "APPOINTMENT_RESPONSE_STATUS_DECLINED"
	//APPOINTMENT_RESPONSE_STATUS_TENTATIVE ..
	APPOINTMENT_RESPONSE_STATUS_TENTATIVE AppointmentResponseStatus = "APPOINTMENT_RESPONSE_STATUS_TENTATIVE"
	//APPOINTMENT_RESPONSE_STATUS_INPROCESS ..
	APPOINTMENT_RESPONSE_STATUS_INPROCESS AppointmentResponseStatus = "APPOINTMENT_RESPONSE_STATUS_INPROCESS"
	//APPOINTMENT_RESPONSE_STATUS_COMPLETED ..
	APPOINTMENT_RESPONSE_STATUS_COMPLETED AppointmentResponseStatus = "APPOINTMENT_RESPONSE_STATUS_COMPLETED"
	//APPOINTMENT_RESPONSE_STATUS_NEEDSACTION ..
	APPOINTMENT_RESPONSE_STATUS_NEEDSACTION AppointmentResponseStatus = "APPOINTMENT_RESPONSE_STATUS_NEEDSACTION"
	//APPOINTMENT_RESPONSE_STATUS_ENTEREDINERROR ..
	APPOINTMENT_RESPONSE_STATUS_ENTEREDINERROR AppointmentResponseStatus = "APPOINTMENT_RESPONSE_STATUS_ENTEREDINERROR"
)

// AppointmentResponseCreate ...
type AppointmentResponseCreate struct {
	Status      AppointmentResponseStatus `json:"status"`
	Appointment ReferenceEntityInput      `json:"appointment"`
	Start       util.Time                 `json:"start"`
	End         util.Time                 `json:"end"`
	Actor       ReferenceActorInput       `json:"actor"`
	Comment     *string                   `json:"comment" bson:"comment"`
}

// AppointmentResponse ...
type AppointmentResponse struct {
	Id          string                    `json:"id" bson:"_id"`
	Status      AppointmentResponseStatus `json:"status" bson:"status"`
	Appointment ReferenceEntity           `json:"appointment" bson:"appointment"`
	Start       util.Time                 `json:"start" bson:"start"`
	End         util.Time                 `json:"end" bson:"end"`
	Actor       ReferenceActor            `json:"actor" bson:"actor"`
	Comment     *string                   `json:"comment" bson:"comment"`
}
