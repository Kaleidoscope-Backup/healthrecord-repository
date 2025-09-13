package service

import (
	"errors"

	"github.com/globalsign/mgo/bson"
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/mongo-lib/mserver"
)

/*==========================================================================================
OrderBagService
==========================================================================================*/

//OrderBagService ..
type OrderBagService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

//NewOrderBagService ..
func NewOrderBagService(dal mserver.DataAccessLayer, log *logging.Logger) *OrderBagService {
	return &OrderBagService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

//FindByID ..
func (u *OrderBagService) FindByID(id string) (*model.OrderBag, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching order bag (if any) from Mongo
	cct, err := u.dal.Get(id, &model.OrderBag{})
	if err != nil {
		return nil, err
	}

	//we need to unmarshal the result from type bson.M{} to model.Dosage
	//Convert BSON (byte) to JSON Fields
	var orderBag *model.OrderBag
	bsonBytes, _ := bson.Marshal(cct)
	bson.Unmarshal(bsonBytes, &orderBag)

	return orderBag, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

//CreateOrderBag will create a new Order Bag in Mongo using the Data Access Layer
func (u *OrderBagService) CreateOrderBag(orderBag *model.OrderBag) (*model.OrderBag, error) {

	id, err := u.dal.Post(orderBag)
	if err != nil {
		return nil, err
	}

	orderBag.Id = id
	return orderBag, nil
}
