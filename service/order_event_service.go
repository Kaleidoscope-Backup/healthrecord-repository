package service

import (
	"errors"

	"github.com/globalsign/mgo/bson"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/mongo-lib/mserver"
	"github.com/op/go-logging"
)

/*==========================================================================================
OrderEventService
==========================================================================================*/

// OrderEventService ..
type OrderEventService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

// NewOrderEventService ..
func NewOrderEventService(dal mserver.DataAccessLayer, log *logging.Logger) *OrderEventService {
	return &OrderEventService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

// FindByID ..
func (u *OrderEventService) FindByID(id string) (*model.OrderEvent, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching order event (if any) from Mongo
	cct, err := u.dal.Get(id, &model.OrderEvent{})
	if err != nil {
		return nil, err
	}

	//we need to unmarshal the result from type bson.M{} to model.Dosage
	//Convert BSON (byte) to JSON Fields
	var orderEvent *model.OrderEvent
	bsonBytes, _ := bson.Marshal(cct)
	bson.Unmarshal(bsonBytes, &orderEvent)

	return orderEvent, nil
}

// FindEvents ...
func (u *OrderEventService) FindEvents(orderID string) (*[]*model.OrderEvent, error) {
	if orderID == "" {
		return nil, errors.New("Missing parameter id")
	}

	var params map[string]string
	params = map[string]string{}
	params["orderID"] = orderID

	//find the matching order id Record (if any) from Mongo
	prArr, err := FindRecords(&params, &model.OrderEvent{}, u.dal)
	if err != nil {
		return nil, err
	}

	var orderEventArr []*model.OrderEvent
	for _, pr := range prArr {
		var orderEvent *model.OrderEvent
		bsonBytes, _ := bson.Marshal(pr)
		bson.Unmarshal(bsonBytes, &orderEvent)
		orderEventArr = append(orderEventArr, orderEvent)
	}

	return &orderEventArr, nil
}

// FindEventsByParam ...
func (u *OrderEventService) FindByParam(param *model.OrderEventQueryParam) (*[]*model.OrderEvent, error) {
	if param == nil {
		return nil, errors.New("Missing parameter")
	}

	var params map[string]string
	params = map[string]string{}

	if param.ExternalID != nil && *param.ExternalID != "" {
		params["externalID"] = *param.ExternalID
	}

	if param.OrderID != nil && *param.OrderID != "" {
		params["orderID"] = *param.OrderID
	}

	if param.FromID != nil && *param.FromID != "" {
		params["fromID"] = *param.FromID
	}

	if param.ToID != nil && *param.ToID != "" {
		params["toID"] = *param.ToID
	}

	if param.Type != nil {
		params["type"] = string(*param.Type)
	}

	//find the matching order id Record (if any) from Mongo
	prArr, err := FindRecords(&params, &model.OrderEvent{}, u.dal)
	if err != nil {
		return nil, err
	}

	var orderEventArr []*model.OrderEvent
	for _, pr := range prArr {
		var orderEvent *model.OrderEvent
		bsonBytes, _ := bson.Marshal(pr)
		bson.Unmarshal(bsonBytes, &orderEvent)
		orderEventArr = append(orderEventArr, orderEvent)
	}

	return &orderEventArr, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

// CreateOrderEvent will create a new Product in Mongo using the Data Access Layer
func (u *OrderEventService) CreateOrderEvent(event *model.OrderEvent) (*model.OrderEvent, error) {

	id, err := u.dal.Post(event)
	if err != nil {
		return nil, err
	}

	event.Id = id
	return event, nil
}
