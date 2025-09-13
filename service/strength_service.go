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
StrengthService
==========================================================================================*/

// StrengthService ..
type StrengthService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

// NewStrengthService ..
func NewStrengthService(dal mserver.DataAccessLayer, log *logging.Logger) *StrengthService {
	return &StrengthService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

// FindById ..
func (u *StrengthService) FindById(id string) (*model.Strength, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching Strength (if any) from Mongo
	cct, err := u.dal.Get(id, &model.Strength{})
	if err != nil {
		return nil, err
	}

	//we need to unmarshal the result from type bson.M{} to model.Dosage
	//Convert BSON (byte) to JSON Fields
	var strength *model.Strength
	bsonBytes, _ := bson.Marshal(cct)
	bson.Unmarshal(bsonBytes, &strength)

	return strength, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

// CreateStrength will create a new Strength in Mongo using the Data Access Layer
func (u *StrengthService) CreateStrength(strength *model.Strength) (*model.Strength, error) {
	//Validate required fields on Model element are being passed in
	if strength.Unit == "" {
		return nil, errors.New("Missing a required field: aborting before saving to the DB")
	}

	if &strength.Number == nil {
		return nil, errors.New("Missing a required field: aborting before saving to the DB")
	}

	id, err := u.dal.Post(strength)
	if err != nil {
		return nil, err
	}

	strength.Id = id

	return strength, nil
}
