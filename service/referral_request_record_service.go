package service

import (
	"errors"

	"github.com/globalsign/mgo/bson"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/mongo-lib/mserver"
	"github.com/op/go-logging"
)

/*==========================================================================================
ReferralRequestService
==========================================================================================*/

// ReferralRequestService ..
type ReferralRequestService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

// NewReferralRequestService ..
func NewReferralRequestService(dal mserver.DataAccessLayer, log *logging.Logger) *ReferralRequestService {
	return &ReferralRequestService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

// FindByID ..
func (u *ReferralRequestService) FindByID(id string) (*model.ReferralRequest, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching ReferralRequest (if any) from Mongo
	refReq, err := u.dal.Get(id, &model.ReferralRequest{})
	if err != nil {
		return nil, err
	}

	//we need to unmarshal the result from type bson.M{} to model.AppointmentRecord
	//Convert BSON (byte) to JSON Fields
	var referralRequest *model.ReferralRequest
	bsonBytes, _ := bson.Marshal(refReq)
	bson.Unmarshal(bsonBytes, &referralRequest)

	return referralRequest, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

// CreateReferralRequest will create a new ReferralRequest in Mongo using the Data Access Layer
func (u *ReferralRequestService) CreateReferralRequest(refRequest *model.ReferralRequest) (*model.ReferralRequest, error) {
	//Validate required fields on Model element are being passed in
	if &refRequest.Status == nil {
		return nil, errors.New("Missing a required field Status: aborting before saving to the DB")
	}

	if &refRequest.Subject == nil {
		return nil, errors.New("Missing a required field: Subject aborting before saving to the DB")
	}

	if &refRequest.Requester == nil {
		return nil, errors.New("Missing a required field: Requester aborting before saving to the DB")
	}

	if &refRequest.Recipient == nil {
		return nil, errors.New("Missing a required field: Recipient aborting before saving to the DB")
	}

	id, err := u.dal.Post(refRequest)
	if err != nil {
		return nil, err
	}

	refRequest.Id = id
	return refRequest, nil
}
