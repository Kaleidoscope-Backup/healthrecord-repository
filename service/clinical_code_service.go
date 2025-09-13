package service

import (
	"context"
	"errors"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/mongo-lib/mserver"
	"github.com/globalsign/mgo/bson"
	"github.com/op/go-logging"
)

/*==========================================================================================
ClinicalCodeService
==========================================================================================*/

// ClinicalCodeService ...
type ClinicalCodeService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

// NewClinicalCodeService ..
func NewClinicalCodeService(dal mserver.DataAccessLayer, log *logging.Logger) *ClinicalCodeService {
	return &ClinicalCodeService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

// FindByID ..
func (u *ClinicalCodeService) FindByID(id string) (*model.ClinicalCode, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching Clinical Code (if any) from Mongo
	cc, err := u.dal.Get(id, &model.ClinicalCode{})
	if err != nil {
		return nil, err
	}

	//we need to unmarshal the result from type bson.M{} to model.ClinicalCode
	//Convert BSON (byte) to JSON Fields
	var clinicalCode *model.ClinicalCode
	bsonBytes, _ := bson.Marshal(cc)
	bson.Unmarshal(bsonBytes, &clinicalCode)

	return clinicalCode, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

// CreateClinicalCode ...
func (u *ClinicalCodeService) CreateClinicalCode(ctx context.Context, clinicalCode *model.ClinicalCode) (*model.ClinicalCode, error) {

	//Validate required fields on Model element are being passed in
	if clinicalCode.Code == "" ||
		&clinicalCode.SystemType == nil ||
		clinicalCode.Definition == "" {
		return nil, errors.New("Missing a required field: aborting before saving to the DB")
	}

	_, err1 := u.dal.Post(clinicalCode)
	if err1 != nil {
		return nil, err1
	}

	return clinicalCode, nil
}
