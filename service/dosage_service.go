package service

import (
	"errors"

	"github.com/globalsign/mgo/bson"
	"github.com/karte/healthrecord-repository/model"
	"github.com/op/go-logging"

	// "github.com/karte/healthrecord-repository/service"

	"github.com/karte/mongo-lib/mserver"
)

/*==========================================================================================
DosageService
==========================================================================================*/

// DosageService ..
type DosageService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

// NewDosageService ..
func NewDosageService(dal mserver.DataAccessLayer, log *logging.Logger) *DosageService {
	return &DosageService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

// FindById ..
func (u *DosageService) FindById(id string) (*model.Dosage, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching Dosage (if any) from Mongo
	cct, err := u.dal.Get(id, &model.Dosage{})
	if err != nil {
		return nil, err
	}

	//we need to unmarshal the result from type bson.M{} to model.Dosage
	//Convert BSON (byte) to JSON Fields
	var dosage *model.Dosage
	bsonBytes, _ := bson.Marshal(cct)
	bson.Unmarshal(bsonBytes, &dosage)

	return dosage, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

// CreateDosage will create a new Dosage in Mongo using the Data Access Layer
func (u *DosageService) CreateDosage(dosage *model.Dosage) (*model.Dosage, error) {
	//Validate required fields on Model element are being passed in
	if dosage.Unit == "" {
		return nil, errors.New("Missing a required field: aborting before saving to the DB")
	}

	if dosage.Frequency == "" {
		return nil, errors.New("Missing a required field: aborting before saving to the DB")
	}

	if &dosage.Value == nil {
		return nil, errors.New("Missing a required field: aborting before saving to the DB")
	}

	id, err := u.dal.Post(dosage)
	if err != nil {
		return nil, err
	}

	dosage.Id = id

	return dosage, nil
}
