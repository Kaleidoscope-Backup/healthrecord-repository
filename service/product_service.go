package service

import (
	"errors"

	"github.com/globalsign/mgo/bson"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/mongo-lib/mserver"
	logging "github.com/op/go-logging"
)

/*==========================================================================================
ProductService
==========================================================================================*/

// ProductService ..
type ProductService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

// NewProductService ..
func NewProductService(dal mserver.DataAccessLayer, log *logging.Logger) *ProductService {
	return &ProductService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

// FindByID ..
func (u *ProductService) FindByID(id string) (*model.Product, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching product (if any) from Mongo
	cct, err := u.dal.Get(id, &model.Product{})
	if err != nil {
		return nil, err
	}

	//we need to unmarshal the result from type bson.M{} to model.Dosage
	//Convert BSON (byte) to JSON Fields
	var product *model.Product
	bsonBytes, _ := bson.Marshal(cct)
	bson.Unmarshal(bsonBytes, &product)

	return product, nil
}

// FindByParam ...
func (u *ProductService) FindByParam(param *model.ProductQueryParam) (*[]*model.Product, error) {
	if param == nil {
		return nil, errors.New("Missing parameter")
	}

	var params map[string]string
	params = map[string]string{}

	//TODO make this generic
	if param.Category != nil && *param.Category != "" {
		params["category"] = *param.Category
	}

	if param.Language != nil {
		params["language"] = string(*param.Language)
	}

	if param.Label != nil && *param.Label != "" {
		params["label"] = *param.Label
	}

	if param.Vendor != nil && *param.Vendor != "" {
		params["vendor"] = *param.Vendor
	}

	if param.Supplier != nil && *param.Supplier != "" {
		params["supplier"] = *param.Supplier
	}

	// if param.Sort != nil {
	// 	params["_sort"] =
	// }

	//find the matching product id Record (if any) from Mongo
	prArr, err := FindRecords(&params, &model.Product{}, u.dal)
	if err != nil {
		return nil, err
	}

	var productArr []*model.Product
	for _, pr := range prArr {
		var product *model.Product
		bsonBytes, _ := bson.Marshal(pr)
		bson.Unmarshal(bsonBytes, &product)
		productArr = append(productArr, product)
	}

	return &productArr, nil
}

// FindProducts ...
func (u *ProductService) FindProducts(supplierID string) (*[]*model.Product, error) {
	if supplierID == "" {
		return nil, errors.New("Missing parameter id")
	}

	var params map[string]string
	params = map[string]string{}
	params["supplier"] = supplierID

	//find the matching product Record (if any) from Mongo
	prArr, err := FindRecords(&params, &model.Product{}, u.dal)
	if err != nil {
		return nil, err
	}

	var productArr []*model.Product
	for _, pr := range prArr {
		var product *model.Product
		bsonBytes, _ := bson.Marshal(pr)
		bson.Unmarshal(bsonBytes, &product)
		productArr = append(productArr, product)
	}

	return &productArr, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

// UpdateProduct ...
func (u *ProductService) UpdateProduct(product *model.Product) (*model.Product, error) {

	if product != nil {
		_, err := u.dal.Put(product.Id, product)
		if err != nil {
			return nil, err
		}

		return product, nil
	}

	return nil, nil
}

// CreateProduct will create a new Product in Mongo using the Data Access Layer
func (u *ProductService) CreateProduct(product *model.Product) (*model.Product, error) {

	id, err := u.dal.Post(product)
	if err != nil {
		return nil, err
	}

	product.Id = id
	return product, nil
}
