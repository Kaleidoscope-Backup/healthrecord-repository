package service

import (
	"errors"
	"strconv"

	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/mongo-lib/mserver"
	logging "github.com/op/go-logging"
	"gopkg.in/mgo.v2/bson"
)

/*==========================================================================================
RelationshipService
==========================================================================================*/

// RelationshipService ..
type RelationshipService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

// NewRelationshipService ..
func NewRelationshipService(dal mserver.DataAccessLayer, log *logging.Logger) *RelationshipService {
	return &RelationshipService{dal: dal, log: log}
}

// FindByID ...
func (u *RelationshipService) FindByID(id string) (*model.Relationship, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching rel (if any) from Mongo
	rel, err := u.dal.Get(id, &model.Relationship{})
	if err != nil {
		return nil, err
	}

	//we need to unmarshal the result from type bson.M{} to model.Notification
	//Convert BSON (byte) to JSON Fields
	var relation *model.Relationship
	bsonBytes, _ := bson.Marshal(rel)
	bson.Unmarshal(bsonBytes, &relation)

	return relation, nil
}

// FindByRelationshipParams ...
func (u *RelationshipService) FindByRelationshipParams(fromID *string, fromType *model.ActorType, toID *string, toType *model.ActorType, relType *model.RelationshipType, label *string) (*[]*model.Relationship, error) {
	if fromType == nil && toType == nil && fromID == nil && toID == nil {
		return nil, errors.New("Missing parameter - either from id, to id, from type, or to type should be provided")
	}

	var params map[string]string
	params = map[string]string{}

	params["active"] = "true"

	if relType != nil {
		params["type"] = string(*relType)
	}

	if label != nil {
		params["label"] = *label
	}

	if fromID != nil {
		params["fromID"] = *fromID
	}

	if toID != nil {
		params["toID"] = *toID
	}

	if fromType != nil {
		params["fromType"] = string(*fromType)
	}

	if toType != nil {
		params["toType"] = string(*toType)
	}

	// Latests will be shown first
	params["_sort:desc"] = "_lastUpdated"
	params["_count"] = strconv.Itoa(int(constant.MAX_RECORD_FETCH_COUNT))

	//find the matching relationship Record (if any) from Mongo
	relArr, err := FindRecords(&params, &model.Relationship{}, u.dal)
	if err != nil {
		return nil, err
	}

	var relationshipArr []*model.Relationship
	for _, relr := range relArr {
		var relationship *model.Relationship
		bsonBytes, _ := bson.Marshal(relr)
		bson.Unmarshal(bsonBytes, &relationship)
		relationshipArr = append(relationshipArr, relationship)
	}

	return &relationshipArr, nil
}

// CreateRelationship will create a new relationship in Mongo using the Data Access Layer
func (u *RelationshipService) CreateRelationship(relationship *model.Relationship) (*model.Relationship, error) {
	//Validate required fields on Model element are being passed in
	if &relationship.Active == nil ||
		&relationship.From == nil ||
		&relationship.To == nil ||
		relationship.Label == "" {
		return nil, errors.New("Missing a required field: aborting before saving to the DB")
	}

	id, err := u.dal.Post(relationship)
	if err != nil {
		return nil, err
	}

	relationship.Id = id
	return relationship, nil
}
