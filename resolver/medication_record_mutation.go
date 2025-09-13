package resolver

import (
	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/constant"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/service"
	"golang.org/x/net/context"
)

// CreateMedicationRecord creates a new medication record in our Mongo DB and then returns the fields asked for in the graphql query
func (r *Resolver) CreateMedicationRecord(ctx context.Context, args *struct {
	MedicationRecord *model.MedicationRecordCreate
}) (*MedicationRecordResolver, error) {

	medicationRecord := &model.MedicationRecord{}
	medicationRecord.PrescribedBy = args.MedicationRecord.PrescribedBy
	medicationRecord.DispensingOrganization = args.MedicationRecord.DispensingOrganization
	medicationRecord.PrescribedOn = args.MedicationRecord.PrescribedOn
	medicationRecord.Expiration = args.MedicationRecord.Expiration

	medicationRecord.Medications = createMedications(ctx, args.MedicationRecord.Medications)

	healthRecord, er := CreateHealthRecord(ctx, &args.MedicationRecord.HealthRecordCreate, model.MEDICATION)
	if er != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", er)
		return nil, er
	}

	medicationRecord.HealthRecord = *healthRecord
	medicationRecord, err := ctx.Value(constant.MedicationRecordService).(*service.MedicationRecordService).CreateMedicationRecord(medicationRecord)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created medication : %v", *medicationRecord)
	healthRecordResolver := HealthRecordResolver{&medicationRecord.HealthRecord}
	return &MedicationRecordResolver{healthRecordResolver, medicationRecord}, nil
}

//CreateMedications will create an array of Medications object from an Array of MedicationCreate input objects
func createMedications(ctx context.Context, medicationsCreate *[]model.MedicationCreate) *[]model.Medication {
	var medications []model.Medication

	if medicationsCreate != nil && len(*medicationsCreate) > 0 {
		for i := 0; i < len(*medicationsCreate); i++ {
			var mc model.MedicationCreate
			mc = (*medicationsCreate)[i]
			if m, _ := createMedication(ctx, &mc); m != nil {
				medications = append(medications, *m)
			}
		}
		return &medications
	}

	return nil
}

//createMedication will create a medication object from a MedicationCreate input object
func createMedication(ctx context.Context, medicationCreate *model.MedicationCreate) (*model.Medication, error) {
	var medication *model.Medication
	medication = &model.Medication{}

	var dosage model.Dosage
	medication.Dosage = &dosage
	var strength model.Strength
	medication.Strength = &strength

	medication.MedicationStatus = medicationCreate.MedicationStatus
	medication.ProductName = medicationCreate.ProductName
	medication.IsOverTheCounter = medicationCreate.IsOverTheCounter
	medication.Route = medicationCreate.Route
	medication.Instructions = medicationCreate.Instructions
	medication.Dosage.Frequency = medicationCreate.DosageFrequency
	medication.Dosage.Unit = medicationCreate.DosageUnit
	medication.Dosage.Value = medicationCreate.DosageValue
	medication.RefillsRemaining = medicationCreate.RefillsRemaining
	medication.RefillsTotal = medicationCreate.RefillsTotal
	medication.Strength.Number = medicationCreate.StrengthNumber
	medication.Strength.Unit = medicationCreate.StrengthUnit
	medication.Start = medicationCreate.Start
	medication.End = medicationCreate.End

	if medicationCreate.Code != nil {
		code, err := CreateCodableConceptFromInput(ctx, medicationCreate.Code)
		if err != nil {
			ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
			return nil, err
		}
		medication.Code = code
	}

	return medication, nil
}
