package service

import (
	"errors"

	"github.com/globalsign/mgo/bson"
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/model"

	"gitlab.com/karte/mongo-lib/mserver"
)

/*==========================================================================================
ContactPointService
==========================================================================================*/

//ContactPointService ..
type ContactPointService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

//NewContactPointService ..
func NewContactPointService(dal mserver.DataAccessLayer, log *logging.Logger) *ContactPointService {
	return &ContactPointService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

//FindById ..
func (u *ContactPointService) FindById(id string) (*model.ContactPoint, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching Dosage (if any) from Mongo
	cct, err := u.dal.Get(id, &model.ContactPoint{})
	if err != nil {
		return nil, err
	}

	//we need to unmarshal the result from type bson.M{} to model.Dosage
	//Convert BSON (byte) to JSON Fields
	var contactPoint *model.ContactPoint
	bsonBytes, _ := bson.Marshal(cct)
	bson.Unmarshal(bsonBytes, &contactPoint)

	return contactPoint, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

//CreateContactPoint will create a new ContactPoint in Mongo using the Data Access Layer
func (u *ContactPointService) CreateContactPoint(contactPoint *model.ContactPoint) (*model.ContactPoint, error) {
	//Validate required fields on Model element are being passed in
	if contactPoint.System == "" || contactPoint.Value == "" {
		return nil, errors.New("Missing a required field: aborting before saving to the DB")
	}

	id, err := u.dal.Post(contactPoint)
	if err != nil {
		return nil, err
	}

	contactPoint.Id = id

	return contactPoint, nil
}
