package service

import (
	"github.com/globalsign/mgo/bson"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/mongo-lib/mserver"
	"github.com/op/go-logging"
)

/*==========================================================================================
Contact service
==========================================================================================*/

// ContactService is for creating contact
type ContactService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

// NewContactervice creates a new Contact service that has all calls to the database, queries and mutations via the Data Access Layer
func NewContactervice(dal mserver.DataAccessLayer, log *logging.Logger) *ContactService {
	return &ContactService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

// FindByID ..
func (u *ContactService) FindByID(id string) (*model.Contact, error) {
	//find the matching Contact (if any) from Mongo
	p, err := u.dal.Get(id, "Contact")
	if err != nil {
		return nil, err
	}

	var contact *model.Contact
	bsonBytes, _ := bson.Marshal(p)
	bson.Unmarshal(bsonBytes, &contact)

	return contact, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

// CreateContact will create a new geo location in Mongo using the Data Access Layer
func (u *ContactService) CreateContact(contact *model.Contact) (*model.Contact, error) {
	_, err := u.dal.Post(contact)
	if err != nil {
		return nil, err
	}

	return contact, nil
}
