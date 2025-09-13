package service

import (
	"errors"

	"github.com/globalsign/mgo/bson"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/mongo-lib/mserver"
	logging "github.com/op/go-logging"
)

/*==========================================================================================
ConceptClass service
==========================================================================================*/

// ConceptClassService is for creating clinical trial
type ConceptClassService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

// NewConceptClassService creates a new ConceptClass service that has all calls to the database, queries and mutations via the Data Access Layer
func NewConceptClassService(dal mserver.DataAccessLayer, log *logging.Logger) *ConceptClassService {
	return &ConceptClassService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

// FindByID ..
func (u *ConceptClassService) FindByID(id string) (*model.ConceptClass, error) {

	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching Contact (if any) from Mongo
	c, err := u.dal.Get(id, &model.ConceptClass{})
	if err != nil {
		return nil, err
	}

	var concept *model.ConceptClass
	bsonBytes, _ := bson.Marshal(c)
	bson.Unmarshal(bsonBytes, &concept)

	return concept, nil
}

// FindByParam ...
func (u *ConceptClassService) FindByParam(param *model.ConceptClassQueryParam) (*[]*model.ConceptClass, error) {
	if param == nil {
		return nil, errors.New("Missing parameter")
	}

	if param.All != nil && *param.All == true {
		ccArr, err := u.dal.GetAll(&model.ConceptClass{})
		if err != nil {
			return nil, err
		}

		var conceptClassArr []*model.ConceptClass
		for _, pr := range ccArr {
			var conceptClass *model.ConceptClass
			bsonBytes, _ := bson.Marshal(pr)
			bson.Unmarshal(bsonBytes, &conceptClass)
			conceptClassArr = append(conceptClassArr, conceptClass)
		}

		return &conceptClassArr, nil

	}

	var params map[string]string
	params = map[string]string{}

	if param.Name != nil && *param.Name != "" {
		params["name"] = *param.Name
	}

	if param.ExternalID != nil && *param.ExternalID != "" {
		params["externalID"] = string(*param.ExternalID)
	}

	//find the matching product id Record (if any) from Mongo
	ccArr, err := FindRecords(&params, &model.ConceptClass{}, u.dal)
	if err != nil {
		return nil, err
	}

	var conceptClassArr []*model.ConceptClass
	for _, pr := range ccArr {
		var conceptClass *model.ConceptClass
		bsonBytes, _ := bson.Marshal(pr)
		bson.Unmarshal(bsonBytes, &conceptClass)
		conceptClassArr = append(conceptClassArr, conceptClass)
	}

	return &conceptClassArr, nil

}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

// CreateConceptClass will create a new clinical trial in Mongo using the Data Access Layer
func (u *ConceptClassService) CreateConceptClass(concept *model.ConceptClass) (*model.ConceptClass, error) {
	id, err := u.dal.Post(concept)
	if err != nil {
		return nil, err
	}

	concept.Id = id
	return concept, nil
}
