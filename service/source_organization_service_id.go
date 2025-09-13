package service

import (
	"errors"

	"github.com/globalsign/mgo/bson"
	"github.com/karte/healthrecord-repository/model"
	"github.com/op/go-logging"

	"github.com/karte/mongo-lib/mserver"
)

/*==========================================================================================
SourceOrganizationIDService
==========================================================================================*/

// SourceOrganizationIDService ..
type SourceOrganizationIDService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

// NewSourceOrganizationIDService ..
func NewSourceOrganizationIDService(dal mserver.DataAccessLayer, log *logging.Logger) *SourceOrganizationIDService {
	return &SourceOrganizationIDService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

// FindById ..
func (u *SourceOrganizationIDService) FindById(id string) (*model.SourceOrganizationID, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching source record (if any) from Mongo
	cct, err := u.dal.Get(id, &model.SourceOrganizationID{})
	if err != nil {
		return nil, err
	}

	//we need to unmarshal the result from type bson.M{} to model.Dosage
	//Convert BSON (byte) to JSON Fields
	var sourceOrganizationID *model.SourceOrganizationID
	bsonBytes, _ := bson.Marshal(cct)
	bson.Unmarshal(bsonBytes, &sourceOrganizationID)

	return sourceOrganizationID, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

// CreateSourceOrganizationID will create a new SourceRecordID in Mongo using the Data Access Layer
func (u *SourceOrganizationIDService) CreateSourceOrganizationID(sourceOrganizationID *model.SourceOrganizationID) (*model.SourceOrganizationID, error) {
	//Validate required fields on Model element are being passed in
	if sourceOrganizationID.SourceID == "" {
		return nil, errors.New("Missing a required field: aborting before saving to the DB")
	}

	id, err := u.dal.Post(sourceOrganizationID)
	if err != nil {
		return nil, err
	}

	sourceOrganizationID.Id = id

	return sourceOrganizationID, nil
}
