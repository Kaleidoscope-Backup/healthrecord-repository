package service

import (
	"errors"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/mongo-lib/mserver"
	"github.com/globalsign/mgo/bson"
	"github.com/op/go-logging"
)

/*==========================================================================================
AppointmentResponseService
==========================================================================================*/

// AppointmentResponseService ..
type AppointmentResponseService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

// NewAppointmentResponseService ..
func NewAppointmentResponseService(dal mserver.DataAccessLayer, log *logging.Logger) *AppointmentResponseService {
	return &AppointmentResponseService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

// FindByID ..
func (u *AppointmentResponseService) FindByID(id string) (*model.AppointmentResponse, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching AppointmentResponse (if any) from Mongo
	appResponse, err := u.dal.Get(id, &model.AppointmentResponse{})
	if err != nil {
		return nil, err
	}

	//we need to unmarshal the result from type bson.M{} to model.AppointmentResponse
	//Convert BSON (byte) to JSON Fields
	var appointmentResponse *model.AppointmentResponse
	bsonBytes, _ := bson.Marshal(appResponse)
	bson.Unmarshal(bsonBytes, &appointmentResponse)

	return appointmentResponse, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

// CreateAppointmentResponse will create a new AppointmentResponse in Mongo using the Data Access Layer
func (u *AppointmentResponseService) CreateAppointmentResponse(appRsp *model.AppointmentResponse) (*model.AppointmentResponse, error) {
	//Validate required fields on Model element are being passed in
	if &appRsp.Status == nil {
		return nil, errors.New("Missing a required field Status: aborting before saving to the DB")
	}

	if &appRsp.Start == nil {
		return nil, errors.New("Missing a required field Start: aborting before saving to the DB")
	}

	if &appRsp.End == nil {
		return nil, errors.New("Missing a required field: End aborting before saving to the DB")
	}

	if &appRsp.Appointment == nil {
		return nil, errors.New("Missing a required field: Appointment aborting before saving to the DB")
	}

	if &appRsp.Actor == nil {
		return nil, errors.New("Missing a required field: Actor aborting before saving to the DB")
	}

	id, err := u.dal.Post(appRsp)
	if err != nil {
		return nil, err
	}

	appRsp.Id = id
	return appRsp, nil
}
