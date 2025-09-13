package service

import (
	"errors"

	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/mongo-lib/mserver"
	"gopkg.in/mgo.v2/bson"
)

/*==========================================================================================
CodableConceptService
==========================================================================================*/

//CodableConceptService ...
type CodableConceptService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

//NewCodableConceptService ...
func NewCodableConceptService(dal mserver.DataAccessLayer, log *logging.Logger) *CodableConceptService {
	return &CodableConceptService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

//FindByID ..
func (u *CodableConceptService) FindByID(id string) (*model.CodableConcept, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching concept (if any) from Mongo
	conpt, err := u.dal.Get(id, &model.CodableConcept{})
	if err != nil {
		return nil, err
	}

	var concept *model.CodableConcept
	bsonBytes, _ := bson.Marshal(conpt)
	bson.Unmarshal(bsonBytes, &concept)

	return concept, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

//CreateCodableConcept will create a new CodableConcept in Mongo using the Data Access Layer
func (u *CodableConceptService) CreateCodableConcept(codableConcept *model.CodableConcept) (*model.CodableConcept, error) {

	_, err := u.dal.Post(codableConcept)
	if err != nil {
		return nil, err
	}

	return codableConcept, nil
}
