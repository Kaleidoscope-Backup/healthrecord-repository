package service

import (
	"errors"
	"fmt"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/mongo-lib/mserver"
	"github.com/op/go-logging"
	"gopkg.in/mgo.v2/bson"
)

/*==========================================================================================
ObservationDefinitionCollection  service
==========================================================================================*/

// ObservationDefinitionCollectionService is for creating habit
type ObservationDefinitionCollectionService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

// NewObservationDefinitionCollectionService Data Access Layer
func NewObservationDefinitionCollectionService(dal mserver.DataAccessLayer, log *logging.Logger) *ObservationDefinitionCollectionService {
	return &ObservationDefinitionCollectionService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

// FindByID ..
func (u *ObservationDefinitionCollectionService) FindByID(id string) (*model.ObservationDefinitionCollection, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching ObservationDefinitionCollection (if any) from Mongo
	obs, err := u.dal.Get(id, &model.ObservationDefinitionCollection{})
	if err != nil {
		return nil, err
	}

	var obsCollectDef *model.ObservationDefinitionCollection
	bsonBytes, _ := bson.Marshal(obs)
	bson.Unmarshal(bsonBytes, &obsCollectDef)

	return obsCollectDef, nil
}

// FindObservationDefinitionCollections ...
func (u *ObservationDefinitionCollectionService) FindObservationDefinitionCollections(param *model.ObservationDefinitionCollectionQueryParam) (*[]*model.ObservationDefinitionCollection, error) {
	if param == nil {
		return nil, errors.New("Missing parameter")
	}

	var params map[string]string
	params = map[string]string{}

	if param.Name != nil {
		params["name"] = *param.Name
	}

	if param.Publisher != nil {
		params["publisher"] = *param.Publisher
	}

	if param.Language != nil {
		params["language"] = string(*param.Language)
	}

	//find the matching order id Record (if any) from Mongo
	oArr, err := FindRecords(&params, &model.ObservationDefinitionCollection{}, u.dal)
	if err != nil {
		return nil, err
	}

	var obsArr []*model.ObservationDefinitionCollection
	for _, o := range oArr {
		var obs *model.ObservationDefinitionCollection
		bsonBytes, _ := bson.Marshal(o)
		bson.Unmarshal(bsonBytes, &obs)
		obsArr = append(obsArr, obs)
	}

	return &obsArr, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

// CreateObservationDefinitionCollection will create a new record in Mongo using the Data Access Layer ...
func (u *ObservationDefinitionCollectionService) CreateObservationDefinitionCollection(obsCollectDef *model.ObservationDefinitionCollection) (*model.ObservationDefinitionCollection, error) {

	if obsCollectDef.Name == "" || &obsCollectDef.Language == nil {
		return nil, errors.New("Missing a required field aborting before saving to the DB")
	}

	id, err := u.dal.Post(obsCollectDef)
	if err != nil {
		fmt.Printf("DB Create Failed For ObservationDefinitionCollection Error : %v", err)
		return nil, err
	}

	obsCollectDef.Id = id
	return obsCollectDef, nil
}
