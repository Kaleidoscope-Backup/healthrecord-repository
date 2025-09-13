package service

import (
	"errors"

	"github.com/globalsign/mgo/bson"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/mongo-lib/mserver"
	logging "github.com/op/go-logging"
)

/*==========================================================================================
SourceRecordIDService
==========================================================================================*/

// SourceRecordIDService ..
type SourceRecordIDService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

// NewSourceRecordIDService ..
func NewSourceRecordIDService(dal mserver.DataAccessLayer, log *logging.Logger) *SourceRecordIDService {
	return &SourceRecordIDService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

// FindById ..
func (u *SourceRecordIDService) FindById(id string) (*model.SourceRecordID, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching source record (if any) from Mongo
	cct, err := u.dal.Get(id, &model.SourceRecordID{})
	if err != nil {
		return nil, err
	}

	//we need to unmarshal the result from type bson.M{} to model.Dosage
	//Convert BSON (byte) to JSON Fields
	var sourceRecordID *model.SourceRecordID
	bsonBytes, _ := bson.Marshal(cct)
	bson.Unmarshal(bsonBytes, &sourceRecordID)

	return sourceRecordID, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

// CreateSourceRecordID will create a new SourceRecordID in Mongo using the Data Access Layer
func (u *SourceRecordIDService) CreateSourceRecordID(sourceRecordID *model.SourceRecordID) (*model.SourceRecordID, error) {
	//Validate required fields on Model element are being passed in
	if sourceRecordID.System == "" {
		return nil, errors.New("Missing a required field: aborting before saving to the DB")
	}

	id, err := u.dal.Post(sourceRecordID)
	if err != nil {
		return nil, err
	}

	sourceRecordID.Id = id

	return sourceRecordID, nil
}
