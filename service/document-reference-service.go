package service

import (
	"errors"

	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/mongo-lib/mserver"
	"gopkg.in/mgo.v2/bson"
)

/*==========================================================================================
DocumentReference service
==========================================================================================*/

//DocumentReferenceService is for creating document reference
type DocumentReferenceService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

//NewDeviceService creates a new device
func NewDocumentReferenceService(dal mserver.DataAccessLayer, log *logging.Logger) *DocumentReferenceService {
	return &DocumentReferenceService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

//FindByID ..
func (u *DocumentReferenceService) FindByID(id string) (*model.DocumentReference, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching consent (if any) from Mongo
	d, err := u.dal.Get(id, &model.DocumentReference{})
	if err != nil {
		return nil, err
	}

	var docReference *model.DocumentReference
	bsonBytes, _ := bson.Marshal(d)
	bson.Unmarshal(bsonBytes, &docReference)

	return docReference, nil
}

//FindByParam ...
func (u *DocumentReferenceService) FindByParam(param *model.DocumentReferenceQueryParam) (*[]*model.DocumentReference, error) {
	if param == nil {
		return nil, errors.New("Missing parameter")
	}

	var params map[string]string
	params = map[string]string{}

	if param.Class != "" {
		params["class"] = param.Class
	}

	if &param.Language != nil {
		params["language"] = string(param.Language)
	}

	if param.Type != nil && *param.Type != "" {
		params["type"] = *param.Type
	}

	if param.Custodian != "" {
		params["custodian"] = param.Custodian
	}

	//find the matching product id Record (if any) from Mongo
	drArr, err := FindRecords(&params, &model.DocumentReference{}, u.dal)
	if err != nil {
		return nil, err
	}

	var documentReferenceArr []*model.DocumentReference
	for _, dr := range drArr {
		var documentReference *model.DocumentReference
		bsonBytes, _ := bson.Marshal(dr)
		bson.Unmarshal(bsonBytes, &documentReference)
		documentReferenceArr = append(documentReferenceArr, documentReference)
	}

	return &documentReferenceArr, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

//CreateDocumentReference ...
func (u *DocumentReferenceService) CreateDocumentReference(docReference *model.DocumentReference) (*model.DocumentReference, error) {

	id, err := u.dal.Post(docReference)
	if err != nil {
		return nil, err
	}

	docReference.Id = id
	return docReference, nil
}
