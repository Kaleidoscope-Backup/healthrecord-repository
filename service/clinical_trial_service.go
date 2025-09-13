package service

import (
	"errors"

	"github.com/globalsign/mgo/bson"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/mongo-lib/mserver"
	"github.com/op/go-logging"
)

/*==========================================================================================
ClinicalTrial service
==========================================================================================*/

// ClinicalTrialService is for creating clinical trial
type ClinicalTrialService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

// NewClinicalTrialService creates a new ClinicalTrial service that has all calls to the database, queries and mutations via the Data Access Layer
func NewClinicalTrialService(dal mserver.DataAccessLayer, log *logging.Logger) *ClinicalTrialService {
	return &ClinicalTrialService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

// FindByID ..
func (u *ClinicalTrialService) FindByID(id string) (*model.ClinicalTrial, error) {

	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching Contact (if any) from Mongo
	p, err := u.dal.Get(id, &model.ClinicalTrial{})
	if err != nil {
		return nil, err
	}

	var clinicalTrial *model.ClinicalTrial
	bsonBytes, _ := bson.Marshal(p)
	bson.Unmarshal(bsonBytes, &clinicalTrial)

	return clinicalTrial, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

// CreateClinicalTrial will create a new clinical trial in Mongo using the Data Access Layer
func (u *ClinicalTrialService) CreateClinicalTrial(clinicalTrial *model.ClinicalTrial) (*model.ClinicalTrial, error) {
	_, err := u.dal.Post(clinicalTrial)
	if err != nil {
		return nil, err
	}

	return clinicalTrial, nil
}
