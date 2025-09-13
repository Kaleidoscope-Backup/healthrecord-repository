package service

import (
	"errors"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/mongo-lib/mserver"
	"github.com/globalsign/mgo/bson"
	logging "github.com/op/go-logging"
)

// ListService ...
type ListService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

// NewListService ...
func NewListService(dal mserver.DataAccessLayer, log *logging.Logger) *ListService {
	return &ListService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

// FindByID ...
func (u *ListService) FindByID(id string) (*model.List, error) {
	//find the matching list (if any) from Mongo
	listBytes, err := u.dal.Get(id, &model.List{})
	if err != nil {
		return nil, err
	}

	//we need to unmarshal the result from type bson.M{} to model.HeartRate
	//Convert BSON (byte) to JSON Fields
	var list *model.List
	bsonBytes, _ := bson.Marshal(listBytes)
	bson.Unmarshal(bsonBytes, &list)

	return list, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

// CreateList ...
func (u *ListService) CreateList(list *model.List) (*model.List, error) {

	if list == nil {
		return nil, errors.New("Input list cannot be nil")
	}
	id, err := u.dal.Post(list)
	if err != nil {
		return nil, err
	}

	list.Id = id
	return list, nil
}
