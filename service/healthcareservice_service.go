package service

import (
	"errors"

	"github.com/globalsign/mgo/bson"
	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/mongo-lib/mserver"
)

//HealthcareServiceService ...
type HealthcareServiceService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

//NewHealthcareServiceService ...
func NewHealthcareServiceService(dal mserver.DataAccessLayer, log *logging.Logger) *HealthcareServiceService {
	return &HealthcareServiceService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

//FindByID ..
func (u *HealthcareServiceService) FindByID(id string) (*model.HealthcareService, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching Record(if any) from Mongo
	p, err := u.dal.Get(id, &model.HealthcareService{})
	if err != nil {
		return nil, err
	}

	var healthcareService *model.HealthcareService
	bsonBytes, _ := bson.Marshal(p)
	bson.Unmarshal(bsonBytes, &healthcareService)

	return healthcareService, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

//CreateHealthcareService ...
func (u *HealthcareServiceService) CreateHealthcareService(healthcareService *model.HealthcareService) (*model.HealthcareService, error) {
	//Validate required fields on Model element are being passed in
	if healthcareService == nil {
		return nil, errors.New("Missing a required field: aborting before saving to the DB")
	}

	id, err := u.dal.Post(healthcareService)
	if err != nil {
		return nil, err
	}

	healthcareService.Id = id
	return healthcareService, nil
}
