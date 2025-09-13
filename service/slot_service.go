package service

import (
	"errors"

	"github.com/globalsign/mgo/bson"
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/mongo-lib/mserver"
)

/*==========================================================================================
SlotService
==========================================================================================*/

//SlotService ..
type SlotService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

//NewSlotService ..
func NewSlotService(dal mserver.DataAccessLayer, log *logging.Logger) *SlotService {
	return &SlotService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

//FindByID ..
func (u *SlotService) FindByID(id string) (*model.Slot, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching slot (if any) from Mongo
	slt, err := u.dal.Get(id, &model.Slot{})
	if err != nil {
		return nil, err
	}

	//we need to unmarshal the result from type bson.M{} to model.Slot
	//Convert BSON (byte) to JSON Fields
	var slot *model.Slot
	bsonBytes, _ := bson.Marshal(slt)
	bson.Unmarshal(bsonBytes, &slot)

	return slot, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

//CreateSlot will create a new Slot in Mongo using the Data Access Layer
func (u *SlotService) CreateSlot(slot *model.Slot) (*model.Slot, error) {
	//Validate required fields on Model element are being passed in
	if &slot.Status == nil {
		return nil, errors.New("Missing a required field Status: aborting before saving to the DB")
	}

	if &slot.Period == nil {
		return nil, errors.New("Missing a required field: Period aborting before saving to the DB")
	}

	id, err := u.dal.Post(slot)
	if err != nil {
		return nil, err
	}

	slot.Id = id
	return slot, nil
}
