package service

import (
	"errors"

	"github.com/globalsign/mgo/bson"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/mongo-lib/mserver"
	"github.com/op/go-logging"
)

/*==========================================================================================
MedicationService
==========================================================================================*/

// MedicationService ..
type MedicationService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

// NewMedicationService ..
func NewMedicationService(dal mserver.DataAccessLayer, log *logging.Logger) *MedicationService {
	return &MedicationService{dal: dal, log: log}
}

// FindById ...
func (u *MedicationService) FindById(id string) (*model.Medication, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching Medication (if any) from Mongo
	cct, err := u.dal.Get(id, &model.Medication{})
	if err != nil {
		return nil, err
	}

	//we need to unmarshal the result from type bson.M{} to model.Medication
	//Convert BSON (byte) to JSON Fields
	var medication *model.Medication
	bsonBytes, _ := bson.Marshal(cct)
	bson.Unmarshal(bsonBytes, &medication)

	return medication, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

// CreateMedication will create a new medication in Mongo using the Data Access Layer
func (u *MedicationService) CreateMedication(medication *model.Medication) (*model.Medication, error) {
	//Validate required fields on Model element are being passed in
	if medication.MedicationStatus == "" ||
		medication.ProductName == "" ||
		&medication.IsOverTheCounter == nil ||
		medication.Route == "" ||
		medication.Instructions == "" ||
		&medication.Start == nil {
		return nil, errors.New("Missing a required field: aborting before saving to the DB")
	}

	id, err := u.dal.Post(medication)
	if err != nil {
		return nil, err
	}

	medication.Id = id

	return medication, nil
}
