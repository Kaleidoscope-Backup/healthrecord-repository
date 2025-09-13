package service

import (
	"errors"

	"github.com/globalsign/mgo/bson"
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/mongo-lib/mserver"
)

/*==========================================================================================
AcknowledgementService
==========================================================================================*/

//AcknowledgementService ..
type AcknowledgementService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

//NewAcknowledgementService ..
func NewAcknowledgementService(dal mserver.DataAccessLayer, log *logging.Logger) *AcknowledgementService {
	return &AcknowledgementService{dal: dal, log: log}
}

//FindByID ...
func (u *AcknowledgementService) FindByID(id string) (*model.Acknowledgement, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching Acknowledgement (if any) from Mongo
	cct, err := u.dal.Get(id, &model.Acknowledgement{})
	if err != nil {
		return nil, err
	}

	//we need to unmarshal the result from type bson.M{} to model.Acknowledgement
	//Convert BSON (byte) to JSON Fields
	var ack *model.Acknowledgement
	bsonBytes, _ := bson.Marshal(cct)
	bson.Unmarshal(bsonBytes, &ack)

	return ack, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

//CreateAcknowledgement will create a new ack in Mongo using the Data Access Layer
func (u *AcknowledgementService) CreateAcknowledgement(ack *model.Acknowledgement) (*model.Acknowledgement, error) {
	//Validate required fields on Model element are being passed in
	if &ack.Created == nil ||
		ack.ConsumerID == "" ||
		ack.RefrenceNotification == "" {
		return nil, errors.New("Missing a required field: aborting before saving to the DB")
	}

	id, err := u.dal.Post(ack)
	if err != nil {
		return nil, err
	}

	ack.Id = id
	return ack, nil
}
