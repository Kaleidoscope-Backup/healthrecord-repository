package service

import (
	"errors"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/globalsign/mgo/bson"
	"github.com/op/go-logging"

	// "github.com/Kaleidoscope-Backup/healthrecord-repository/service"

	"github.com/Kaleidoscope-Backup/mongo-lib/mserver"
)

/*==========================================================================================
SourceConsumerIDService
==========================================================================================*/

// SourceConsumerIDService ..
type SourceConsumerIDService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

// NewSourceConsumerIDService ..
func NewSourceConsumerIDService(dal mserver.DataAccessLayer, log *logging.Logger) *SourceConsumerIDService {
	return &SourceConsumerIDService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

// FindById ..
func (u *SourceConsumerIDService) FindById(id string) (*model.SourceConsumerID, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching source record (if any) from Mongo
	cct, err := u.dal.Get(id, &model.SourceConsumerID{})
	if err != nil {
		return nil, err
	}

	//we need to unmarshal the result from type bson.M{} to model.Dosage
	//Convert BSON (byte) to JSON Fields
	var sourceConsumerID *model.SourceConsumerID
	bsonBytes, _ := bson.Marshal(cct)
	bson.Unmarshal(bsonBytes, &sourceConsumerID)

	return sourceConsumerID, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

// CreateSourceConsumerID will create a new SourceConsumerID in Mongo using the Data Access Layer
func (u *SourceConsumerIDService) CreateSourceConsumerID(sourceConsumerID *model.SourceConsumerID) (*model.SourceConsumerID, error) {
	//Validate required fields on Model element are being passed in
	if sourceConsumerID.System == "" {
		return nil, errors.New("Missing a required field: aborting before saving to the DB")
	}

	id, err := u.dal.Post(sourceConsumerID)
	if err != nil {
		return nil, err
	}

	sourceConsumerID.Id = id

	return sourceConsumerID, nil
}
