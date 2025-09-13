package service

import (
	"errors"
	"strconv"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
	"github.com/Kaleidoscope-Backup/mongo-lib/mserver"
	"github.com/globalsign/mgo/bson"
	logging "github.com/op/go-logging"
)

/*==========================================================================================
OrderService
==========================================================================================*/

// OrderService ..
type OrderService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

// NewOrderService ..
func NewOrderService(dal mserver.DataAccessLayer, log *logging.Logger) *OrderService {
	return &OrderService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

// FindByID ..
func (u *OrderService) FindByID(id string) (*model.Order, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching order  (if any) from Mongo
	cct, err := u.dal.Get(id, &model.Order{})
	if err != nil {
		return nil, err
	}

	//we need to unmarshal the result from type bson.M{} to model.Order
	//Convert BSON (byte) to JSON Fields
	var order *model.Order
	bsonBytes, _ := bson.Marshal(cct)
	bson.Unmarshal(bsonBytes, &order)

	return order, nil
}

// FindByConsumerID ...
func (u *OrderService) FindByConsumerID(id string) (*[]*model.Order, error) {
	if id == "" {
		return nil, errors.New("Missing parameter")
	}

	orders := []*model.Order{}
	paramFrom := model.OrderQueryParam{}
	paramFrom.FromID = id
	ordersFrom, err := u.FindOrders(&paramFrom)

	if err != nil {
		return nil, err
	}

	if ordersFrom != nil {
		orders = append(orders, *ordersFrom...)
	}

	paramTo := model.OrderQueryParam{}
	paramFrom.ToID = &id
	ordersTo, err := u.FindOrders(&paramTo)

	if err != nil {
		return nil, err
	}

	if ordersTo != nil {
		orders = append(orders, *ordersTo...)
	}
	return &orders, nil
}

// FindOrders ...
func (u *OrderService) FindOrders(param *model.OrderQueryParam) (*[]*model.Order, error) {
	if param == nil {
		return nil, errors.New("Missing parameter")
	}

	var params map[string]string
	params = map[string]string{}

	if param.Status != nil {
		params["status"] = string(*param.Status)
	}

	if param.FromID != "" {
		params["fromID"] = param.FromID
	}

	if param.ToID != nil && *param.ToID != "" {
		params["toID"] = *param.ToID
	}

	if param.Supplier != nil && *param.Supplier != "" {
		params["supplier"] = *param.Supplier
	}

	if param.OrderedItem != nil && *param.OrderedItem != "" {
		params["orderedItem"] = *param.OrderedItem
	}

	// Latests will be shown first
	params["_count"] = strconv.Itoa(int(constant.MAX_RECORD_FETCH_COUNT))

	//find the matching order id Record (if any) from Mongo
	orArr, err := FindRecords(&params, &model.Order{}, u.dal)
	if err != nil {
		return nil, err
	}

	var orderArr []*model.Order
	for _, or := range orArr {
		var order *model.Order
		bsonBytes, _ := bson.Marshal(or)
		bson.Unmarshal(bsonBytes, &order)
		orderArr = append(orderArr, order)
	}

	return &orderArr, nil
}

// FindOrdersBySupplier ...
func (u *OrderService) FindOrdersBySupplier(supplierID string) (*[]*model.Order, error) {
	if supplierID == "" {
		return nil, errors.New("Missing parameter id")
	}

	var params map[string]string
	params = map[string]string{}

	params["supplier"] = supplierID
	params["_count"] = strconv.Itoa(int(constant.MAX_RECORD_FETCH_COUNT))

	//find the matching order id Record (if any) from Mongo
	orArr, err := FindRecords(&params, &model.Order{}, u.dal)
	if err != nil {
		return nil, err
	}

	var orderArr []*model.Order
	for _, or := range orArr {
		var order *model.Order
		bsonBytes, _ := bson.Marshal(or)
		bson.Unmarshal(bsonBytes, &order)
		orderArr = append(orderArr, order)
	}

	return &orderArr, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

// CreateOrder will create a new Product in Mongo using the Data Access Layer
func (u *OrderService) CreateOrder(order *model.Order) (*model.Order, error) {

	id, err := u.dal.Post(order)
	if err != nil {
		return nil, err
	}

	order.Id = id
	return order, nil
}

// Export ...
func (u *OrderService) Export(id string) (*[]model.HealthRecordExportElement, error) {
	if id == "" {
		return nil, errors.New("Missing parameter")
	}

	orders, _ := u.FindByConsumerID(id)
	if orders != nil {
		ordersIterable := *orders
		records := []model.HealthRecordExportElement{}
		for _, order := range ordersIterable {
			recordElements, _ := u.ExportElements(order, id)
			records = append(records, *recordElements...)
		}
		return &records, nil
	}

	return nil, nil
}

// ExportElements ...
func (u *OrderService) ExportElements(order *model.Order, consumerID string) (*[]model.HealthRecordExportElement, error) {
	if order == nil {
		return nil, errors.New("Missing parameter")
	}

	// all records array
	records := []model.HealthRecordExportElement{}
	recordID := util.UUID()

	// Supplier
	valueSupplier := model.Value{}
	valueSupplier.ValueType = model.TEXT
	valueSupplier.ValueText = &order.Supplier

	// populate record
	recordSupplier := model.HealthRecordExportElement{}
	recordSupplier.Name = "Order Supplier"
	recordSupplier.TimeStamp = order.TimeStamp
	recordSupplier.Value = valueSupplier
	recordSupplier.RecordID = recordID
	recordSupplier.ConsumerID = consumerID
	//populate the record
	records = append(records, recordSupplier)

	// Ordered item
	valueOrderedItem := model.Value{}
	valueOrderedItem.ValueType = model.TEXT
	valueOrderedItem.ValueText = &order.OrderedItem

	// populate record
	recordOrderedItem := model.HealthRecordExportElement{}
	recordOrderedItem.Name = "Ordered Item"
	recordOrderedItem.TimeStamp = order.TimeStamp
	recordOrderedItem.Value = valueOrderedItem
	recordOrderedItem.RecordID = recordID
	recordOrderedItem.ConsumerID = consumerID
	//populate the record
	records = append(records, recordOrderedItem)

	// Total price
	valueTotalPrice := model.Value{}
	valueTotalPrice.ValueType = model.QUANTITY
	valueTotalPrice.ValueQuantity = &order.Quantity

	// populate record
	recordTotalPrice := model.HealthRecordExportElement{}
	recordTotalPrice.Name = "Order Price"
	recordTotalPrice.TimeStamp = order.TimeStamp
	recordTotalPrice.Value = valueTotalPrice
	recordTotalPrice.RecordID = recordID
	recordTotalPrice.ConsumerID = consumerID
	//populate the record
	records = append(records, recordTotalPrice)

	if order.Attributes != nil {
		attributes := *order.Attributes

		for _, attribute := range attributes {
			//value
			value := model.Value{}
			value.ValueType = model.TEXT
			value.ValueText = attribute.Value.ValueText
			value.Unit = attribute.Value.Unit

			// populate record
			record := model.HealthRecordExportElement{}
			record.Name = attribute.Name
			record.TimeStamp = order.TimeStamp
			record.Value = value
			record.RecordID = recordID
			record.ConsumerID = consumerID

			//populate the record
			records = append(records, record)
		}
	}

	return &records, nil
}
