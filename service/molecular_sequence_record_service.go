package service

import (
	"errors"
	"fmt"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
	"github.com/Kaleidoscope-Backup/mongo-lib/mserver"
	"github.com/globalsign/mgo/bson"
	logging "github.com/op/go-logging"
)

/*==========================================================================================
MolecularSequenceRecord  service
==========================================================================================*/

// MolecularSequenceRecordService ...
type MolecularSequenceRecordService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

// NewMolecularSequenceRecordService Data Access Layer
func NewMolecularSequenceRecordService(dal mserver.DataAccessLayer, log *logging.Logger) *MolecularSequenceRecordService {
	return &MolecularSequenceRecordService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

// FindByID ..
func (u *MolecularSequenceRecordService) FindByID(id string) (*model.MolecularSequenceRecord, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching NutritionOrderRecord (if any) from Mongo
	ms, err := u.dal.Get(id, &model.MolecularSequenceRecord{})
	if err != nil {
		return nil, err
	}

	var molSequence *model.MolecularSequenceRecord
	bsonBytes, _ := bson.Marshal(ms)
	bson.Unmarshal(bsonBytes, &molSequence)

	return molSequence, nil
}

// FindByConsumerID ..
func (u *MolecularSequenceRecordService) FindByConsumerID(id string) (*[]*model.MolecularSequenceRecord, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching NutritionOrderRecord (if any) from Mongo
	molArr, err := FindHealthRecordsByConsumerID(id, &model.MolecularSequenceRecord{}, u.dal)
	if err != nil {
		return nil, err
	}

	var molSequenceRecordArr []*model.MolecularSequenceRecord
	for _, mol := range molArr {
		var molSequenceRecord *model.MolecularSequenceRecord
		bsonBytes, _ := bson.Marshal(mol)
		bson.Unmarshal(bsonBytes, &molSequenceRecord)
		molSequenceRecordArr = append(molSequenceRecordArr, molSequenceRecord)
	}

	return &molSequenceRecordArr, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

// CreateMolecularSequenceRecord ...
func (u *MolecularSequenceRecordService) CreateMolecularSequenceRecord(molecularSequenceRecord *model.MolecularSequenceRecord) (*model.MolecularSequenceRecord, error) {

	id, err := u.dal.Post(molecularSequenceRecord)
	if err != nil {
		fmt.Printf("DB Create Failed For Record Error : %v", err)
		return nil, err
	}

	errHr := u.dal.PostWithID(id, &molecularSequenceRecord.HealthRecord)
	if errHr != nil {
		return nil, err
	}

	PostRecord(molecularSequenceRecord, molecularSequenceRecord.ConsumerID, id)
	return molecularSequenceRecord, nil
}

// Export ...
func (u *MolecularSequenceRecordService) Export(id string) (*[]model.HealthRecordExportElement, error) {
	if id == "" {
		return nil, errors.New("Missing parameter")
	}

	molecularSequenceRecords, _ := u.FindByConsumerID(id)

	if molecularSequenceRecords != nil {
		molecularSequenceRecordsIterable := *molecularSequenceRecords
		records := []model.HealthRecordExportElement{}
		for _, molecularSequenceRecord := range molecularSequenceRecordsIterable {
			recordElements, _ := u.ExportElements(molecularSequenceRecord)
			records = append(records, *recordElements...)
		}
		return &records, nil
	}

	return nil, nil
}

// ExportElements ...
func (u *MolecularSequenceRecordService) ExportElements(molecularSequenceRecord *model.MolecularSequenceRecord) (*[]model.HealthRecordExportElement, error) {

	if molecularSequenceRecord == nil {
		return nil, errors.New("Missing parameter")
	}

	// all records array
	records := []model.HealthRecordExportElement{}
	recordUUID := util.UUID()

	// Refrence sequence
	if molecularSequenceRecord.ReferenceSeq != nil {

		// Genome build
		valueGenomeBuild := model.Value{}
		recordGenomeBuild := model.HealthRecordExportElement{}

		valueGenomeBuild.ValueType = model.TEXT
		valueGenomeBuild.ValueText = &molecularSequenceRecord.ReferenceSeq.GenomeBuild

		recordGenomeBuild.Name = "Reference Sequence Genome Build"
		recordGenomeBuild.RecordID = recordUUID
		recordGenomeBuild.Value = valueGenomeBuild
		recordGenomeBuild.TimeStamp = molecularSequenceRecord.Occurred
		recordGenomeBuild.ConsumerID = molecularSequenceRecord.ConsumerID

		//populate the record
		records = append(records, recordGenomeBuild)

		// AccessionID
		valueAccessionID := model.Value{}
		recordAccessionID := model.HealthRecordExportElement{}

		valueAccessionID.ValueType = model.TEXT
		valueAccessionID.ValueText = molecularSequenceRecord.ReferenceSeq.AccessionID

		recordAccessionID.Name = "Reference Sequence AccessionID"
		recordAccessionID.RecordID = recordUUID
		recordAccessionID.Value = valueAccessionID
		recordAccessionID.TimeStamp = molecularSequenceRecord.Occurred
		recordAccessionID.ConsumerID = molecularSequenceRecord.ConsumerID

		//populate the record
		records = append(records, recordAccessionID)

		// ReferenceSeqString
		valueReferenceSeqString := model.Value{}
		recordReferenceSeqString := model.HealthRecordExportElement{}

		valueReferenceSeqString.ValueType = model.TEXT
		valueReferenceSeqString.ValueText = molecularSequenceRecord.ReferenceSeq.ReferenceSeqString

		recordReferenceSeqString.Name = "Reference Sequence ReferenceSeqString"
		recordReferenceSeqString.RecordID = recordUUID
		recordReferenceSeqString.Value = valueReferenceSeqString
		recordReferenceSeqString.TimeStamp = molecularSequenceRecord.Occurred
		recordReferenceSeqString.ConsumerID = molecularSequenceRecord.ConsumerID

		//populate the record
		records = append(records, recordReferenceSeqString)

		// WindowStart
		valueWindowStart := model.Value{}
		recordWindowStart := model.HealthRecordExportElement{}

		valueWindowStart.ValueType = model.QUANTITY
		valueWindowStart.ValueQuantity = molecularSequenceRecord.ReferenceSeq.WindowStart

		recordWindowStart.Name = "Reference Sequence WindowStart"
		recordWindowStart.RecordID = recordUUID
		recordWindowStart.Value = valueWindowStart
		recordWindowStart.TimeStamp = molecularSequenceRecord.Occurred
		recordWindowStart.ConsumerID = molecularSequenceRecord.ConsumerID

		//populate the record
		records = append(records, recordWindowStart)

		// WindowEnd
		valueWindowEnd := model.Value{}
		recordWindowEnd := model.HealthRecordExportElement{}

		valueWindowStart.ValueType = model.QUANTITY
		valueWindowStart.ValueQuantity = molecularSequenceRecord.ReferenceSeq.WindowEnd

		recordWindowEnd.Name = "Reference Sequence WindowEnd"
		recordWindowEnd.RecordID = recordUUID
		recordWindowEnd.Value = valueWindowEnd
		recordWindowEnd.TimeStamp = molecularSequenceRecord.Occurred
		recordWindowEnd.ConsumerID = molecularSequenceRecord.ConsumerID

		//populate the record
		records = append(records, recordWindowEnd)
	}

	// ObservedSeq
	if molecularSequenceRecord.ObservedSeq != nil {
		// ObservedSeq
		valueObservedSeq := model.Value{}
		recordObservedSeq := model.HealthRecordExportElement{}

		valueObservedSeq.ValueType = model.TEXT
		valueObservedSeq.ValueText = molecularSequenceRecord.ObservedSeq

		recordObservedSeq.Name = "Reference Sequence ObservedSeq"
		recordObservedSeq.RecordID = recordUUID
		recordObservedSeq.Value = valueObservedSeq
		recordObservedSeq.TimeStamp = molecularSequenceRecord.Occurred
		recordObservedSeq.ConsumerID = molecularSequenceRecord.ConsumerID

		//populate the record
		records = append(records, recordObservedSeq)
	}

	//Variants ...
	if molecularSequenceRecord.Variants != nil {
		variants := *molecularSequenceRecord.Variants

		for i := 0; i < len(variants); i++ {
			variant := variants[i]

			value := model.Value{}
			record := model.HealthRecordExportElement{}

			value.ValueType = model.TEXT
			value.ValueText = variant.ObservedAllele

			record.Name = "Variant"
			record.RecordID = recordUUID
			record.Value = value
			record.TimeStamp = molecularSequenceRecord.Occurred
			record.ConsumerID = molecularSequenceRecord.ConsumerID

			//populate the record
			records = append(records, record)
		}
	}

	return &records, nil
}
