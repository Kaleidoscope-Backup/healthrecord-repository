package service

import (
	"errors"

	"github.com/globalsign/mgo/bson"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/mongo-lib/mserver"
	"github.com/op/go-logging"
)

/*==========================================================================================
AppointmentRecordService
==========================================================================================*/

// AppointmentRecordService ..
type AppointmentRecordService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

// NewAppointmentRecordService ..
func NewAppointmentRecordService(dal mserver.DataAccessLayer, log *logging.Logger) *SlotService {
	return &SlotService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

// FindByID ..
func (u *AppointmentRecordService) FindByID(id string) (*model.AppointmentRecord, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching appointment (if any) from Mongo
	appt, err := u.dal.Get(id, &model.AppointmentRecord{})
	if err != nil {
		return nil, err
	}

	//we need to unmarshal the result from type bson.M{} to model.AppointmentRecord
	//Convert BSON (byte) to JSON Fields
	var appointmentRecord *model.AppointmentRecord
	bsonBytes, _ := bson.Marshal(appt)
	bson.Unmarshal(bsonBytes, &appointmentRecord)

	return appointmentRecord, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

// CreateAppointmentRecord will create a new AppointmentRecord in Mongo using the Data Access Layer
func (u *AppointmentRecordService) CreateAppointmentRecord(appointment *model.AppointmentRecord) (*model.AppointmentRecord, error) {
	//Validate required fields on Model element are being passed in
	if &appointment.Status == nil {
		return nil, errors.New("Missing a required field Status: aborting before saving to the DB")
	}

	if &appointment.AppointmentType == nil {
		return nil, errors.New("Missing a required field: AppointmentType aborting before saving to the DB")
	}

	if &appointment.RequestedPeriod == nil {
		return nil, errors.New("Missing a required field: RequestedPeriod aborting before saving to the DB")
	}

	id, err := u.dal.Post(appointment)
	if err != nil {
		return nil, err
	}

	appointment.Id = id

	PostRecord(appointment, appointment.ConsumerID, id)
	return appointment, nil
}
