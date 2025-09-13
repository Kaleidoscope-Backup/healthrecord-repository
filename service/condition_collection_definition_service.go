package service

import (
	"errors"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/mongo-lib/mserver"
	"github.com/globalsign/mgo/bson"
	"github.com/op/go-logging"
)

/*==========================================================================================
ConditionDefinitionCollection service
==========================================================================================*/

// ConditionDefinitionCollectionService is for creating condition templates
type ConditionDefinitionCollectionService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

// NewConditionDefinitionCollectionService ...
func NewConditionDefinitionCollectionService(dal mserver.DataAccessLayer, log *logging.Logger) *ConditionDefinitionCollectionService {
	return &ConditionDefinitionCollectionService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

// FindByID ..
func (u *ConditionDefinitionCollectionService) FindByID(id string) (*model.ConditionDefinitionCollection, error) {

	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching Contact (if any) from Mongo
	p, err := u.dal.Get(id, &model.ConditionDefinitionCollection{})
	if err != nil {
		return nil, err
	}

	var conditionDefinition *model.ConditionDefinitionCollection
	bsonBytes, _ := bson.Marshal(p)
	bson.Unmarshal(bsonBytes, &conditionDefinition)

	return conditionDefinition, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

// CreateConditionDefinitionCollection will create a new condition template in Mongo using the Data Access Layer
func (u *ConditionDefinitionCollectionService) CreateConditionDefinitionCollection(conditionTemplate *model.ConditionDefinitionCollection) (*model.ConditionDefinitionCollection, error) {
	id, err := u.dal.Post(conditionTemplate)
	if err != nil {
		return nil, err
	}
	conditionTemplate.Id = id
	return conditionTemplate, nil
}
